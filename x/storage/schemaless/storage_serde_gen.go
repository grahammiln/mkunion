// Code generated by mkunion. DO NOT EDIT.
package schemaless

import (
	"encoding/json"
	"fmt"
	"github.com/widmogrod/mkunion/x/shape"
	"github.com/widmogrod/mkunion/x/shared"
)

func init() {
	shape.Register(SortFieldShape())
	shape.Register(PageResultShape())
	shape.Register(RecordShape())
}

var (
	_ json.Unmarshaler = (*SortField)(nil)
	_ json.Marshaler   = (*SortField)(nil)
)

func (r *SortField) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONSortField(*r)
}
func (r *SortField) _marshalJSONSortField(x SortField) ([]byte, error) {
	partial := make(map[string]json.RawMessage)
	var err error
	var fieldField []byte
	fieldField, err = r._marshalJSONstring(x.Field)
	if err != nil {
		return nil, fmt.Errorf("schemaless: SortField._marshalJSONSortField: field name Field; %w", err)
	}
	partial["Field"] = fieldField
	var fieldDescending []byte
	fieldDescending, err = r._marshalJSONbool(x.Descending)
	if err != nil {
		return nil, fmt.Errorf("schemaless: SortField._marshalJSONSortField: field name Descending; %w", err)
	}
	partial["Descending"] = fieldDescending
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("schemaless: SortField._marshalJSONSortField: struct; %w", err)
	}
	return result, nil
}
func (r *SortField) _marshalJSONstring(x string) ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, fmt.Errorf("schemaless: SortField._marshalJSONstring:; %w", err)
	}
	return result, nil
}
func (r *SortField) _marshalJSONbool(x bool) ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, fmt.Errorf("schemaless: SortField._marshalJSONbool:; %w", err)
	}
	return result, nil
}
func (r *SortField) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONSortField(data)
	if err != nil {
		return fmt.Errorf("schemaless: SortField.UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *SortField) _unmarshalJSONSortField(data []byte) (SortField, error) {
	result := SortField{}
	var partial map[string]json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("schemaless: SortField._unmarshalJSONSortField: native struct unwrap; %w", err)
	}
	if fieldField, ok := partial["Field"]; ok {
		result.Field, err = r._unmarshalJSONstring(fieldField)
		if err != nil {
			return result, fmt.Errorf("schemaless: SortField._unmarshalJSONSortField: field Field; %w", err)
		}
	}
	if fieldDescending, ok := partial["Descending"]; ok {
		result.Descending, err = r._unmarshalJSONbool(fieldDescending)
		if err != nil {
			return result, fmt.Errorf("schemaless: SortField._unmarshalJSONSortField: field Descending; %w", err)
		}
	}
	return result, nil
}
func (r *SortField) _unmarshalJSONstring(data []byte) (string, error) {
	var result string
	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, fmt.Errorf("schemaless: SortField._unmarshalJSONstring: native primitive unwrap; %w", err)
	}
	return result, nil
}
func (r *SortField) _unmarshalJSONbool(data []byte) (bool, error) {
	var result bool
	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, fmt.Errorf("schemaless: SortField._unmarshalJSONbool: native primitive unwrap; %w", err)
	}
	return result, nil
}
func SortFieldShape() shape.Shape {
	return &shape.StructLike{
		Name:          "SortField",
		PkgName:       "schemaless",
		PkgImportName: "github.com/widmogrod/mkunion/x/storage/schemaless",
		Fields: []*shape.FieldLike{
			{
				Name: "Field",
				Type: &shape.PrimitiveLike{Kind: &shape.StringLike{}},
			},
			{
				Name: "Descending",
				Type: &shape.PrimitiveLike{Kind: &shape.BooleanLike{}},
			},
		},
		Tags: map[string]shape.Tag{
			"serde": {
				Value: "json",
			},
		},
	}
}

var (
	_ json.Unmarshaler = (*PageResult[any])(nil)
	_ json.Marshaler   = (*PageResult[any])(nil)
)

