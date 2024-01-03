// Code generated by mkunion. DO NOT EDIT.
package testutils

import (
	"encoding/json"
	"fmt"
	"github.com/widmogrod/mkunion/f"
	"github.com/widmogrod/mkunion/x/shape"
	"github.com/widmogrod/mkunion/x/shared"
)

func init() {
	shape.Register(AlphabetShape())
	shape.Register(A1Shape())
	shape.Register(B2Shape())
	shape.Register(C3Shape())
}

//mkunion-extension:visitor

type AlphabetVisitor interface {
	VisitA1(v *A1) any
	VisitB2(v *B2) any
	VisitC3(v *C3) any
}

type Alphabet interface {
	AcceptAlphabet(g AlphabetVisitor) any
}

func (r *A1) AcceptAlphabet(v AlphabetVisitor) any { return v.VisitA1(r) }
func (r *B2) AcceptAlphabet(v AlphabetVisitor) any { return v.VisitB2(r) }
func (r *C3) AcceptAlphabet(v AlphabetVisitor) any { return v.VisitC3(r) }

var (
	_ Alphabet = (*A1)(nil)
	_ Alphabet = (*B2)(nil)
	_ Alphabet = (*C3)(nil)
)

func MatchAlphabet[TOut any](
	x Alphabet,
	f1 func(x *A1) TOut,
	f2 func(x *B2) TOut,
	f3 func(x *C3) TOut,
	df func(x Alphabet) TOut,
) TOut {
	return f.Match3(x, f1, f2, f3, df)
}

func MatchAlphabetR2[TOut1, TOut2 any](
	x Alphabet,
	f1 func(x *A1) (TOut1, TOut2),
	f2 func(x *B2) (TOut1, TOut2),
	f3 func(x *C3) (TOut1, TOut2),
	df func(x Alphabet) (TOut1, TOut2),
) (TOut1, TOut2) {
	return f.Match3R2(x, f1, f2, f3, df)
}

func MustMatchAlphabet[TOut any](
	x Alphabet,
	f1 func(x *A1) TOut,
	f2 func(x *B2) TOut,
	f3 func(x *C3) TOut,
) TOut {
	return f.MustMatch3(x, f1, f2, f3)
}

func MustMatchAlphabetR0(
	x Alphabet,
	f1 func(x *A1),
	f2 func(x *B2),
	f3 func(x *C3),
) {
	f.MustMatch3R0(x, f1, f2, f3)
}

func MustMatchAlphabetR2[TOut1, TOut2 any](
	x Alphabet,
	f1 func(x *A1) (TOut1, TOut2),
	f2 func(x *B2) (TOut1, TOut2),
	f3 func(x *C3) (TOut1, TOut2),
) (TOut1, TOut2) {
	return f.MustMatch3R2(x, f1, f2, f3)
}

//mkunion-extension:shape

func AlphabetShape() shape.Shape {
	return &shape.UnionLike{
		Name:          "Alphabet",
		PkgName:       "testutils",
		PkgImportName: "github.com/widmogrod/mkunion/x/generators/testutils",
		Variant: []shape.Shape{
			A1Shape(),
			B2Shape(),
			C3Shape(),
		},
	}
}

func A1Shape() shape.Shape {
	return &shape.StructLike{
		Name:          "A1",
		PkgName:       "testutils",
		PkgImportName: "github.com/widmogrod/mkunion/x/generators/testutils",
	}
}

func B2Shape() shape.Shape {
	return &shape.StructLike{
		Name:          "B2",
		PkgName:       "testutils",
		PkgImportName: "github.com/widmogrod/mkunion/x/generators/testutils",
	}
}

func C3Shape() shape.Shape {
	return &shape.StructLike{
		Name:          "C3",
		PkgName:       "testutils",
		PkgImportName: "github.com/widmogrod/mkunion/x/generators/testutils",
	}
}

// mkunion-extension:json
func init() {
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/generators/testutils.Alphabet", AlphabetFromJSON, AlphabetToJSON)
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/generators/testutils.A1", A1FromJSON, A1ToJSON)
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/generators/testutils.B2", B2FromJSON, B2ToJSON)
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/generators/testutils.C3", C3FromJSON, C3ToJSON)
}

type AlphabetUnionJSON struct {
	Type string          `json:"$type,omitempty"`
	A1   json.RawMessage `json:"testutils.A1,omitempty"`
	B2   json.RawMessage `json:"testutils.B2,omitempty"`
	C3   json.RawMessage `json:"testutils.C3,omitempty"`
}

