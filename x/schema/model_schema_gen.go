// Code generated by mkunion. DO NOT EDIT.
package schema

import (
	"encoding/json"
	"fmt"
	"github.com/widmogrod/mkunion/f"
	"github.com/widmogrod/mkunion/x/shape"
	"github.com/widmogrod/mkunion/x/shared"
)

func init() {
	shape.Register(SchemaShape())
	shape.Register(NoneShape())
	shape.Register(BoolShape())
	shape.Register(NumberShape())
	shape.Register(StringShape())
	shape.Register(BinaryShape())
	shape.Register(ListShape())
	shape.Register(MapShape())
}

//mkunion-extension:visitor

type SchemaVisitor interface {
	VisitNone(v *None) any
	VisitBool(v *Bool) any
	VisitNumber(v *Number) any
	VisitString(v *String) any
	VisitBinary(v *Binary) any
	VisitList(v *List) any
	VisitMap(v *Map) any
}

type Schema interface {
	AcceptSchema(g SchemaVisitor) any
}

func (r *None) AcceptSchema(v SchemaVisitor) any   { return v.VisitNone(r) }
func (r *Bool) AcceptSchema(v SchemaVisitor) any   { return v.VisitBool(r) }
func (r *Number) AcceptSchema(v SchemaVisitor) any { return v.VisitNumber(r) }
func (r *String) AcceptSchema(v SchemaVisitor) any { return v.VisitString(r) }
func (r *Binary) AcceptSchema(v SchemaVisitor) any { return v.VisitBinary(r) }
func (r *List) AcceptSchema(v SchemaVisitor) any   { return v.VisitList(r) }
func (r *Map) AcceptSchema(v SchemaVisitor) any    { return v.VisitMap(r) }

var (
	_ Schema = (*None)(nil)
	_ Schema = (*Bool)(nil)
	_ Schema = (*Number)(nil)
	_ Schema = (*String)(nil)
	_ Schema = (*Binary)(nil)
	_ Schema = (*List)(nil)
	_ Schema = (*Map)(nil)
)

func MatchSchema[TOut any](
	x Schema,
	f1 func(x *None) TOut,
	f2 func(x *Bool) TOut,
	f3 func(x *Number) TOut,
	f4 func(x *String) TOut,
	f5 func(x *Binary) TOut,
	f6 func(x *List) TOut,
	f7 func(x *Map) TOut,
	df func(x Schema) TOut,
) TOut {
	return f.Match7(x, f1, f2, f3, f4, f5, f6, f7, df)
}

func MatchSchemaR2[TOut1, TOut2 any](
	x Schema,
	f1 func(x *None) (TOut1, TOut2),
	f2 func(x *Bool) (TOut1, TOut2),
	f3 func(x *Number) (TOut1, TOut2),
	f4 func(x *String) (TOut1, TOut2),
	f5 func(x *Binary) (TOut1, TOut2),
	f6 func(x *List) (TOut1, TOut2),
	f7 func(x *Map) (TOut1, TOut2),
	df func(x Schema) (TOut1, TOut2),
) (TOut1, TOut2) {
	return f.Match7R2(x, f1, f2, f3, f4, f5, f6, f7, df)
}

func MustMatchSchema[TOut any](
	x Schema,
	f1 func(x *None) TOut,
	f2 func(x *Bool) TOut,
	f3 func(x *Number) TOut,
	f4 func(x *String) TOut,
	f5 func(x *Binary) TOut,
	f6 func(x *List) TOut,
	f7 func(x *Map) TOut,
) TOut {
	return f.MustMatch7(x, f1, f2, f3, f4, f5, f6, f7)
}

func MustMatchSchemaR0(
	x Schema,
	f1 func(x *None),
	f2 func(x *Bool),
	f3 func(x *Number),
	f4 func(x *String),
	f5 func(x *Binary),
	f6 func(x *List),
	f7 func(x *Map),
) {
	f.MustMatch7R0(x, f1, f2, f3, f4, f5, f6, f7)
}

