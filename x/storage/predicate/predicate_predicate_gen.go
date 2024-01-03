// Code generated by mkunion. DO NOT EDIT.
package predicate

import (
	"encoding/json"
	"fmt"
	"github.com/widmogrod/mkunion/f"
	"github.com/widmogrod/mkunion/x/shape"
	"github.com/widmogrod/mkunion/x/shared"
)

func init() {
	shape.Register(PredicateShape())
	shape.Register(AndShape())
	shape.Register(OrShape())
	shape.Register(NotShape())
	shape.Register(CompareShape())
}

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

//mkunion-extension:shape

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
					},
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
					},
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
	if x == nil || len(x) == 0 {
		return nil, nil
	}
	if string(x[:4]) == "null" {
		return nil, nil
	}

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
	result := new(And)
	err := result.UnmarshalJSON(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func AndToJSON(x *And) ([]byte, error) {
	return x.MarshalJSON()
}

var (
	_ json.Unmarshaler = (*And)(nil)
	_ json.Marshaler   = (*And)(nil)
)

func (r *And) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONAnd(*r)
}
func (r *And) _marshalJSONAnd(x And) ([]byte, error) {
	partial := make(map[string]json.RawMessage)
	var err error
	var fieldL []byte
	fieldL, err = r._marshalJSONSlicePredicate(x.L)
	if err != nil {
		return nil, fmt.Errorf("predicate: And._marshalJSONAnd: field name L; %w", err)
	}
	partial["L"] = fieldL
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("predicate: And._marshalJSONAnd: struct; %w", err)
	}
	return result, nil
}
func (r *And) _marshalJSONSlicePredicate(x []Predicate) ([]byte, error) {
	partial := make([]json.RawMessage, len(x))
	for i, v := range x {
		item, err := r._marshalJSONPredicate(v)
		if err != nil {
			return nil, fmt.Errorf("predicate: And._marshalJSONSlicePredicate: at index %d; %w", i, err)
		}
		partial[i] = item
	}
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("predicate: And._marshalJSONSlicePredicate:; %w", err)
	}
	return result, nil
}
func (r *And) _marshalJSONPredicate(x Predicate) ([]byte, error) {
	result, err := shared.JSONMarshal[Predicate](x)
	if err != nil {
		return nil, fmt.Errorf("predicate: And._marshalJSONPredicate:; %w", err)
	}
	return result, nil
}
func (r *And) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONAnd(data)
	if err != nil {
		return fmt.Errorf("predicate: And.UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *And) _unmarshalJSONAnd(data []byte) (And, error) {
	result := And{}
	var partial map[string]json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("predicate: And._unmarshalJSONAnd: native struct unwrap; %w", err)
	}
	if fieldL, ok := partial["L"]; ok {
		result.L, err = r._unmarshalJSONSlicePredicate(fieldL)
		if err != nil {
			return result, fmt.Errorf("predicate: And._unmarshalJSONAnd: field L; %w", err)
		}
	}
	return result, nil
}
func (r *And) _unmarshalJSONSlicePredicate(data []byte) ([]Predicate, error) {
	result := make([]Predicate, 0)
	var partial []json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("predicate: And._unmarshalJSONSlicePredicate: native list unwrap; %w", err)
	}
	for i, v := range partial {
		item, err := r._unmarshalJSONPredicate(v)
		if err != nil {
			return result, fmt.Errorf("predicate: And._unmarshalJSONSlicePredicate: at index %d; %w", i, err)
		}
		result = append(result, item)
	}
	return result, nil
}
func (r *And) _unmarshalJSONPredicate(data []byte) (Predicate, error) {
	result, err := shared.JSONUnmarshal[Predicate](data)
	if err != nil {
		return result, fmt.Errorf("predicate: And._unmarshalJSONPredicate: native ref unwrap; %w", err)
	}
	return result, nil
}

