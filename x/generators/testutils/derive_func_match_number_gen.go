// Code generated by mkunion. DO NOT EDIT.
package testutils

import "github.com/widmogrod/mkunion/f"
import "github.com/widmogrod/mkunion/x/schema"
import "github.com/widmogrod/mkunion/x/shape"
import "github.com/widmogrod/mkunion/x/shared"
import "encoding/json"
import "fmt"

//mkunion-extension:visitor

type NumberVisitor interface {
	VisitN0(v *N0) any
	VisitN1(v *N1) any
}

type Number interface {
	AcceptNumber(g NumberVisitor) any
}

func (r *N0) AcceptNumber(v NumberVisitor) any { return v.VisitN0(r) }
func (r *N1) AcceptNumber(v NumberVisitor) any { return v.VisitN1(r) }

var (
	_ Number = (*N0)(nil)
	_ Number = (*N1)(nil)
)

func MatchNumber[TOut any](
	x Number,
	f1 func(x *N0) TOut,
	f2 func(x *N1) TOut,
	df func(x Number) TOut,
) TOut {
	return f.Match2(x, f1, f2, df)
}

func MatchNumberR2[TOut1, TOut2 any](
	x Number,
	f1 func(x *N0) (TOut1, TOut2),
	f2 func(x *N1) (TOut1, TOut2),
	df func(x Number) (TOut1, TOut2),
) (TOut1, TOut2) {
	return f.Match2R2(x, f1, f2, df)
}

func MustMatchNumber[TOut any](
	x Number,
	f1 func(x *N0) TOut,
	f2 func(x *N1) TOut,
) TOut {
	return f.MustMatch2(x, f1, f2)
}

func MustMatchNumberR0(
	x Number,
	f1 func(x *N0),
	f2 func(x *N1),
) {
	f.MustMatch2R0(x, f1, f2)
}

func MustMatchNumberR2[TOut1, TOut2 any](
	x Number,
	f1 func(x *N0) (TOut1, TOut2),
	f2 func(x *N1) (TOut1, TOut2),
) (TOut1, TOut2) {
	return f.MustMatch2R2(x, f1, f2)
}

// mkunion-extension:schema
func init() {
	schema.RegisterUnionTypes(NumberSchemaDef())
}

func NumberSchemaDef() *schema.UnionVariants[Number] {
	return schema.MustDefineUnion[Number](
		new(N0),
		new(N1),
	)
}

// mkunion-extension:shape
func NumberShape() shape.Shape {
	return &shape.UnionLike{
		Name:          "Number",
		PkgName:       "testutils",
		PkgImportName: "github.com/widmogrod/mkunion/x/generators/testutils",
		Variant: []shape.Shape{
			N0Shape(),
			N1Shape(),
		},
	}
}

func N0Shape() shape.Shape {
	return &shape.StructLike{
		Name:          "N0",
		PkgName:       "testutils",
		PkgImportName: "github.com/widmogrod/mkunion/x/generators/testutils",
	}
}

func N1Shape() shape.Shape {
	return &shape.StructLike{
		Name:          "N1",
		PkgName:       "testutils",
		PkgImportName: "github.com/widmogrod/mkunion/x/generators/testutils",
	}
}

// mkunion-extension:json
func init() {
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/generators/testutils.Number", NumberFromJSON, NumberToJSON)
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/generators/testutils.N0", N0FromJSON, N0ToJSON)
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/generators/testutils.N1", N1FromJSON, N1ToJSON)
}

type NumberUnionJSON struct {
	Type string          `json:"$type,omitempty"`
	N0   json.RawMessage `json:"testutils.N0,omitempty"`
	N1   json.RawMessage `json:"testutils.N1,omitempty"`
}

func NumberFromJSON(x []byte) (Number, error) {
	var data NumberUnionJSON
	err := json.Unmarshal(x, &data)
	if err != nil {
		return nil, err
	}

	switch data.Type {
	case "testutils.N0":
		return N0FromJSON(data.N0)
	case "testutils.N1":
		return N1FromJSON(data.N1)
	}

	if data.N0 != nil {
		return N0FromJSON(data.N0)
	} else if data.N1 != nil {
		return N1FromJSON(data.N1)
	}

	return nil, fmt.Errorf("testutils.Number: unknown type %s", data.Type)
}

func NumberToJSON(x Number) ([]byte, error) {
	if x == nil {
		return nil, nil
	}
	return MustMatchNumberR2(
		x,
		func(x *N0) ([]byte, error) {
			body, err := N0ToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(NumberUnionJSON{
				Type: "testutils.N0",
				N0:   body,
			})
		},
		func(x *N1) ([]byte, error) {
			body, err := N1ToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(NumberUnionJSON{
				Type: "testutils.N1",
				N1:   body,
			})
		},
	)
}

func N0FromJSON(x []byte) (*N0, error) {
	var result *N0 = new(N0)
	// if is Struct
	err := shared.JSONParseObject(x, func(key string, value []byte) error {
		switch key {
		}

		return fmt.Errorf("testutils.N0FromJSON: unknown key %s", key)
	})

	return result, err
}

func N0ToJSON(x *N0) ([]byte, error) {
	return json.Marshal(map[string]json.RawMessage{})
}
func (self *N0) MarshalJSON() ([]byte, error) {
	return N0ToJSON(self)
}

func (self *N0) UnmarshalJSON(x []byte) error {
	n, err := N0FromJSON(x)
	if err != nil {
		return err
	}
	*self = *n
	return nil
}

func N1FromJSON(x []byte) (*N1, error) {
	var result *N1 = new(N1)
	// if is Struct
	err := shared.JSONParseObject(x, func(key string, value []byte) error {
		switch key {
		}

		return fmt.Errorf("testutils.N1FromJSON: unknown key %s", key)
	})

	return result, err
}

func N1ToJSON(x *N1) ([]byte, error) {
	return json.Marshal(map[string]json.RawMessage{})
}
func (self *N1) MarshalJSON() ([]byte, error) {
	return N1ToJSON(self)
}

func (self *N1) UnmarshalJSON(x []byte) error {
	n, err := N1FromJSON(x)
	if err != nil {
		return err
	}
	*self = *n
	return nil
}
