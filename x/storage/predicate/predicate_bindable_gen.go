// Code generated by mkunion. DO NOT EDIT.
package predicate

import (
	"encoding/json"
	"fmt"
	"github.com/widmogrod/mkunion/f"
	"github.com/widmogrod/mkunion/x/schema"
	"github.com/widmogrod/mkunion/x/shape"
	"github.com/widmogrod/mkunion/x/shared"
)

func init() {
	shape.Register(BindableShape())
	shape.Register(BindValueShape())
	shape.Register(LiteralShape())
	shape.Register(LocatableShape())
}

//mkunion-extension:visitor

type BindableVisitor interface {
	VisitBindValue(v *BindValue) any
	VisitLiteral(v *Literal) any
	VisitLocatable(v *Locatable) any
}

type Bindable interface {
	AcceptBindable(g BindableVisitor) any
}

func (r *BindValue) AcceptBindable(v BindableVisitor) any { return v.VisitBindValue(r) }
func (r *Literal) AcceptBindable(v BindableVisitor) any   { return v.VisitLiteral(r) }
func (r *Locatable) AcceptBindable(v BindableVisitor) any { return v.VisitLocatable(r) }

var (
	_ Bindable = (*BindValue)(nil)
	_ Bindable = (*Literal)(nil)
	_ Bindable = (*Locatable)(nil)
)

func MatchBindable[TOut any](
	x Bindable,
	f1 func(x *BindValue) TOut,
	f2 func(x *Literal) TOut,
	f3 func(x *Locatable) TOut,
	df func(x Bindable) TOut,
) TOut {
	return f.Match3(x, f1, f2, f3, df)
}

func MatchBindableR2[TOut1, TOut2 any](
	x Bindable,
	f1 func(x *BindValue) (TOut1, TOut2),
	f2 func(x *Literal) (TOut1, TOut2),
	f3 func(x *Locatable) (TOut1, TOut2),
	df func(x Bindable) (TOut1, TOut2),
) (TOut1, TOut2) {
	return f.Match3R2(x, f1, f2, f3, df)
}

func MustMatchBindable[TOut any](
	x Bindable,
	f1 func(x *BindValue) TOut,
	f2 func(x *Literal) TOut,
	f3 func(x *Locatable) TOut,
) TOut {
	return f.MustMatch3(x, f1, f2, f3)
}

func MustMatchBindableR0(
	x Bindable,
	f1 func(x *BindValue),
	f2 func(x *Literal),
	f3 func(x *Locatable),
) {
	f.MustMatch3R0(x, f1, f2, f3)
}

func MustMatchBindableR2[TOut1, TOut2 any](
	x Bindable,
	f1 func(x *BindValue) (TOut1, TOut2),
	f2 func(x *Literal) (TOut1, TOut2),
	f3 func(x *Locatable) (TOut1, TOut2),
) (TOut1, TOut2) {
	return f.MustMatch3R2(x, f1, f2, f3)
}

// mkunion-extension:shape
func BindableShape() shape.Shape {
	return &shape.UnionLike{
		Name:          "Bindable",
		PkgName:       "predicate",
		PkgImportName: "github.com/widmogrod/mkunion/x/storage/predicate",
		Variant: []shape.Shape{
			BindValueShape(),
			LiteralShape(),
			LocatableShape(),
		},
	}
}
func BindValueShape() shape.Shape {
	return &shape.StructLike{
		Name:          "BindValue",
		PkgName:       "predicate",
		PkgImportName: "github.com/widmogrod/mkunion/x/storage/predicate",
		Fields: []*shape.FieldLike{
			{
				Name: "BindName",
				Type: &shape.RefName{
					Name:          "BindName",
					PkgName:       "predicate",
					PkgImportName: "github.com/widmogrod/mkunion/x/storage/predicate",
				},
			},
		},
	}
}
func LiteralShape() shape.Shape {
	return &shape.StructLike{
		Name:          "Literal",
		PkgName:       "predicate",
		PkgImportName: "github.com/widmogrod/mkunion/x/storage/predicate",
		Fields: []*shape.FieldLike{
			{
				Name: "Value",
				Type: &shape.RefName{
					Name:          "Schema",
					PkgName:       "schema",
					PkgImportName: "github.com/widmogrod/mkunion/x/schema",
				},
			},
		},
	}
}
func LocatableShape() shape.Shape {
	return &shape.StructLike{
		Name:          "Locatable",
		PkgName:       "predicate",
		PkgImportName: "github.com/widmogrod/mkunion/x/storage/predicate",
		Fields: []*shape.FieldLike{
			{
				Name: "Location",
				Type: &shape.StringLike{},
			},
		},
	}
}