func MustMatchSchemaR2[TOut1, TOut2 any](
	x Schema,
	f1 func(x *None) (TOut1, TOut2),
	f2 func(x *Bool) (TOut1, TOut2),
	f3 func(x *Number) (TOut1, TOut2),
	f4 func(x *String) (TOut1, TOut2),
	f5 func(x *Binary) (TOut1, TOut2),
	f6 func(x *List) (TOut1, TOut2),
	f7 func(x *Map) (TOut1, TOut2),
) (TOut1, TOut2) {
	return f.MustMatch7R2(x, f1, f2, f3, f4, f5, f6, f7)
}

//mkunion-extension:shape

func SchemaShape() shape.Shape {
	return &shape.UnionLike{
		Name:          "Schema",
		PkgName:       "schema",
		PkgImportName: "github.com/widmogrod/mkunion/x/schema",
		Variant: []shape.Shape{
			NoneShape(),
			BoolShape(),
			NumberShape(),
			StringShape(),
			BinaryShape(),
			ListShape(),
			MapShape(),
		},
	}
}

func NoneShape() shape.Shape {
	return &shape.StructLike{
		Name:          "None",
		PkgName:       "schema",
		PkgImportName: "github.com/widmogrod/mkunion/x/schema",
	}
}

func BoolShape() shape.Shape {
	return &shape.AliasLike{
		Name:          "Bool",
		PkgName:       "schema",
		PkgImportName: "github.com/widmogrod/mkunion/x/schema",
		Type:          &shape.BooleanLike{},
	}
}

func NumberShape() shape.Shape {
	return &shape.AliasLike{
		Name:          "Number",
		PkgName:       "schema",
		PkgImportName: "github.com/widmogrod/mkunion/x/schema",
		Type: &shape.NumberLike{
			Kind: &shape.Float64{},
		},
	}
}

func StringShape() shape.Shape {
	return &shape.AliasLike{
		Name:          "String",
		PkgName:       "schema",
		PkgImportName: "github.com/widmogrod/mkunion/x/schema",
		Type:          &shape.StringLike{},
	}
}

func BinaryShape() shape.Shape {
	return &shape.AliasLike{
		Name:          "Binary",
		PkgName:       "schema",
		PkgImportName: "github.com/widmogrod/mkunion/x/schema",
		Type: &shape.ListLike{
			Element: &shape.NumberLike{
				Kind: &shape.UInt8{},
			},
		},
	}
}

func ListShape() shape.Shape {
	return &shape.AliasLike{
		Name:          "List",
		PkgName:       "schema",
		PkgImportName: "github.com/widmogrod/mkunion/x/schema",
		Type: &shape.ListLike{
			Element: &shape.RefName{
				Name:          "Schema",
				PkgName:       "schema",
				PkgImportName: "github.com/widmogrod/mkunion/x/schema",
			},
		},
	}
}

func MapShape() shape.Shape {
	return &shape.AliasLike{
		Name:          "Map",
		PkgName:       "schema",
		PkgImportName: "github.com/widmogrod/mkunion/x/schema",
		Type: &shape.MapLike{
			Key: &shape.StringLike{},
			Val: &shape.RefName{
				Name:          "Schema",
				PkgName:       "schema",
				PkgImportName: "github.com/widmogrod/mkunion/x/schema",
			},
		},
	}
}

// mkunion-extension:json
func init() {
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/schema.Schema", SchemaFromJSON, SchemaToJSON)
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/schema.None", NoneFromJSON, NoneToJSON)
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/schema.Bool", BoolFromJSON, BoolToJSON)
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/schema.Number", NumberFromJSON, NumberToJSON)
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/schema.String", StringFromJSON, StringToJSON)
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/schema.Binary", BinaryFromJSON, BinaryToJSON)
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/schema.List", ListFromJSON, ListToJSON)
	shared.JSONMarshallerRegister("github.com/widmogrod/mkunion/x/schema.Map", MapFromJSON, MapToJSON)
}

