// Code generated by mkunion. DO NOT EDIT.
package schemaless

import (
	"encoding/json"
	"fmt"
	"github.com/widmogrod/mkunion/x/shape"
	"github.com/widmogrod/mkunion/x/shared"
)

func init() {
	shape.Register(PageResultShape())
	shape.Register(RecordShape())
}

var (
	_ json.Unmarshaler = (*PageResult[any])(nil)
	_ json.Marshaler   = (*PageResult[any])(nil)
)

func (r *PageResult[A]) MarshalJSON() ([]byte, error) {
	var err error
	result := make(map[string]json.RawMessage)

	fieldItems := make([]json.RawMessage, len(r.Items))
	for i, v := range r.Items {
		fieldItems[i], err = shared.JSONMarshal[A](v)
		if err != nil {
			return nil, fmt.Errorf("schemaless.PageResult[A].MarshalJSON: field Items[%d]; %w", i, err)
		}
	}
	result["Items"], err = json.Marshal(fieldItems)
	if err != nil {
		return nil, fmt.Errorf("schemaless.PageResult[A].MarshalJSON: field Items; %w", err)
	}

	if r.Next != nil {
		fieldNext, err := shared.JSONMarshal[*FindingRecords[A]](r.Next)
		if err != nil {
			return nil, fmt.Errorf("schemaless.PageResult[A].MarshalJSON: field Next; %w", err)
		}
		result["Next"] = fieldNext
	}

	output, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("schemaless.PageResult[A].MarshalJSON: final step; %w", err)
	}

	return output, nil
}

func (r *PageResult[A]) UnmarshalJSON(bytes []byte) error {
	return shared.JSONParseObject(bytes, func(key string, bytes []byte) error {
		switch key {
		case "Items":
			err := shared.JSONParseList(bytes, func(index int, bytes []byte) error {
				item, err := shared.JSONUnmarshal[A](bytes)
				if err != nil {
					return fmt.Errorf("schemaless.PageResult[A].UnmarshalJSON: field Items[%d]; %w", index, err)
				}
				r.Items = append(r.Items, item)
				return nil
			})
			if err != nil {
				return fmt.Errorf("schemaless.PageResult[A].UnmarshalJSON: field Items; %w", err)
			}
			return nil

		case "Next":
			var err error
			r.Next, err = shared.JSONUnmarshal[*FindingRecords[A]](bytes)
			if err != nil {
				return fmt.Errorf("schemaless.PageResult[A].UnmarshalJSON: field Next; %w", err)
			}
			return nil

		}

		return nil
	})
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
				Type: &shape.RefName{
					Name:          "FindingRecords",
					PkgName:       "schemaless",
					PkgImportName: "github.com/widmogrod/mkunion/x/storage/schemaless",
					IsPointer:     true,
					Indexed: []shape.Shape{
						&shape.RefName{
							Name:          "A",
							PkgName:       "",
							PkgImportName: "",
						},
					},
				},
				IsPointer: true,
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
	var err error
	result := make(map[string]json.RawMessage)

	fieldID, err := shared.JSONMarshal[string](r.ID)
	if err != nil {
		return nil, fmt.Errorf("schemaless.Record[A].MarshalJSON: field ID; %w", err)
	}
	result["ID"] = fieldID

	fieldType, err := shared.JSONMarshal[string](r.Type)
	if err != nil {
		return nil, fmt.Errorf("schemaless.Record[A].MarshalJSON: field Type; %w", err)
	}
	result["Type"] = fieldType

	fieldData, err := shared.JSONMarshal[A](r.Data)
	if err != nil {
		return nil, fmt.Errorf("schemaless.Record[A].MarshalJSON: field Data; %w", err)
	}
	result["Data"] = fieldData

	fieldVersion, err := shared.JSONMarshal[uint16](r.Version)
	if err != nil {
		return nil, fmt.Errorf("schemaless.Record[A].MarshalJSON: field Version; %w", err)
	}
	result["Version"] = fieldVersion

	output, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("schemaless.Record[A].MarshalJSON: final step; %w", err)
	}

	return output, nil
}

func (r *Record[A]) UnmarshalJSON(bytes []byte) error {
	return shared.JSONParseObject(bytes, func(key string, bytes []byte) error {
		switch key {
		case "ID":
			var err error
			r.ID, err = shared.JSONUnmarshal[string](bytes)
			if err != nil {
				return fmt.Errorf("schemaless.Record[A].UnmarshalJSON: field ID; %w", err)
			}
			return nil

		case "Type":
			var err error
			r.Type, err = shared.JSONUnmarshal[string](bytes)
			if err != nil {
				return fmt.Errorf("schemaless.Record[A].UnmarshalJSON: field Type; %w", err)
			}
			return nil

		case "Data":
			var err error
			r.Data, err = shared.JSONUnmarshal[A](bytes)
			if err != nil {
				return fmt.Errorf("schemaless.Record[A].UnmarshalJSON: field Data; %w", err)
			}
			return nil

		case "Version":
			var err error
			r.Version, err = shared.JSONUnmarshal[uint16](bytes)
			if err != nil {
				return fmt.Errorf("schemaless.Record[A].UnmarshalJSON: field Version; %w", err)
			}
			return nil

		}

		return nil
	})
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
				Type: &shape.StringLike{},
			},
			{
				Name: "Type",
				Type: &shape.StringLike{},
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
				Type: &shape.NumberLike{
					Kind: &shape.UInt16{},
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