func (r *PageResult[A]) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONPageResultLb_A_bL(*r)
}
func (r *PageResult[A]) _marshalJSONPageResultLb_A_bL(x PageResult[A]) ([]byte, error) {
	partial := make(map[string]json.RawMessage)
	var err error
	var fieldItems []byte
	fieldItems, err = r._marshalJSONSliceA(x.Items)
	if err != nil {
		return nil, fmt.Errorf("schemaless: PageResult[A]._marshalJSONPageResultLb_A_bL: field name Items; %w", err)
	}
	partial["Items"] = fieldItems
	var fieldNext []byte
	fieldNext, err = r._marshalJSONPtrFindingRecordsLb_A_bL(x.Next)
	if err != nil {
		return nil, fmt.Errorf("schemaless: PageResult[A]._marshalJSONPageResultLb_A_bL: field name Next; %w", err)
	}
	if fieldNext != nil {
		partial["Next"] = fieldNext
	}
	var fieldPrev []byte
	fieldPrev, err = r._marshalJSONPtrFindingRecordsLb_A_bL(x.Prev)
	if err != nil {
		return nil, fmt.Errorf("schemaless: PageResult[A]._marshalJSONPageResultLb_A_bL: field name Prev; %w", err)
	}
	if fieldPrev != nil {
		partial["Prev"] = fieldPrev
	}
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("schemaless: PageResult[A]._marshalJSONPageResultLb_A_bL: struct; %w", err)
	}
	return result, nil
}
func (r *PageResult[A]) _marshalJSONSliceA(x []A) ([]byte, error) {
	partial := make([]json.RawMessage, len(x))
	for i, v := range x {
		item, err := r._marshalJSONA(v)
		if err != nil {
			return nil, fmt.Errorf("schemaless: PageResult[A]._marshalJSONSliceA: at index %d; %w", i, err)
		}
		partial[i] = item
	}
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("schemaless: PageResult[A]._marshalJSONSliceA:; %w", err)
	}
	return result, nil
}
func (r *PageResult[A]) _marshalJSONA(x A) ([]byte, error) {
	result, err := shared.JSONMarshal[A](x)
	if err != nil {
		return nil, fmt.Errorf("schemaless: PageResult[A]._marshalJSONA:; %w", err)
	}
	return result, nil
}
func (r *PageResult[A]) _marshalJSONPtrFindingRecordsLb_A_bL(x *FindingRecords[A]) ([]byte, error) {
	if x == nil {
		return nil, nil
	}
	return r._marshalJSONFindingRecordsLb_A_bL(*x)
}
func (r *PageResult[A]) _marshalJSONFindingRecordsLb_A_bL(x FindingRecords[A]) ([]byte, error) {
	result, err := shared.JSONMarshal[FindingRecords[A]](x)
	if err != nil {
		return nil, fmt.Errorf("schemaless: PageResult[A]._marshalJSONFindingRecordsLb_A_bL:; %w", err)
	}
	return result, nil
}
func (r *PageResult[A]) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONPageResultLb_A_bL(data)
	if err != nil {
		return fmt.Errorf("schemaless: PageResult[A].UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *PageResult[A]) _unmarshalJSONPageResultLb_A_bL(data []byte) (PageResult[A], error) {
	result := PageResult[A]{}
	var partial map[string]json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("schemaless: PageResult[A]._unmarshalJSONPageResultLb_A_bL: native struct unwrap; %w", err)
	}
	if fieldItems, ok := partial["Items"]; ok {
		result.Items, err = r._unmarshalJSONSliceA(fieldItems)
		if err != nil {
			return result, fmt.Errorf("schemaless: PageResult[A]._unmarshalJSONPageResultLb_A_bL: field Items; %w", err)
		}
	}
	if fieldNext, ok := partial["Next"]; ok {
		result.Next, err = r._unmarshalJSONPtrFindingRecordsLb_A_bL(fieldNext)
		if err != nil {
			return result, fmt.Errorf("schemaless: PageResult[A]._unmarshalJSONPageResultLb_A_bL: field Next; %w", err)
		}
	}
	if fieldPrev, ok := partial["Prev"]; ok {
		result.Prev, err = r._unmarshalJSONPtrFindingRecordsLb_A_bL(fieldPrev)
		if err != nil {
			return result, fmt.Errorf("schemaless: PageResult[A]._unmarshalJSONPageResultLb_A_bL: field Prev; %w", err)
		}
	}
	return result, nil
}
func (r *PageResult[A]) _unmarshalJSONSliceA(data []byte) ([]A, error) {
	result := make([]A, 0)
	var partial []json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("schemaless: PageResult[A]._unmarshalJSONSliceA: native list unwrap; %w", err)
	}
	for i, v := range partial {
		item, err := r._unmarshalJSONA(v)
		if err != nil {
			return result, fmt.Errorf("schemaless: PageResult[A]._unmarshalJSONSliceA: at index %d; %w", i, err)
		}
		result = append(result, item)
	}
	return result, nil
}
func (r *PageResult[A]) _unmarshalJSONA(data []byte) (A, error) {
	result, err := shared.JSONUnmarshal[A](data)
	if err != nil {
		return result, fmt.Errorf("schemaless: PageResult[A]._unmarshalJSONA: native ref unwrap; %w", err)
	}
	return result, nil
}
func (r *PageResult[A]) _unmarshalJSONPtrFindingRecordsLb_A_bL(data []byte) (*FindingRecords[A], error) {
	if len(data) == 0 {
		return nil, nil
	}
	if string(data[:4]) == "null" {
		return nil, nil
	}
	result, err := r._unmarshalJSONFindingRecordsLb_A_bL(data)
	if err != nil {
		return nil, fmt.Errorf("schemaless: PageResult[A]._unmarshalJSONPtrFindingRecordsLb_A_bL: pointer; %w", err)
	}
	return &result, nil
}
func (r *PageResult[A]) _unmarshalJSONFindingRecordsLb_A_bL(data []byte) (FindingRecords[A], error) {
	result, err := shared.JSONUnmarshal[FindingRecords[A]](data)
	if err != nil {
		return result, fmt.Errorf("schemaless: PageResult[A]._unmarshalJSONFindingRecordsLb_A_bL: native ref unwrap; %w", err)
	}
	return result, nil
}
func PageResultShape() shape.Shape {
	return &shape.StructLike{
		Name:          "PageResult",
		PkgName:       "schemaless",
		PkgImportName: "github.com/widmogrod/mkunion/x/storage/schemaless",
		TypeParams: []shape.TypeParam{
			shape.TypeParam{
				Name: "A",
				Type: &shape.Any{},
			},
		},
		Fields: []*shape.FieldLike{
			{
				Name: "Items",
				Type: &shape.ListLike{
					Element: &shape.RefName{
						Name:          "A",
						PkgName:       "",
						PkgImportName: "",
					},
				},
			},
			{
				Name: "Next",
				Type: &shape.PointerLike{
					Type: &shape.RefName{
						Name:          "FindingRecords",
						PkgName:       "schemaless",
						PkgImportName: "github.com/widmogrod/mkunion/x/storage/schemaless",
						Indexed: []shape.Shape{
							&shape.RefName{
								Name:          "A",
								PkgName:       "",
								PkgImportName: "",
							},
						},
					},
				},
			},
			{
				Name: "Prev",
				Type: &shape.PointerLike{
					Type: &shape.RefName{
						Name:          "FindingRecords",
						PkgName:       "schemaless",
						PkgImportName: "github.com/widmogrod/mkunion/x/storage/schemaless",
						Indexed: []shape.Shape{
							&shape.RefName{
								Name:          "A",
								PkgName:       "",
								PkgImportName: "",
							},
						},
					},
				},
			},
		},
		Tags: map[string]shape.Tag{
			"serde": {
				Value: "json",
			},
		},
	}
}