type SchemaUnionJSON struct {
	Type   string          `json:"$type,omitempty"`
	None   json.RawMessage `json:"schema.None,omitempty"`
	Bool   json.RawMessage `json:"schema.Bool,omitempty"`
	Number json.RawMessage `json:"schema.Number,omitempty"`
	String json.RawMessage `json:"schema.String,omitempty"`
	Binary json.RawMessage `json:"schema.Binary,omitempty"`
	List   json.RawMessage `json:"schema.List,omitempty"`
	Map    json.RawMessage `json:"schema.Map,omitempty"`
}

func SchemaFromJSON(x []byte) (Schema, error) {
	if x == nil || len(x) == 0 {
		return nil, nil
	}
	if string(x[:4]) == "null" {
		return nil, nil
	}

	var data SchemaUnionJSON
	err := json.Unmarshal(x, &data)
	if err != nil {
		return nil, err
	}

	switch data.Type {
	case "schema.None":
		return NoneFromJSON(data.None)
	case "schema.Bool":
		return BoolFromJSON(data.Bool)
	case "schema.Number":
		return NumberFromJSON(data.Number)
	case "schema.String":
		return StringFromJSON(data.String)
	case "schema.Binary":
		return BinaryFromJSON(data.Binary)
	case "schema.List":
		return ListFromJSON(data.List)
	case "schema.Map":
		return MapFromJSON(data.Map)
	}

	if data.None != nil {
		return NoneFromJSON(data.None)
	} else if data.Bool != nil {
		return BoolFromJSON(data.Bool)
	} else if data.Number != nil {
		return NumberFromJSON(data.Number)
	} else if data.String != nil {
		return StringFromJSON(data.String)
	} else if data.Binary != nil {
		return BinaryFromJSON(data.Binary)
	} else if data.List != nil {
		return ListFromJSON(data.List)
	} else if data.Map != nil {
		return MapFromJSON(data.Map)
	}

	return nil, fmt.Errorf("schema.Schema: unknown type %s", data.Type)
}

func SchemaToJSON(x Schema) ([]byte, error) {
	if x == nil {
		return nil, nil
	}
	return MustMatchSchemaR2(
		x,
		func(x *None) ([]byte, error) {
			body, err := NoneToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(SchemaUnionJSON{
				Type: "schema.None",
				None: body,
			})
		},
		func(x *Bool) ([]byte, error) {
			body, err := BoolToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(SchemaUnionJSON{
				Type: "schema.Bool",
				Bool: body,
			})
		},
		func(x *Number) ([]byte, error) {
			body, err := NumberToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(SchemaUnionJSON{
				Type:   "schema.Number",
				Number: body,
			})
		},
		func(x *String) ([]byte, error) {
			body, err := StringToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(SchemaUnionJSON{
				Type:   "schema.String",
				String: body,
			})
		},
		func(x *Binary) ([]byte, error) {
			body, err := BinaryToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(SchemaUnionJSON{
				Type:   "schema.Binary",
				Binary: body,
			})
		},
		func(x *List) ([]byte, error) {
			body, err := ListToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(SchemaUnionJSON{
				Type: "schema.List",
				List: body,
			})
		},
		func(x *Map) ([]byte, error) {
			body, err := MapToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(SchemaUnionJSON{
				Type: "schema.Map",
				Map:  body,
			})
		},
	)
}

func NoneFromJSON(x []byte) (*None, error) {
	result := new(None)
	err := result.UnmarshalJSON(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func NoneToJSON(x *None) ([]byte, error) {
	return x.MarshalJSON()
}

var (
	_ json.Unmarshaler = (*None)(nil)
	_ json.Marshaler   = (*None)(nil)
)

func (r *None) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONNone(*r)
}
func (r *None) _marshalJSONNone(x None) ([]byte, error) {
	partial := make(map[string]json.RawMessage)
	var err error
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("schema: None._marshalJSONNone: struct; %w", err)
	}
	return result, nil
}
func (r *None) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONNone(data)
	if err != nil {
		return fmt.Errorf("schema: None.UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *None) _unmarshalJSONNone(data []byte) (None, error) {
	result := None{}
	var partial map[string]json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("schema: None._unmarshalJSONNone: native struct unwrap; %w", err)
	}
	return result, nil
}

