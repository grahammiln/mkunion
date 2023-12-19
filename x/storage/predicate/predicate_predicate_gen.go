// Code generated by mkunion. DO NOT EDIT.
package predicate

import "github.com/widmogrod/mkunion/f"
import "github.com/widmogrod/mkunion/x/shape"
import "github.com/widmogrod/mkunion/x/shared"
import "encoding/json"
import "fmt"

//mkunion-extension:visitor

type PredicateVisitor interface {
	VisitAnd(v *And) any
	VisitOr(v *Or) any
	VisitNot(v *Not) any
	VisitCompare(v *Compare) any
}

type Predicate interface {
	AcceptPredicate(g PredicateVisitor) any
}

func (r *And) AcceptPredicate(v PredicateVisitor) any     { return v.VisitAnd(r) }
func (r *Or) AcceptPredicate(v PredicateVisitor) any      { return v.VisitOr(r) }
func (r *Not) AcceptPredicate(v PredicateVisitor) any     { return v.VisitNot(r) }
func (r *Compare) AcceptPredicate(v PredicateVisitor) any { return v.VisitCompare(r) }

var (
	_ Predicate = (*And)(nil)
	_ Predicate = (*Or)(nil)
	_ Predicate = (*Not)(nil)
	_ Predicate = (*Compare)(nil)
)

func MatchPredicate[TOut any](
	x Predicate,
	f1 func(x *And) TOut,
	f2 func(x *Or) TOut,
	f3 func(x *Not) TOut,
	f4 func(x *Compare) TOut,
	df func(x Predicate) TOut,
) TOut {
	return f.Match4(x, f1, f2, f3, f4, df)
}

func MatchPredicateR2[TOut1, TOut2 any](
	x Predicate,
	f1 func(x *And) (TOut1, TOut2),
	f2 func(x *Or) (TOut1, TOut2),
	f3 func(x *Not) (TOut1, TOut2),
	f4 func(x *Compare) (TOut1, TOut2),
	df func(x Predicate) (TOut1, TOut2),
) (TOut1, TOut2) {
	return f.Match4R2(x, f1, f2, f3, f4, df)
}

func MustMatchPredicate[TOut any](
	x Predicate,
	f1 func(x *And) TOut,
	f2 func(x *Or) TOut,
	f3 func(x *Not) TOut,
	f4 func(x *Compare) TOut,
) TOut {
	return f.MustMatch4(x, f1, f2, f3, f4)
}

func MustMatchPredicateR0(
	x Predicate,
	f1 func(x *And),
	f2 func(x *Or),
	f3 func(x *Not),
	f4 func(x *Compare),
) {
	f.MustMatch4R0(x, f1, f2, f3, f4)
}

func MustMatchPredicateR2[TOut1, TOut2 any](
	x Predicate,
	f1 func(x *And) (TOut1, TOut2),
	f2 func(x *Or) (TOut1, TOut2),
	f3 func(x *Not) (TOut1, TOut2),
	f4 func(x *Compare) (TOut1, TOut2),
) (TOut1, TOut2) {
	return f.MustMatch4R2(x, f1, f2, f3, f4)
}

// mkunion-extension:shape
func PredicateShape() shape.Shape {
	return &shape.UnionLike{
		Name:          "Predicate",
		PkgName:       "predicate",
		PkgImportName: "github.com/widmogrod/mkunion/x/storage/predicate",
		Variant: []shape.Shape{
			AndShape(),
			OrShape(),
			NotShape(),
			CompareShape(),
		},
	}
}

func AndShape() shape.Shape {
	return &shape.StructLike{
		Name:          "And",
		PkgName:       "predicate",
		PkgImportName: "github.com/widmogrod/mkunion/x/storage/predicate",
		Fields: []*shape.FieldLike{
			{
				Name: "L",
				Type: &shape.ListLike{
					Element: &shape.RefName{
						Name:          "Predicate",
						PkgName:       "predicate",
						PkgImportName: "github.com/widmogrod/mkunion/x/storage/predicate",
						IsPointer:     false,
					},
					ElementIsPointer: false,
				},
			},
		},
	}
}

func OrShape() shape.Shape {
	return &shape.StructLike{
		Name:          "Or",
		PkgName:       "predicate",
		PkgImportName: "github.com/widmogrod/mkunion/x/storage/predicate",
		Fields: []*shape.FieldLike{
			{
				Name: "L",
				Type: &shape.ListLike{
					Element: &shape.RefName{
						Name:          "Predicate",
						PkgName:       "predicate",
						PkgImportName: "github.com/widmogrod/mkunion/x/storage/predicate",
						IsPointer:     false,
					},
					ElementIsPointer: false,
				},
			},
		},
	}
}

