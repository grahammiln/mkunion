package schemaless

import (
	"fmt"
	"github.com/widmogrod/mkunion/x/schema"
	"github.com/widmogrod/mkunion/x/storage/predicate"
	"sort"
	"sync"
)

func NewInMemoryRepository[A any]() *InMemoryRepository[A] {
	return &InMemoryRepository[A]{
		store:     make(map[string]Record[A]),
		appendLog: NewAppendLog[A](),
	}
}

var _ Repository[any] = (*InMemoryRepository[any])(nil)

type InMemoryRepository[A any] struct {
	store     map[string]Record[A]
	appendLog *AppendLog[A]
	mux       sync.RWMutex
}

func (s *InMemoryRepository[A]) Get(recordID, recordType string) (Record[A], error) {
	result, err := s.FindingRecords(FindingRecords[Record[A]]{
		RecordType: recordType,
		Where: predicate.MustWhere("ID = :id", predicate.ParamBinds{
			":id": schema.MkString(recordID),
		}),
		Limit: 1,
	})
	if err != nil {
		return Record[A]{}, err
	}

	if len(result.Items) == 0 {
		return Record[A]{}, ErrNotFound
	}

	return result.Items[0], nil
}

func (s *InMemoryRepository[A]) UpdateRecords(x UpdateRecords[Record[A]]) error {
	if x.IsEmpty() {
		return fmt.Errorf("store.InMemoryRepository.UpdateRecords: empty command %w", ErrEmptyCommand)
	}

	s.mux.Lock()
	defer s.mux.Unlock()

	newLog := NewAppendLog[A]()

	for _, record := range x.Saving {
		stored, ok := s.store[s.toKey(record)]
		if !ok {
			// new record, should have version 1
			// and since few lines below we increment version
			// we need to set it to 0
			record.Version = 0
			continue
		}

		//storedVersion := schema.AsDefault[uint16](schema.GetSchema(stored, "Version"), 0)
		storedVersion := stored.Version

		if x.UpdatingPolicy == PolicyIfServerNotChanged {
			if storedVersion != record.Version {
				return fmt.Errorf("store.InMemoryRepository.UpdateRecords ID=%s Type=%s %d != %d %w",
					record.ID, record.Type, storedVersion, record.Version, ErrVersionConflict)
			}
		} else if x.UpdatingPolicy == PolicyOverwriteServerChanges {
			record.Version = storedVersion
		}
	}

	for _, record := range x.Saving {
		var err error
		var before *Record[A]
		if b, ok := s.store[s.toKey(record)]; ok {
			before = &b
		}

		record.Version += 1
		s.store[s.toKey(record)] = record

		if before == nil {
			err = newLog.Change(Record[A]{}, record)
		} else {
			err = newLog.Change(*before, record)
		}
		if err != nil {
			panic(fmt.Errorf("store.InMemoryRepository.UpdateRecords: append log failed (1) %s %w", err, ErrInternalError))
		}
	}

	for _, record := range x.Deleting {
		if before, ok := s.store[s.toKey(record)]; ok {
			err := newLog.Delete(before)
			if err != nil {
				panic(fmt.Errorf("store.InMemoryRepository.UpdateRecords: append log failed (2) %s %w", err, ErrInternalError))
			}
		}

		delete(s.store, s.toKey(record))
	}

	s.appendLog.Append(newLog)

	return nil
}

func (s *InMemoryRepository[A]) toKey(record Record[A]) string {
	return record.ID + record.Type
}

func (s *InMemoryRepository[A]) FindingRecords(query FindingRecords[Record[A]]) (PageResult[Record[A]], error) {
	s.mux.RLock()
	defer s.mux.RUnlock()
	records := make([]Record[A], 0)
	for _, v := range s.store {
		records = append(records, v)
	}

	if query.RecordType != "" {
		newRecords := make([]Record[A], 0)
		for _, record := range records {
			if predicate.EvaluateEqual[Record[A]](record, "Type", query.RecordType) {
				newRecords = append(newRecords, record)
			}
		}
		records = newRecords
	}

	if query.Where != nil {
		newRecords := make([]Record[A], 0)
		for _, record := range records {
			if predicate.Evaluate[Record[A]](query.Where.Predicate, record, query.Where.Params) {
				newRecords = append(newRecords, record)
			}
		}
		records = newRecords
	}

	if len(query.Sort) > 0 {
		records = sortRecords(records, query.Sort)
	}

	if query.After != nil {
		found := false
		newRecords := make([]Record[A], 0)
		for _, record := range records {
			if predicate.EvaluateEqual[Record[A]](record, "ID", *query.After) {
				found = true
				continue // we're interested in records after this one
			}
			if found {
				newRecords = append(newRecords, record)
			}
		}
		records = newRecords
	}

	// Use limit to reduce number of records
	var next *FindingRecords[Record[A]]
	if query.Limit > 0 {
		if len(records) > int(query.Limit) {
			records = records[:query.Limit]

			next = &FindingRecords[Record[A]]{
				Where: query.Where,
				Sort:  query.Sort,
				Limit: query.Limit,
				After: &records[len(records)-1].ID,
			}
		}
	}

	result := PageResult[Record[A]]{
		Items: records,
		Next:  next,
	}

	return result, nil
}

func (s *InMemoryRepository[A]) AppendLog() *AppendLog[A] {
	return s.appendLog
}

func sortRecords[A any](records []Record[A], sortFields []SortField) []Record[A] {
	sort.Slice(records, func(i, j int) bool {
		for _, sortField := range sortFields {
			fieldA, _ := schema.Get[Record[A]](records[i], sortField.Field)
			fieldB, _ := schema.Get[Record[A]](records[j], sortField.Field)
			cmp := schema.Compare(fieldA, fieldB)
			if sortField.Descending {
				cmp = -cmp
			}
			if cmp != 0 {
				return cmp < 0
			}
		}
		return false
	})
	return records
}