func BoolFromJSON(x []byte) (*Bool, error) {
	result := new(Bool)
	err := result.UnmarshalJSON(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func BoolToJSON(x *Bool) ([]byte, error) {
	return x.MarshalJSON()
}

var (
	_ json.Unmarshaler = (*Bool)(nil)
	_ json.Marshaler   = (*Bool)(nil)
)

func (r *Bool) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONBool(*r)
}
func (r *Bool) _marshalJSONBool(x Bool) ([]byte, error) {
	return r._marshalJSONbool(bool(x))
}
func (r *Bool) _marshalJSONbool(x bool) ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, fmt.Errorf("schema: Bool._marshalJSONbool:; %w", err)
	}
	return result, nil
}
func (r *Bool) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONBool(data)
	if err != nil {
		return fmt.Errorf("schema: Bool.UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *Bool) _unmarshalJSONBool(data []byte) (Bool, error) {
	var result Bool
	intermidiary, err := r._unmarshalJSONbool(data)
	if err != nil {
		return result, fmt.Errorf("schema: Bool._unmarshalJSONBool: alias; %w", err)
	}
	result = Bool(intermidiary)
	return result, nil
}
func (r *Bool) _unmarshalJSONbool(data []byte) (bool, error) {
	var result bool
	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, fmt.Errorf("schema: Bool._unmarshalJSONbool: native boolean unwrap; %w", err)
	}
	return result, nil
}

func NumberFromJSON(x []byte) (*Number, error) {
	result := new(Number)
	err := result.UnmarshalJSON(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func NumberToJSON(x *Number) ([]byte, error) {
	return x.MarshalJSON()
}

var (
	_ json.Unmarshaler = (*Number)(nil)
	_ json.Marshaler   = (*Number)(nil)
)

func (r *Number) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONNumber(*r)
}
func (r *Number) _marshalJSONNumber(x Number) ([]byte, error) {
	return r._marshalJSONfloat64(float64(x))
}
func (r *Number) _marshalJSONfloat64(x float64) ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, fmt.Errorf("schema: Number._marshalJSONfloat64:; %w", err)
	}
	return result, nil
}
func (r *Number) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONNumber(data)
	if err != nil {
		return fmt.Errorf("schema: Number.UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *Number) _unmarshalJSONNumber(data []byte) (Number, error) {
	var result Number
	intermidiary, err := r._unmarshalJSONfloat64(data)
	if err != nil {
		return result, fmt.Errorf("schema: Number._unmarshalJSONNumber: alias; %w", err)
	}
	result = Number(intermidiary)
	return result, nil
}
func (r *Number) _unmarshalJSONfloat64(data []byte) (float64, error) {
	var result float64
	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, fmt.Errorf("schema: Number._unmarshalJSONfloat64: native number unwrap; %w", err)
	}
	return result, nil
}

func StringFromJSON(x []byte) (*String, error) {
	result := new(String)
	err := result.UnmarshalJSON(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func StringToJSON(x *String) ([]byte, error) {
	return x.MarshalJSON()
}

var (
	_ json.Unmarshaler = (*String)(nil)
	_ json.Marshaler   = (*String)(nil)
)

func (r *String) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONString(*r)
}
func (r *String) _marshalJSONString(x String) ([]byte, error) {
	return r._marshalJSONstring(string(x))
}
func (r *String) _marshalJSONstring(x string) ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, fmt.Errorf("schema: String._marshalJSONstring:; %w", err)
	}
	return result, nil
}
func (r *String) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONString(data)
	if err != nil {
		return fmt.Errorf("schema: String.UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *String) _unmarshalJSONString(data []byte) (String, error) {
	var result String
	intermidiary, err := r._unmarshalJSONstring(data)
	if err != nil {
		return result, fmt.Errorf("schema: String._unmarshalJSONString: alias; %w", err)
	}
	result = String(intermidiary)
	return result, nil
}
func (r *String) _unmarshalJSONstring(data []byte) (string, error) {
	var result string
	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, fmt.Errorf("schema: String._unmarshalJSONstring: native string unwrap; %w", err)
	}
	return result, nil
}