func NotShape() shape.Shape {
	return &shape.StructLike{
		Name:          "Not",
		PkgName:       "predicate",
		PkgImportName: "github.com/widmogrod/mkunion/x/storage/predicate",
		Fields: []*shape.FieldLike{
			{
				Name: "P",
				Type: &shape.RefName{
					Name:          "Predicate",
					PkgName:       "predicate",
					PkgImportName: "github.com/widmogrod/mkunion/x/storage/predicate",
					IsPointer:     false,
				},
			},
		},
	}
}

func CompareShape() shape.Shape {
	return &shape.StructLike{
		Name:          "Compare",
		PkgName:       "predicate",
		PkgImportName: "github.com/widmogrod/mkunion/x/storage/predicate",
		Fields: []*shape.FieldLike{
			{
				Name: "Location",
				Type: &shape.StringLike{},
			},
			{
				Name: "Operation",
				Type: &shape.StringLike{},
			},
			{
				Name: "BindValue",
				Type: &shape.RefName{
					Name:          "Bindable",
					PkgName:       "predicate",
					PkgImportName: "github.com/widmogrod/mkunion/x/storage/predicate",
					IsPointer:     false,
				},
			},
		},
	}
}

// mkunion-extension:json
func init() {
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/storage/predicate.Predicate", PredicateFromJSON, PredicateToJSON)
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/storage/predicate.And", AndFromJSON, AndToJSON)
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/storage/predicate.Or", OrFromJSON, OrToJSON)
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/storage/predicate.Not", NotFromJSON, NotToJSON)
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/storage/predicate.Compare", CompareFromJSON, CompareToJSON)
}

type PredicateUnionJSON struct {
	Type    string          `json:"$type,omitempty"`
	And     json.RawMessage `json:"predicate.And,omitempty"`
	Or      json.RawMessage `json:"predicate.Or,omitempty"`
	Not     json.RawMessage `json:"predicate.Not,omitempty"`
	Compare json.RawMessage `json:"predicate.Compare,omitempty"`
}

func PredicateFromJSON(x []byte) (Predicate, error) {
	var data PredicateUnionJSON
	err := json.Unmarshal(x, &data)
	if err != nil {
		return nil, err
	}

	switch data.Type {
	case "predicate.And":
		return AndFromJSON(data.And)
	case "predicate.Or":
		return OrFromJSON(data.Or)
	case "predicate.Not":
		return NotFromJSON(data.Not)
	case "predicate.Compare":
		return CompareFromJSON(data.Compare)
	}

	if data.And != nil {
		return AndFromJSON(data.And)
	} else if data.Or != nil {
		return OrFromJSON(data.Or)
	} else if data.Not != nil {
		return NotFromJSON(data.Not)
	} else if data.Compare != nil {
		return CompareFromJSON(data.Compare)
	}

	return nil, fmt.Errorf("predicate.Predicate: unknown type %s", data.Type)
}

func PredicateToJSON(x Predicate) ([]byte, error) {
	if x == nil {
		return nil, nil
	}
	return MustMatchPredicateR2(
		x,
		func(x *And) ([]byte, error) {
			body, err := AndToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(PredicateUnionJSON{
				Type: "predicate.And",
				And:  body,
			})
		},
		func(x *Or) ([]byte, error) {
			body, err := OrToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(PredicateUnionJSON{
				Type: "predicate.Or",
				Or:   body,
			})
		},
		func(x *Not) ([]byte, error) {
			body, err := NotToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(PredicateUnionJSON{
				Type: "predicate.Not",
				Not:  body,
			})
		},
		func(x *Compare) ([]byte, error) {
			body, err := CompareToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(PredicateUnionJSON{
				Type:    "predicate.Compare",
				Compare: body,
			})
		},
	)
}

func AndFromJSON(x []byte) (*And, error) {
	var result *And = new(And)
	// if is Struct
	err := shared.JSONParseObject(x, func(key string, value []byte) error {
		switch key {
		case "L":
			res, err := shared.JSONToListWithDeserializer(value, result.L, PredicateFromJSON)
			if err != nil {
				return fmt.Errorf("predicate._FromJSON: field Predicate %w", err)
			}
			result.L = res
			return nil
		}

		return fmt.Errorf("predicate.AndFromJSON: unknown key %s", key)
	})

	return result, err
}