func OrFromJSON(x []byte) (*Or, error) {
	result := new(Or)
	err := result.UnmarshalJSON(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func OrToJSON(x *Or) ([]byte, error) {
	return x.MarshalJSON()
}

var (
	_ json.Unmarshaler = (*Or)(nil)
	_ json.Marshaler   = (*Or)(nil)
)

func (r *Or) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONOr(*r)
}
func (r *Or) _marshalJSONOr(x Or) ([]byte, error) {
	partial := make(map[string]json.RawMessage)
	var err error
	var fieldL []byte
	fieldL, err = r._marshalJSONSlicePredicate(x.L)
	if err != nil {
		return nil, fmt.Errorf("predicate: Or._marshalJSONOr: field name L; %w", err)
	}
	partial["L"] = fieldL
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("predicate: Or._marshalJSONOr: struct; %w", err)
	}
	return result, nil
}
func (r *Or) _marshalJSONSlicePredicate(x []Predicate) ([]byte, error) {
	partial := make([]json.RawMessage, len(x))
	for i, v := range x {
		item, err := r._marshalJSONPredicate(v)
		if err != nil {
			return nil, fmt.Errorf("predicate: Or._marshalJSONSlicePredicate: at index %d; %w", i, err)
		}
		partial[i] = item
	}
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("predicate: Or._marshalJSONSlicePredicate:; %w", err)
	}
	return result, nil
}
func (r *Or) _marshalJSONPredicate(x Predicate) ([]byte, error) {
	result, err := shared.JSONMarshal[Predicate](x)
	if err != nil {
		return nil, fmt.Errorf("predicate: Or._marshalJSONPredicate:; %w", err)
	}
	return result, nil
}
func (r *Or) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONOr(data)
	if err != nil {
		return fmt.Errorf("predicate: Or.UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *Or) _unmarshalJSONOr(data []byte) (Or, error) {
	result := Or{}
	var partial map[string]json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("predicate: Or._unmarshalJSONOr: native struct unwrap; %w", err)
	}
	if fieldL, ok := partial["L"]; ok {
		result.L, err = r._unmarshalJSONSlicePredicate(fieldL)
		if err != nil {
			return result, fmt.Errorf("predicate: Or._unmarshalJSONOr: field L; %w", err)
		}
	}
	return result, nil
}
func (r *Or) _unmarshalJSONSlicePredicate(data []byte) ([]Predicate, error) {
	result := make([]Predicate, 0)
	var partial []json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("predicate: Or._unmarshalJSONSlicePredicate: native list unwrap; %w", err)
	}
	for i, v := range partial {
		item, err := r._unmarshalJSONPredicate(v)
		if err != nil {
			return result, fmt.Errorf("predicate: Or._unmarshalJSONSlicePredicate: at index %d; %w", i, err)
		}
		result = append(result, item)
	}
	return result, nil
}
func (r *Or) _unmarshalJSONPredicate(data []byte) (Predicate, error) {
	result, err := shared.JSONUnmarshal[Predicate](data)
	if err != nil {
		return result, fmt.Errorf("predicate: Or._unmarshalJSONPredicate: native ref unwrap; %w", err)
	}
	return result, nil
}

func NotFromJSON(x []byte) (*Not, error) {
	result := new(Not)
	err := result.UnmarshalJSON(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func NotToJSON(x *Not) ([]byte, error) {
	return x.MarshalJSON()
}

var (
	_ json.Unmarshaler = (*Not)(nil)
	_ json.Marshaler   = (*Not)(nil)
)

func (r *Not) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONNot(*r)
}
func (r *Not) _marshalJSONNot(x Not) ([]byte, error) {
	partial := make(map[string]json.RawMessage)
	var err error
	var fieldP []byte
	fieldP, err = r._marshalJSONPredicate(x.P)
	if err != nil {
		return nil, fmt.Errorf("predicate: Not._marshalJSONNot: field name P; %w", err)
	}
	partial["P"] = fieldP
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("predicate: Not._marshalJSONNot: struct; %w", err)
	}
	return result, nil
}
func (r *Not) _marshalJSONPredicate(x Predicate) ([]byte, error) {
	result, err := shared.JSONMarshal[Predicate](x)
	if err != nil {
		return nil, fmt.Errorf("predicate: Not._marshalJSONPredicate:; %w", err)
	}
	return result, nil
}
func (r *Not) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONNot(data)
	if err != nil {
		return fmt.Errorf("predicate: Not.UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *Not) _unmarshalJSONNot(data []byte) (Not, error) {
	result := Not{}
	var partial map[string]json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("predicate: Not._unmarshalJSONNot: native struct unwrap; %w", err)
	}
	if fieldP, ok := partial["P"]; ok {
		result.P, err = r._unmarshalJSONPredicate(fieldP)
		if err != nil {
			return result, fmt.Errorf("predicate: Not._unmarshalJSONNot: field P; %w", err)
		}
	}
	return result, nil
}
func (r *Not) _unmarshalJSONPredicate(data []byte) (Predicate, error) {
	result, err := shared.JSONUnmarshal[Predicate](data)
	if err != nil {
		return result, fmt.Errorf("predicate: Not._unmarshalJSONPredicate: native ref unwrap; %w", err)
	}
	return result, nil
}