func AlphabetFromJSON(x []byte) (Alphabet, error) {
	if x == nil || len(x) == 0 {
		return nil, nil
	}
	if string(x[:4]) == "null" {
		return nil, nil
	}

	var data AlphabetUnionJSON
	err := json.Unmarshal(x, &data)
	if err != nil {
		return nil, err
	}

	switch data.Type {
	case "testutils.A1":
		return A1FromJSON(data.A1)
	case "testutils.B2":
		return B2FromJSON(data.B2)
	case "testutils.C3":
		return C3FromJSON(data.C3)
	}

	if data.A1 != nil {
		return A1FromJSON(data.A1)
	} else if data.B2 != nil {
		return B2FromJSON(data.B2)
	} else if data.C3 != nil {
		return C3FromJSON(data.C3)
	}

	return nil, fmt.Errorf("testutils.Alphabet: unknown type %s", data.Type)
}

func AlphabetToJSON(x Alphabet) ([]byte, error) {
	if x == nil {
		return nil, nil
	}
	return MustMatchAlphabetR2(
		x,
		func(x *A1) ([]byte, error) {
			body, err := A1ToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(AlphabetUnionJSON{
				Type: "testutils.A1",
				A1:   body,
			})
		},
		func(x *B2) ([]byte, error) {
			body, err := B2ToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(AlphabetUnionJSON{
				Type: "testutils.B2",
				B2:   body,
			})
		},
		func(x *C3) ([]byte, error) {
			body, err := C3ToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(AlphabetUnionJSON{
				Type: "testutils.C3",
				C3:   body,
			})
		},
	)
}

func A1FromJSON(x []byte) (*A1, error) {
	result := new(A1)
	err := result.UnmarshalJSON(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func A1ToJSON(x *A1) ([]byte, error) {
	return x.MarshalJSON()
}

var (
	_ json.Unmarshaler = (*A1)(nil)
	_ json.Marshaler   = (*A1)(nil)
)

func (r *A1) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONA1(*r)
}
func (r *A1) _marshalJSONA1(x A1) ([]byte, error) {
	partial := make(map[string]json.RawMessage)
	var err error
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("testutils: A1._marshalJSONA1: struct; %w", err)
	}
	return result, nil
}
func (r *A1) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONA1(data)
	if err != nil {
		return fmt.Errorf("testutils: A1.UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *A1) _unmarshalJSONA1(data []byte) (A1, error) {
	result := A1{}
	var partial map[string]json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("testutils: A1._unmarshalJSONA1: native struct unwrap; %w", err)
	}
	return result, nil
}

func B2FromJSON(x []byte) (*B2, error) {
	result := new(B2)
	err := result.UnmarshalJSON(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func B2ToJSON(x *B2) ([]byte, error) {
	return x.MarshalJSON()
}

var (
	_ json.Unmarshaler = (*B2)(nil)
	_ json.Marshaler   = (*B2)(nil)
)

func (r *B2) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONB2(*r)
}
func (r *B2) _marshalJSONB2(x B2) ([]byte, error) {
	partial := make(map[string]json.RawMessage)
	var err error
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("testutils: B2._marshalJSONB2: struct; %w", err)
	}
	return result, nil
}
func (r *B2) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONB2(data)
	if err != nil {
		return fmt.Errorf("testutils: B2.UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *B2) _unmarshalJSONB2(data []byte) (B2, error) {
	result := B2{}
	var partial map[string]json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("testutils: B2._unmarshalJSONB2: native struct unwrap; %w", err)
	}
	return result, nil
}

func C3FromJSON(x []byte) (*C3, error) {
	result := new(C3)
	err := result.UnmarshalJSON(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func C3ToJSON(x *C3) ([]byte, error) {
	return x.MarshalJSON()
}

var (
	_ json.Unmarshaler = (*C3)(nil)
	_ json.Marshaler   = (*C3)(nil)
)

func (r *C3) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONC3(*r)
}
func (r *C3) _marshalJSONC3(x C3) ([]byte, error) {
	partial := make(map[string]json.RawMessage)
	var err error
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("testutils: C3._marshalJSONC3: struct; %w", err)
	}
	return result, nil
}
func (r *C3) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONC3(data)
	if err != nil {
		return fmt.Errorf("testutils: C3.UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *C3) _unmarshalJSONC3(data []byte) (C3, error) {
	result := C3{}
	var partial map[string]json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("testutils: C3._unmarshalJSONC3: native struct unwrap; %w", err)
	}
	return result, nil
}