func AndToJSON(x *And) ([]byte, error) {
	field_L, err := shared.JSONListFromSerializer(x.L, PredicateToJSON)
	if err != nil {
		return nil, err
	}
	return json.Marshal(map[string]json.RawMessage{
		"L": field_L,
	})
}
func (self *And) MarshalJSON() ([]byte, error) {
	return AndToJSON(self)
}

func (self *And) UnmarshalJSON(x []byte) error {
	n, err := AndFromJSON(x)
	if err != nil {
		return err
	}
	*self = *n
	return nil
}

func OrFromJSON(x []byte) (*Or, error) {
	var result *Or = new(Or)
	// if is Struct
	err := shared.JSONParseObject(x, func(key string, value []byte) error {
		switch key {
		case "L":
			res, err := shared.JSONToListWithDeserializer(value, result.L, PredicateFromJSON)
			if err != nil {
				return fmt.Errorf("predicate._FromJSON: field Predicate %w", err)
			}
			result.L = res
			return nil
		}

		return fmt.Errorf("predicate.OrFromJSON: unknown key %s", key)
	})

	return result, err
}

func OrToJSON(x *Or) ([]byte, error) {
	field_L, err := shared.JSONListFromSerializer(x.L, PredicateToJSON)
	if err != nil {
		return nil, err
	}
	return json.Marshal(map[string]json.RawMessage{
		"L": field_L,
	})
}
func (self *Or) MarshalJSON() ([]byte, error) {
	return OrToJSON(self)
}

func (self *Or) UnmarshalJSON(x []byte) error {
	n, err := OrFromJSON(x)
	if err != nil {
		return err
	}
	*self = *n
	return nil
}

func NotFromJSON(x []byte) (*Not, error) {
	var result *Not = new(Not)
	// if is Struct
	err := shared.JSONParseObject(x, func(key string, value []byte) error {
		switch key {
		case "P":
			res, err := PredicateFromJSON(value)
			if err != nil {
				return fmt.Errorf("predicate._FromJSON: field Predicate %w", err)
			}
			result.P = res
			return nil
		}

		return fmt.Errorf("predicate.NotFromJSON: unknown key %s", key)
	})

	return result, err
}

func NotToJSON(x *Not) ([]byte, error) {
	field_P, err := PredicateToJSON(x.P)
	if err != nil {
		return nil, err
	}
	return json.Marshal(map[string]json.RawMessage{
		"P": field_P,
	})
}
func (self *Not) MarshalJSON() ([]byte, error) {
	return NotToJSON(self)
}

func (self *Not) UnmarshalJSON(x []byte) error {
	n, err := NotFromJSON(x)
	if err != nil {
		return err
	}
	*self = *n
	return nil
}

func CompareFromJSON(x []byte) (*Compare, error) {
	var result *Compare = new(Compare)
	// if is Struct
	err := shared.JSONParseObject(x, func(key string, value []byte) error {
		switch key {
		case "Location":
			return json.Unmarshal(value, &result.Location)
		case "Operation":
			return json.Unmarshal(value, &result.Operation)
		case "BindValue":
			res, err := BindableFromJSON(value)
			if err != nil {
				return fmt.Errorf("predicate._FromJSON: field Bindable %w", err)
			}
			result.BindValue = res
			return nil
		}

		return fmt.Errorf("predicate.CompareFromJSON: unknown key %s", key)
	})

	return result, err
}

func CompareToJSON(x *Compare) ([]byte, error) {
	field_Location, err := json.Marshal(x.Location)
	if err != nil {
		return nil, err
	}
	field_Operation, err := json.Marshal(x.Operation)
	if err != nil {
		return nil, err
	}
	field_BindValue, err := BindableToJSON(x.BindValue)
	if err != nil {
		return nil, err
	}
	return json.Marshal(map[string]json.RawMessage{
		"Location":  field_Location,
		"Operation": field_Operation,
		"BindValue": field_BindValue,
	})
}
func (self *Compare) MarshalJSON() ([]byte, error) {
	return CompareToJSON(self)
}

func (self *Compare) UnmarshalJSON(x []byte) error {
	n, err := CompareFromJSON(x)
	if err != nil {
		return err
	}
	*self = *n
	return nil
}
