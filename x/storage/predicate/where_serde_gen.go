// Code generated by mkunion. DO NOT EDIT.
package predicate

import (
	"encoding/json"
	"fmt"
	"github.com/widmogrod/mkunion/x/shape"
	"github.com/widmogrod/mkunion/x/shared"
)

func init() {
	shape.Register(WherePredicatesShape())
}

var (
	_ json.Unmarshaler = (*WherePredicates)(nil)
	_ json.Marshaler   = (*WherePredicates)(nil)
)

func (r *WherePredicates) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONWherePredicates(*r)
}
func (r *WherePredicates) _marshalJSONWherePredicates(x WherePredicates) ([]byte, error) {
	partial := make(map[string]json.RawMessage)
	var err error
	var fieldPredicate []byte
	fieldPredicate, err = r._marshalJSONPredicate(x.Predicate)
	if err != nil {
		return nil, fmt.Errorf("predicate: WherePredicates._marshalJSONWherePredicates: field name Predicate; %w", err)
	}
	partial["Predicate"] = fieldPredicate
	var fieldParams []byte
	fieldParams, err = r._marshalJSONParamBinds(x.Params)
	if err != nil {
		return nil, fmt.Errorf("predicate: WherePredicates._marshalJSONWherePredicates: field name Params; %w", err)
	}
	partial["Params"] = fieldParams
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("predicate: WherePredicates._marshalJSONWherePredicates: struct; %w", err)
	}
	return result, nil
}
func (r *WherePredicates) _marshalJSONPredicate(x Predicate) ([]byte, error) {
	result, err := shared.JSONMarshal[Predicate](x)
	if err != nil {
		return nil, fmt.Errorf("predicate: WherePredicates._marshalJSONPredicate:; %w", err)
	}
	return result, nil
}
func (r *WherePredicates) _marshalJSONParamBinds(x ParamBinds) ([]byte, error) {
	result, err := shared.JSONMarshal[ParamBinds](x)
	if err != nil {
		return nil, fmt.Errorf("predicate: WherePredicates._marshalJSONParamBinds:; %w", err)
	}
	return result, nil
}
func (r *WherePredicates) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONWherePredicates(data)
	if err != nil {
		return fmt.Errorf("predicate: WherePredicates.UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *WherePredicates) _unmarshalJSONWherePredicates(data []byte) (WherePredicates, error) {
	result := WherePredicates{}
	var partial map[string]json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("predicate: WherePredicates._unmarshalJSONWherePredicates: native struct unwrap; %w", err)
	}
	if fieldPredicate, ok := partial["Predicate"]; ok {
		result.Predicate, err = r._unmarshalJSONPredicate(fieldPredicate)
		if err != nil {
			return result, fmt.Errorf("predicate: WherePredicates._unmarshalJSONWherePredicates: field Predicate; %w", err)
		}
	}
	if fieldParams, ok := partial["Params"]; ok {
		result.Params, err = r._unmarshalJSONParamBinds(fieldParams)
		if err != nil {
			return result, fmt.Errorf("predicate: WherePredicates._unmarshalJSONWherePredicates: field Params; %w", err)
		}
	}
	return result, nil
}
func (r *WherePredicates) _unmarshalJSONPredicate(data []byte) (Predicate, error) {
	result, err := shared.JSONUnmarshal[Predicate](data)
	if err != nil {
		return result, fmt.Errorf("predicate: WherePredicates._unmarshalJSONPredicate: native ref unwrap; %w", err)
	}
	return result, nil
}
func (r *WherePredicates) _unmarshalJSONParamBinds(data []byte) (ParamBinds, error) {
	result, err := shared.JSONUnmarshal[ParamBinds](data)
	if err != nil {
		return result, fmt.Errorf("predicate: WherePredicates._unmarshalJSONParamBinds: native ref unwrap; %w", err)
	}
	return result, nil
}

//shape:shape
func WherePredicatesShape() shape.Shape {
	return &shape.StructLike{
		Name:          "WherePredicates",
		PkgName:       "predicate",
		PkgImportName: "github.com/widmogrod/mkunion/x/storage/predicate",
		Fields: []*shape.FieldLike{
			{
				Name: "Predicate",
				Type: &shape.RefName{
					Name:          "Predicate",
					PkgName:       "predicate",
					PkgImportName: "github.com/widmogrod/mkunion/x/storage/predicate",
				},
			},
			{
				Name: "Params",
				Type: &shape.RefName{
					Name:          "ParamBinds",
					PkgName:       "predicate",
					PkgImportName: "github.com/widmogrod/mkunion/x/storage/predicate",
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