// mkunion-extension:json
func init() {
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/storage/predicate.Bindable", BindableFromJSON, BindableToJSON)
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/storage/predicate.BindValue", BindValueFromJSON, BindValueToJSON)
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/storage/predicate.Literal", LiteralFromJSON, LiteralToJSON)
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/storage/predicate.Locatable", LocatableFromJSON, LocatableToJSON)
}

type BindableUnionJSON struct {
	Type      string          `json:"$type,omitempty"`
	BindValue json.RawMessage `json:"predicate.BindValue,omitempty"`
	Literal   json.RawMessage `json:"predicate.Literal,omitempty"`
	Locatable json.RawMessage `json:"predicate.Locatable,omitempty"`
}

func BindableFromJSON(x []byte) (Bindable, error) {
	if x == nil || len(x) == 0 {
		return nil, nil
	}
	if string(x[:4]) == "null" {
		return nil, nil
	}

	var data BindableUnionJSON
	err := json.Unmarshal(x, &data)
	if err != nil {
		return nil, err
	}

	switch data.Type {
	case "predicate.BindValue":
		return BindValueFromJSON(data.BindValue)
	case "predicate.Literal":
		return LiteralFromJSON(data.Literal)
	case "predicate.Locatable":
		return LocatableFromJSON(data.Locatable)
	}

	if data.BindValue != nil {
		return BindValueFromJSON(data.BindValue)
	} else if data.Literal != nil {
		return LiteralFromJSON(data.Literal)
	} else if data.Locatable != nil {
		return LocatableFromJSON(data.Locatable)
	}

	return nil, fmt.Errorf("predicate.Bindable: unknown type %s", data.Type)
}

func BindableToJSON(x Bindable) ([]byte, error) {
	if x == nil {
		return nil, nil
	}
	return MustMatchBindableR2(
		x,
		func(x *BindValue) ([]byte, error) {
			body, err := BindValueToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(BindableUnionJSON{
				Type:      "predicate.BindValue",
				BindValue: body,
			})
		},
		func(x *Literal) ([]byte, error) {
			body, err := LiteralToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(BindableUnionJSON{
				Type:    "predicate.Literal",
				Literal: body,
			})
		},
		func(x *Locatable) ([]byte, error) {
			body, err := LocatableToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(BindableUnionJSON{
				Type:      "predicate.Locatable",
				Locatable: body,
			})
		},
	)
}

func BindValueFromJSON(x []byte) (*BindValue, error) {
	result := new(BindValue)
	err := result.UnmarshalJSON(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func BindValueToJSON(x *BindValue) ([]byte, error) {
	return x.MarshalJSON()
}

var (
	_ json.Unmarshaler = (*BindValue)(nil)
	_ json.Marshaler   = (*BindValue)(nil)
)

func (r *BindValue) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONBindValue(*r)
}
func (r *BindValue) _marshalJSONBindValue(x BindValue) ([]byte, error) {
	partial := make(map[string]json.RawMessage)
	var err error
	var fieldBindName []byte
	fieldBindName, err = r._marshalJSONBindName(x.BindName)
	if err != nil {
		return nil, fmt.Errorf("predicate: BindValue._marshalJSONBindValue: field name BindName; %w", err)
	}
	partial["BindName"] = fieldBindName
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("predicate: BindValue._marshalJSONBindValue: struct; %w", err)
	}
	return result, nil
}
func (r *BindValue) _marshalJSONBindName(x BindName) ([]byte, error) {
	result, err := shared.JSONMarshal[BindName](x)
	if err != nil {
		return nil, fmt.Errorf("predicate: BindValue._marshalJSONBindName:; %w", err)
	}
	return result, nil
}
func (r *BindValue) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONBindValue(data)
	if err != nil {
		return fmt.Errorf("predicate: BindValue.UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *BindValue) _unmarshalJSONBindValue(data []byte) (BindValue, error) {
	result := BindValue{}
	var partial map[string]json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("predicate: BindValue._unmarshalJSONBindValue: native struct unwrap; %w", err)
	}
	if fieldBindName, ok := partial["BindName"]; ok {
		result.BindName, err = r._unmarshalJSONBindName(fieldBindName)
		if err != nil {
			return result, fmt.Errorf("predicate: BindValue._unmarshalJSONBindValue: field BindName; %w", err)
		}
	}
	return result, nil
}
func (r *BindValue) _unmarshalJSONBindName(data []byte) (BindName, error) {
	result, err := shared.JSONUnmarshal[BindName](data)
	if err != nil {
		return result, fmt.Errorf("predicate: BindValue._unmarshalJSONBindName: native ref unwrap; %w", err)
	}
	return result, nil
}