func BinaryFromJSON(x []byte) (*Binary, error) {
	result := new(Binary)
	err := result.UnmarshalJSON(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func BinaryToJSON(x *Binary) ([]byte, error) {
	return x.MarshalJSON()
}

var (
	_ json.Unmarshaler = (*Binary)(nil)
	_ json.Marshaler   = (*Binary)(nil)
)

func (r *Binary) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONBinary(*r)
}
func (r *Binary) _marshalJSONBinary(x Binary) ([]byte, error) {
	return r._marshalJSONSliceuint8([]uint8(x))
}
func (r *Binary) _marshalJSONSliceuint8(x []uint8) ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, fmt.Errorf("schema: Binary._marshalJSONSliceuint8:; %w", err)
	}
	return result, nil
}
func (r *Binary) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONBinary(data)
	if err != nil {
		return fmt.Errorf("schema: Binary.UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *Binary) _unmarshalJSONBinary(data []byte) (Binary, error) {
	var result Binary
	intermidiary, err := r._unmarshalJSONSliceuint8(data)
	if err != nil {
		return result, fmt.Errorf("schema: Binary._unmarshalJSONBinary: alias; %w", err)
	}
	result = Binary(intermidiary)
	return result, nil
}
func (r *Binary) _unmarshalJSONSliceuint8(data []byte) ([]uint8, error) {
	var result []uint8
	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, fmt.Errorf("schema: Binary._unmarshalJSONSliceuint8: native list unwrap; %w", err)
	}
	return result, nil
}

func ListFromJSON(x []byte) (*List, error) {
	result := new(List)
	err := result.UnmarshalJSON(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func ListToJSON(x *List) ([]byte, error) {
	return x.MarshalJSON()
}

var (
	_ json.Unmarshaler = (*List)(nil)
	_ json.Marshaler   = (*List)(nil)
)

func (r *List) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONList(*r)
}
func (r *List) _marshalJSONList(x List) ([]byte, error) {
	return r._marshalJSONSliceSchema([]Schema(x))
}
func (r *List) _marshalJSONSliceSchema(x []Schema) ([]byte, error) {
	partial := make([]json.RawMessage, len(x))
	for i, v := range x {
		item, err := r._marshalJSONSchema(v)
		if err != nil {
			return nil, fmt.Errorf("schema: List._marshalJSONSliceSchema: at index %d; %w", i, err)
		}
		partial[i] = item
	}
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("schema: List._marshalJSONSliceSchema:; %w", err)
	}
	return result, nil
}
func (r *List) _marshalJSONSchema(x Schema) ([]byte, error) {
	result, err := shared.JSONMarshal[Schema](x)
	if err != nil {
		return nil, fmt.Errorf("schema: List._marshalJSONSchema:; %w", err)
	}
	return result, nil
}
func (r *List) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONList(data)
	if err != nil {
		return fmt.Errorf("schema: List.UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *List) _unmarshalJSONList(data []byte) (List, error) {
	var result List
	intermidiary, err := r._unmarshalJSONSliceSchema(data)
	if err != nil {
		return result, fmt.Errorf("schema: List._unmarshalJSONList: alias; %w", err)
	}
	result = List(intermidiary)
	return result, nil
}
func (r *List) _unmarshalJSONSliceSchema(data []byte) ([]Schema, error) {
	result := make([]Schema, 0)
	var partial []json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return result, fmt.Errorf("schema: List._unmarshalJSONSliceSchema: native list unwrap; %w", err)
	}
	for i, v := range partial {
		item, err := r._unmarshalJSONSchema(v)
		if err != nil {
			return result, fmt.Errorf("schema: List._unmarshalJSONSliceSchema: at index %d; %w", i, err)
		}
		result = append(result, item)
	}
	return result, nil
}
func (r *List) _unmarshalJSONSchema(data []byte) (Schema, error) {
	result, err := shared.JSONUnmarshal[Schema](data)
	if err != nil {
		return result, fmt.Errorf("schema: List._unmarshalJSONSchema: native ref unwrap; %w", err)
	}
	return result, nil
}