func CompareFromJSON(x []byte) (*Compare, error) {
	result := new(Compare)
	err := result.UnmarshalJSON(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func CompareToJSON(x *Compare) ([]byte, error) {
	return x.MarshalJSON()
}

var (
	_ json.Unmarshaler = (*Compare)(nil)
	_ json.Marshaler   = (*Compare)(nil)
)

func (r *Compare) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONCompare(*r)
}
func (r *Compare) _marshalJSONCompare(x Compare) ([]byte, error) {
	partial := make(map[string]json.RawMessage)
	var err error
	var fieldLocation []byte
	fieldLocation, err = r._marshalJSONstring(x.Location)
	if err != nil {
		return nil, fmt.Errorf("predicate: Compare._marshalJSONCompare: field name Location; %w", err)
	}
	partial["Location"] = fieldLocation
	var fieldOperation []byte
	fieldOperation, err = r._marshalJSONstring(x.Operation)
	if err != nil {
		return nil, fmt.Errorf("predicate: Compare._marshalJSONCompare: field name Operation; %w", err)
	}
	partial["Operation"] = fieldOperation
	var fieldBindValue []byte
	fieldBindValue, err = r._marshalJSONBindable(x.BindValue)
	if err != nil {
		return nil, fmt.Errorf("predicate: Compare._marshalJSONCompare: field name BindValue; %w", err)
	}
	partial["BindValue"] = fieldBindValue
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("predicate: Compare._marshalJSONCompare: struct; %w", err)
	}
	return result, nil
}
func (r *Compare) _marshalJSONstring(x string) ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, fmt.Errorf("predicate: Compare._marshalJSONstring:; %w", err)
	}
	return result, nil
}
func (r *Compare) _marshalJSONBindable(x Bindable) ([]byte, error) {
	result, err := shared.JSONMarshal[Bindable](x)
	if err != nil {
		return nil, fmt.Errorf("predicate: Compare._marshalJSONBindable:; %w", err)
	}
	return result, nil
}
func (r *Compare) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONCompare(data)
	if err != nil {
		return fmt.Errorf("predicate: Compare.UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *Compare) _unmarshalJSONCompare(data []byte) (Compare, error) {
	result := Compare{}
	var partial map[string]json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("predicate: Compare._unmarshalJSONCompare: native struct unwrap; %w", err)
	}
	if fieldLocation, ok := partial["Location"]; ok {
		result.Location, err = r._unmarshalJSONstring(fieldLocation)
		if err != nil {
			return result, fmt.Errorf("predicate: Compare._unmarshalJSONCompare: field Location; %w", err)
		}
	}
	if fieldOperation, ok := partial["Operation"]; ok {
		result.Operation, err = r._unmarshalJSONstring(fieldOperation)
		if err != nil {
			return result, fmt.Errorf("predicate: Compare._unmarshalJSONCompare: field Operation; %w", err)
		}
	}
	if fieldBindValue, ok := partial["BindValue"]; ok {
		result.BindValue, err = r._unmarshalJSONBindable(fieldBindValue)
		if err != nil {
			return result, fmt.Errorf("predicate: Compare._unmarshalJSONCompare: field BindValue; %w", err)
		}
	}
	return result, nil
}
func (r *Compare) _unmarshalJSONstring(data []byte) (string, error) {
	var result string
	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, fmt.Errorf("predicate: Compare._unmarshalJSONstring: native string unwrap; %w", err)
	}
	return result, nil
}
func (r *Compare) _unmarshalJSONBindable(data []byte) (Bindable, error) {
	result, err := shared.JSONUnmarshal[Bindable](data)
	if err != nil {
		return result, fmt.Errorf("predicate: Compare._unmarshalJSONBindable: native ref unwrap; %w", err)
	}
	return result, nil
}