func LiteralFromJSON(x []byte) (*Literal, error) {
	result := new(Literal)
	err := result.UnmarshalJSON(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func LiteralToJSON(x *Literal) ([]byte, error) {
	return x.MarshalJSON()
}

var (
	_ json.Unmarshaler = (*Literal)(nil)
	_ json.Marshaler   = (*Literal)(nil)
)

func (r *Literal) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONLiteral(*r)
}
func (r *Literal) _marshalJSONLiteral(x Literal) ([]byte, error) {
	partial := make(map[string]json.RawMessage)
	var err error
	var fieldValue []byte
	fieldValue, err = r._marshalJSONschema_Schema(x.Value)
	if err != nil {
		return nil, fmt.Errorf("predicate: Literal._marshalJSONLiteral: field name Value; %w", err)
	}
	partial["Value"] = fieldValue
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("predicate: Literal._marshalJSONLiteral: struct; %w", err)
	}
	return result, nil
}
func (r *Literal) _marshalJSONschema_Schema(x schema.Schema) ([]byte, error) {
	result, err := shared.JSONMarshal[schema.Schema](x)
	if err != nil {
		return nil, fmt.Errorf("predicate: Literal._marshalJSONschema_Schema:; %w", err)
	}
	return result, nil
}
func (r *Literal) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONLiteral(data)
	if err != nil {
		return fmt.Errorf("predicate: Literal.UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *Literal) _unmarshalJSONLiteral(data []byte) (Literal, error) {
	result := Literal{}
	var partial map[string]json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("predicate: Literal._unmarshalJSONLiteral: native struct unwrap; %w", err)
	}
	if fieldValue, ok := partial["Value"]; ok {
		result.Value, err = r._unmarshalJSONschema_Schema(fieldValue)
		if err != nil {
			return result, fmt.Errorf("predicate: Literal._unmarshalJSONLiteral: field Value; %w", err)
		}
	}
	return result, nil
}
func (r *Literal) _unmarshalJSONschema_Schema(data []byte) (schema.Schema, error) {
	result, err := shared.JSONUnmarshal[schema.Schema](data)
	if err != nil {
		return result, fmt.Errorf("predicate: Literal._unmarshalJSONschema_Schema: native ref unwrap; %w", err)
	}
	return result, nil
}

func LocatableFromJSON(x []byte) (*Locatable, error) {
	result := new(Locatable)
	err := result.UnmarshalJSON(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func LocatableToJSON(x *Locatable) ([]byte, error) {
	return x.MarshalJSON()
}

var (
	_ json.Unmarshaler = (*Locatable)(nil)
	_ json.Marshaler   = (*Locatable)(nil)
)

func (r *Locatable) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONLocatable(*r)
}
func (r *Locatable) _marshalJSONLocatable(x Locatable) ([]byte, error) {
	partial := make(map[string]json.RawMessage)
	var err error
	var fieldLocation []byte
	fieldLocation, err = r._marshalJSONstring(x.Location)
	if err != nil {
		return nil, fmt.Errorf("predicate: Locatable._marshalJSONLocatable: field name Location; %w", err)
	}
	partial["Location"] = fieldLocation
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("predicate: Locatable._marshalJSONLocatable: struct; %w", err)
	}
	return result, nil
}
func (r *Locatable) _marshalJSONstring(x string) ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, fmt.Errorf("predicate: Locatable._marshalJSONstring:; %w", err)
	}
	return result, nil
}
func (r *Locatable) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONLocatable(data)
	if err != nil {
		return fmt.Errorf("predicate: Locatable.UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *Locatable) _unmarshalJSONLocatable(data []byte) (Locatable, error) {
	result := Locatable{}
	var partial map[string]json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("predicate: Locatable._unmarshalJSONLocatable: native struct unwrap; %w", err)
	}
	if fieldLocation, ok := partial["Location"]; ok {
		result.Location, err = r._unmarshalJSONstring(fieldLocation)
		if err != nil {
			return result, fmt.Errorf("predicate: Locatable._unmarshalJSONLocatable: field Location; %w", err)
		}
	}
	return result, nil
}
func (r *Locatable) _unmarshalJSONstring(data []byte) (string, error) {
	var result string
	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, fmt.Errorf("predicate: Locatable._unmarshalJSONstring: native string unwrap; %w", err)
	}
	return result, nil
}