var (
	_ json.Unmarshaler = (*Record[any])(nil)
	_ json.Marshaler   = (*Record[any])(nil)
)

func (r *Record[A]) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONRecordLb_A_bL(*r)
}
func (r *Record[A]) _marshalJSONRecordLb_A_bL(x Record[A]) ([]byte, error) {
	partial := make(map[string]json.RawMessage)
	var err error
	var fieldID []byte
	fieldID, err = r._marshalJSONstring(x.ID)
	if err != nil {
		return nil, fmt.Errorf("schemaless: Record[A]._marshalJSONRecordLb_A_bL: field name ID; %w", err)
	}
	partial["ID"] = fieldID
	var fieldType []byte
	fieldType, err = r._marshalJSONstring(x.Type)
	if err != nil {
		return nil, fmt.Errorf("schemaless: Record[A]._marshalJSONRecordLb_A_bL: field name Type; %w", err)
	}
	partial["Type"] = fieldType
	var fieldData []byte
	fieldData, err = r._marshalJSONA(x.Data)
	if err != nil {
		return nil, fmt.Errorf("schemaless: Record[A]._marshalJSONRecordLb_A_bL: field name Data; %w", err)
	}
	partial["Data"] = fieldData
	var fieldVersion []byte
	fieldVersion, err = r._marshalJSONuint16(x.Version)
	if err != nil {
		return nil, fmt.Errorf("schemaless: Record[A]._marshalJSONRecordLb_A_bL: field name Version; %w", err)
	}
	partial["Version"] = fieldVersion
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("schemaless: Record[A]._marshalJSONRecordLb_A_bL: struct; %w", err)
	}
	return result, nil
}
func (r *Record[A]) _marshalJSONstring(x string) ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, fmt.Errorf("schemaless: Record[A]._marshalJSONstring:; %w", err)
	}
	return result, nil
}
func (r *Record[A]) _marshalJSONA(x A) ([]byte, error) {
	result, err := shared.JSONMarshal[A](x)
	if err != nil {
		return nil, fmt.Errorf("schemaless: Record[A]._marshalJSONA:; %w", err)
	}
	return result, nil
}
func (r *Record[A]) _marshalJSONuint16(x uint16) ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, fmt.Errorf("schemaless: Record[A]._marshalJSONuint16:; %w", err)
	}
	return result, nil
}
func (r *Record[A]) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONRecordLb_A_bL(data)
	if err != nil {
		return fmt.Errorf("schemaless: Record[A].UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *Record[A]) _unmarshalJSONRecordLb_A_bL(data []byte) (Record[A], error) {
	result := Record[A]{}
	var partial map[string]json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("schemaless: Record[A]._unmarshalJSONRecordLb_A_bL: native struct unwrap; %w", err)
	}
	if fieldID, ok := partial["ID"]; ok {
		result.ID, err = r._unmarshalJSONstring(fieldID)
		if err != nil {
			return result, fmt.Errorf("schemaless: Record[A]._unmarshalJSONRecordLb_A_bL: field ID; %w", err)
		}
	}
	if fieldType, ok := partial["Type"]; ok {
		result.Type, err = r._unmarshalJSONstring(fieldType)
		if err != nil {
			return result, fmt.Errorf("schemaless: Record[A]._unmarshalJSONRecordLb_A_bL: field Type; %w", err)
		}
	}
	if fieldData, ok := partial["Data"]; ok {
		result.Data, err = r._unmarshalJSONA(fieldData)
		if err != nil {
			return result, fmt.Errorf("schemaless: Record[A]._unmarshalJSONRecordLb_A_bL: field Data; %w", err)
		}
	}
	if fieldVersion, ok := partial["Version"]; ok {
		result.Version, err = r._unmarshalJSONuint16(fieldVersion)
		if err != nil {
			return result, fmt.Errorf("schemaless: Record[A]._unmarshalJSONRecordLb_A_bL: field Version; %w", err)
		}
	}
	return result, nil
}
func (r *Record[A]) _unmarshalJSONstring(data []byte) (string, error) {
	var result string
	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, fmt.Errorf("schemaless: Record[A]._unmarshalJSONstring: native primitive unwrap; %w", err)
	}
	return result, nil
}
func (r *Record[A]) _unmarshalJSONA(data []byte) (A, error) {
	result, err := shared.JSONUnmarshal[A](data)
	if err != nil {
		return result, fmt.Errorf("schemaless: Record[A]._unmarshalJSONA: native ref unwrap; %w", err)
	}
	return result, nil
}
func (r *Record[A]) _unmarshalJSONuint16(data []byte) (uint16, error) {
	var result uint16
	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, fmt.Errorf("schemaless: Record[A]._unmarshalJSONuint16: native primitive unwrap; %w", err)
	}
	return result, nil
}
func RecordShape() shape.Shape {
	return &shape.StructLike{
		Name:          "Record",
		PkgName:       "schemaless",
		PkgImportName: "github.com/widmogrod/mkunion/x/storage/schemaless",
		TypeParams: []shape.TypeParam{
			shape.TypeParam{
				Name: "A",
				Type: &shape.Any{},
			},
		},
		Fields: []*shape.FieldLike{
			{
				Name: "ID",
				Type: &shape.PrimitiveLike{Kind: &shape.StringLike{}},
			},
			{
				Name: "Type",
				Type: &shape.PrimitiveLike{Kind: &shape.StringLike{}},
			},
			{
				Name: "Data",
				Type: &shape.RefName{
					Name:          "A",
					PkgName:       "",
					PkgImportName: "",
				},
			},
			{
				Name: "Version",
				Type: &shape.PrimitiveLike{
					Kind: &shape.NumberLike{
						Kind: &shape.UInt16{},
					},
				},
			},
		},
		Tags: map[string]shape.Tag{
			"serde": {
				Value: "json",
			},
		},
	}
}