func MapFromJSON(x []byte) (*Map, error) {
	result := new(Map)
	err := result.UnmarshalJSON(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func MapToJSON(x *Map) ([]byte, error) {
	return x.MarshalJSON()
}

var (
	_ json.Unmarshaler = (*Map)(nil)
	_ json.Marshaler   = (*Map)(nil)
)

func (r *Map) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	return r._marshalJSONMap(*r)
}
func (r *Map) _marshalJSONMap(x Map) ([]byte, error) {
	return r._marshalJSONmapLb_string_bLSchema(map[string]Schema(x))
}
func (r *Map) _marshalJSONmapLb_string_bLSchema(x map[string]Schema) ([]byte, error) {
	partial := make(map[string]json.RawMessage)
	for k, v := range x {
		key := string(k)
		value, err := r._marshalJSONSchema(v)
		if err != nil {
			return nil, fmt.Errorf("schema: Map._marshalJSONmapLb_string_bLSchema: value; %w", err)
		}
		partial[string(key)] = value
	}
	result, err := json.Marshal(partial)
	if err != nil {
		return nil, fmt.Errorf("schema: Map._marshalJSONmapLb_string_bLSchema:; %w", err)
	}
	return result, nil
}
func (r *Map) _marshalJSONstring(x string) ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, fmt.Errorf("schema: Map._marshalJSONstring:; %w", err)
	}
	return result, nil
}
func (r *Map) _marshalJSONSchema(x Schema) ([]byte, error) {
	result, err := shared.JSONMarshal[Schema](x)
	if err != nil {
		return nil, fmt.Errorf("schema: Map._marshalJSONSchema:; %w", err)
	}
	return result, nil
}
func (r *Map) UnmarshalJSON(data []byte) error {
	result, err := r._unmarshalJSONMap(data)
	if err != nil {
		return fmt.Errorf("schema: Map.UnmarshalJSON: %w", err)
	}
	*r = result
	return nil
}
func (r *Map) _unmarshalJSONMap(data []byte) (Map, error) {
	var result Map
	intermidiary, err := r._unmarshalJSONmapLb_string_bLSchema(data)
	if err != nil {
		return result, fmt.Errorf("schema: Map._unmarshalJSONMap: alias; %w", err)
	}
	result = Map(intermidiary)
	return result, nil
}
func (r *Map) _unmarshalJSONmapLb_string_bLSchema(data []byte) (map[string]Schema, error) {
	var partial map[string]json.RawMessage
	err := json.Unmarshal(data, &partial)
	if err != nil {
		return nil, fmt.Errorf("schema: Map._unmarshalJSONmapLb_string_bLSchema: native map unwrap; %w", err)
	}
	result := make(map[string]Schema)
	for k, v := range partial {
		key := string(k)
		value, err := r._unmarshalJSONSchema(v)
		if err != nil {
			return nil, fmt.Errorf("schema: Map._unmarshalJSONmapLb_string_bLSchema: value; %w", err)
		}
		result[key] = value
	}
	return result, nil
}
func (r *Map) _unmarshalJSONstring(data []byte) (string, error) {
	var result string
	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, fmt.Errorf("schema: Map._unmarshalJSONstring: native string unwrap; %w", err)
	}
	return result, nil
}
func (r *Map) _unmarshalJSONSchema(data []byte) (Schema, error) {
	result, err := shared.JSONUnmarshal[Schema](data)
	if err != nil {
		return result, fmt.Errorf("schema: Map._unmarshalJSONSchema: native ref unwrap; %w", err)
	}
	return result, nil
}
