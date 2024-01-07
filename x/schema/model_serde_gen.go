// Code generated by mkunion. DO NOT EDIT.
package schema

import (
	"encoding/json"
	"fmt"
	"github.com/widmogrod/mkunion/x/shape"
	"github.com/widmogrod/mkunion/x/shared"
)

func init() {
	shape.Register(FieldShape())
}

var (
	_ json.Unmarshaler = (*Field)(nil)
	_ json.Marshaler   = (*Field)(nil)
)

func (r *Field) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONField(*r)
}
func (r *Field) _marshalJSONField(x Field) ([]byte, error) {
	partial := make(map[string]json.RawMessage)
	var err error
	var fieldName []byte
	fieldName, err = r._marshalJSONstring(x.Name)
	if err != nil {
		return nil, fmt.Errorf("schema: Field._marshalJSONField: field name Name; %w", err)
	}
	partial["Name"] = fieldName
	var fieldValue []byte
	fieldValue, err = r._marshalJSONSchema(x.Value)
	if err != nil {
		return nil, fmt.Errorf("schema: Field._marshalJSONField: field name Value; %w", err)
	}
	partial["Value"] = fieldValue
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("schema: Field._marshalJSONField: struct; %w", err)
	}
	return result, nil
}
func (r *Field) _marshalJSONstring(x string) ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, fmt.Errorf("schema: Field._marshalJSONstring:; %w", err)
	}
	return result, nil
}
func (r *Field) _marshalJSONSchema(x Schema) ([]byte, error) {
	result, err := shared.JSONMarshal[Schema](x)
	if err != nil {
		return nil, fmt.Errorf("schema: Field._marshalJSONSchema:; %w", err)
	}
	return result, nil
}
func (r *Field) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONField(data)
	if err != nil {
		return fmt.Errorf("schema: Field.UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *Field) _unmarshalJSONField(data []byte) (Field, error) {
	result := Field{}
	var partial map[string]json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("schema: Field._unmarshalJSONField: native struct unwrap; %w", err)
	}
	if fieldName, ok := partial["Name"]; ok {
		result.Name, err = r._unmarshalJSONstring(fieldName)
		if err != nil {
			return result, fmt.Errorf("schema: Field._unmarshalJSONField: field Name; %w", err)
		}
	}
	if fieldValue, ok := partial["Value"]; ok {
		result.Value, err = r._unmarshalJSONSchema(fieldValue)
		if err != nil {
			return result, fmt.Errorf("schema: Field._unmarshalJSONField: field Value; %w", err)
		}
	}
	return result, nil
}
func (r *Field) _unmarshalJSONstring(data []byte) (string, error) {
	var result string
	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, fmt.Errorf("schema: Field._unmarshalJSONstring: native primitive unwrap; %w", err)
	}
	return result, nil
}
func (r *Field) _unmarshalJSONSchema(data []byte) (Schema, error) {
	result, err := shared.JSONUnmarshal[Schema](data)
	if err != nil {
		return result, fmt.Errorf("schema: Field._unmarshalJSONSchema: native ref unwrap; %w", err)
	}
	return result, nil
}

//shape:shape
func FieldShape() shape.Shape {
	return &shape.StructLike{
		Name:          "Field",
		PkgName:       "schema",
		PkgImportName: "github.com/widmogrod/mkunion/x/schema",
		Fields: []*shape.FieldLike{
			{
				Name: "Name",
				Type: &shape.PrimitiveLike{Kind: &shape.StringLike{}},
			},
			{
				Name: "Value",
				Type: &shape.RefName{
					Name:          "Schema",
					PkgName:       "schema",
					PkgImportName: "github.com/widmogrod/mkunion/x/schema",
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
