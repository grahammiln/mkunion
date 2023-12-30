// Code generated by mkunion. DO NOT EDIT.
package predicate

import "github.com/widmogrod/mkunion/f"
import "github.com/widmogrod/mkunion/x/shape"
import "github.com/widmogrod/mkunion/x/shared"
import "encoding/json"
import "fmt"
import "github.com/widmogrod/mkunion/x/schema"

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
func init() {
	shape.Register(BindableShape())
	shape.Register(BindValueShape())
	shape.Register(LiteralShape())
	shape.Register(LocatableShape())
}

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
					IsPointer:     false,
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
					IsPointer:     false,
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
	var result *BindValue = new(BindValue)
	// if is Struct
	err := shared.JSONParseObject(x, func(key string, value []byte) error {
		switch key {
		case "BindName":
			return json.Unmarshal(value, &result.BindName)
		}

		return fmt.Errorf("predicate.BindValueFromJSON: unknown key %s", key)
	})

	return result, err
}

func BindValueToJSON(x *BindValue) ([]byte, error) {
	var err error
	field_BindName, err := shared.JSONMarshal[BindName](x.BindName)
	if err != nil {
		return nil, err
	}
	return json.Marshal(map[string]json.RawMessage{
		"BindName": field_BindName,
	})
}
func (self *BindValue) MarshalJSON() ([]byte, error) {
	return BindValueToJSON(self)
}

func (self *BindValue) UnmarshalJSON(x []byte) error {
	n, err := BindValueFromJSON(x)
	if err != nil {
		return err
	}
	*self = *n
	return nil
}

func LiteralFromJSON(x []byte) (*Literal, error) {
	var result *Literal = new(Literal)
	// if is Struct
	err := shared.JSONParseObject(x, func(key string, value []byte) error {
		switch key {
		case "Value":
			res, err := schema.SchemaFromJSON(value)
			if err != nil {
				return fmt.Errorf("predicate._FromJSON: field Schema %w", err)
			}
			result.Value = res
			return nil
		}

		return fmt.Errorf("predicate.LiteralFromJSON: unknown key %s", key)
	})

	return result, err
}

func LiteralToJSON(x *Literal) ([]byte, error) {
	var err error
	field_Value, err := schema.SchemaToJSON(x.Value)
	if err != nil {
		return nil, err
	}
	return json.Marshal(map[string]json.RawMessage{
		"Value": field_Value,
	})
}
func (self *Literal) MarshalJSON() ([]byte, error) {
	return LiteralToJSON(self)
}

func (self *Literal) UnmarshalJSON(x []byte) error {
	n, err := LiteralFromJSON(x)
	if err != nil {
		return err
	}
	*self = *n
	return nil
}

func LocatableFromJSON(x []byte) (*Locatable, error) {
	var result *Locatable = new(Locatable)
	// if is Struct
	err := shared.JSONParseObject(x, func(key string, value []byte) error {
		switch key {
		case "Location":
			return json.Unmarshal(value, &result.Location)
		}

		return fmt.Errorf("predicate.LocatableFromJSON: unknown key %s", key)
	})

	return result, err
}

func LocatableToJSON(x *Locatable) ([]byte, error) {
	var err error
	field_Location, err := json.Marshal(x.Location)
	if err != nil {
		return nil, err
	}
	return json.Marshal(map[string]json.RawMessage{
		"Location": field_Location,
	})
}
func (self *Locatable) MarshalJSON() ([]byte, error) {
	return LocatableToJSON(self)
}

func (self *Locatable) UnmarshalJSON(x []byte) error {
	n, err := LocatableFromJSON(x)
	if err != nil {
		return err
	}
	*self = *n
	return nil
}
