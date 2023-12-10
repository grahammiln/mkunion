// Code generated by mkunion. DO NOT EDIT.
package schema

import "github.com/widmogrod/mkunion/f"
import "github.com/widmogrod/mkunion/x/shared"
import "encoding/json"
import "fmt"

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

// mkunion-extension:reducer_dfs
type (
	SchemaReducer[A any] interface {
		ReduceNone(x *None, agg A) (result A, stop bool)
		ReduceBool(x *Bool, agg A) (result A, stop bool)
		ReduceNumber(x *Number, agg A) (result A, stop bool)
		ReduceString(x *String, agg A) (result A, stop bool)
		ReduceBinary(x *Binary, agg A) (result A, stop bool)
		ReduceList(x *List, agg A) (result A, stop bool)
		ReduceMap(x *Map, agg A) (result A, stop bool)
	}
)

type SchemaDepthFirstVisitor[A any] struct {
	stop   bool
	result A
	reduce SchemaReducer[A]
}

var _ SchemaVisitor = (*SchemaDepthFirstVisitor[any])(nil)

func (d *SchemaDepthFirstVisitor[A]) VisitNone(v *None) any {
	d.result, d.stop = d.reduce.ReduceNone(v, d.result)
	if d.stop {
		return nil
	}

	return nil
}

func (d *SchemaDepthFirstVisitor[A]) VisitBool(v *Bool) any {
	d.result, d.stop = d.reduce.ReduceBool(v, d.result)
	if d.stop {
		return nil
	}

	return nil
}

func (d *SchemaDepthFirstVisitor[A]) VisitNumber(v *Number) any {
	d.result, d.stop = d.reduce.ReduceNumber(v, d.result)
	if d.stop {
		return nil
	}

	return nil
}

func (d *SchemaDepthFirstVisitor[A]) VisitString(v *String) any {
	d.result, d.stop = d.reduce.ReduceString(v, d.result)
	if d.stop {
		return nil
	}

	return nil
}

func (d *SchemaDepthFirstVisitor[A]) VisitBinary(v *Binary) any {
	d.result, d.stop = d.reduce.ReduceBinary(v, d.result)
	if d.stop {
		return nil
	}

	return nil
}

func (d *SchemaDepthFirstVisitor[A]) VisitList(v *List) any {
	d.result, d.stop = d.reduce.ReduceList(v, d.result)
	if d.stop {
		return nil
	}
	for idx := range v.Items {
		if _ = v.Items[idx].AcceptSchema(d); d.stop {
			return nil
		}
	}

	return nil
}

func (d *SchemaDepthFirstVisitor[A]) VisitMap(v *Map) any {
	d.result, d.stop = d.reduce.ReduceMap(v, d.result)
	if d.stop {
		return nil
	}

	return nil
}

func ReduceSchemaDepthFirst[A any](r SchemaReducer[A], v Schema, init A) A {
	reducer := &SchemaDepthFirstVisitor[A]{
		result: init,
		reduce: r,
	}

	_ = v.AcceptSchema(reducer)

	return reducer.result
}

// mkunion-extension:reducer_bfs
var _ SchemaVisitor = (*SchemaBreadthFirstVisitor[any])(nil)

type SchemaBreadthFirstVisitor[A any] struct {
	stop   bool
	result A
	reduce SchemaReducer[A]

	queue         []Schema
	visited       map[Schema]bool
	shouldExecute map[Schema]bool
}

func (d *SchemaBreadthFirstVisitor[A]) VisitNone(v *None) any {
	d.queue = append(d.queue, v)

	if d.shouldExecute[v] {
		d.shouldExecute[v] = false
		d.result, d.stop = d.reduce.ReduceNone(v, d.result)
	} else {
		d.execute()
	}
	return nil
}

func (d *SchemaBreadthFirstVisitor[A]) VisitBool(v *Bool) any {
	d.queue = append(d.queue, v)

	if d.shouldExecute[v] {
		d.shouldExecute[v] = false
		d.result, d.stop = d.reduce.ReduceBool(v, d.result)
	} else {
		d.execute()
	}
	return nil
}

func (d *SchemaBreadthFirstVisitor[A]) VisitNumber(v *Number) any {
	d.queue = append(d.queue, v)

	if d.shouldExecute[v] {
		d.shouldExecute[v] = false
		d.result, d.stop = d.reduce.ReduceNumber(v, d.result)
	} else {
		d.execute()
	}
	return nil
}

func (d *SchemaBreadthFirstVisitor[A]) VisitString(v *String) any {
	d.queue = append(d.queue, v)

	if d.shouldExecute[v] {
		d.shouldExecute[v] = false
		d.result, d.stop = d.reduce.ReduceString(v, d.result)
	} else {
		d.execute()
	}
	return nil
}

func (d *SchemaBreadthFirstVisitor[A]) VisitBinary(v *Binary) any {
	d.queue = append(d.queue, v)

	if d.shouldExecute[v] {
		d.shouldExecute[v] = false
		d.result, d.stop = d.reduce.ReduceBinary(v, d.result)
	} else {
		d.execute()
	}
	return nil
}

func (d *SchemaBreadthFirstVisitor[A]) VisitList(v *List) any {
	d.queue = append(d.queue, v)
	for idx := range v.Items {
		d.queue = append(d.queue, v.Items[idx])
	}

	if d.shouldExecute[v] {
		d.shouldExecute[v] = false
		d.result, d.stop = d.reduce.ReduceList(v, d.result)
	} else {
		d.execute()
	}
	return nil
}

func (d *SchemaBreadthFirstVisitor[A]) VisitMap(v *Map) any {
	d.queue = append(d.queue, v)

	if d.shouldExecute[v] {
		d.shouldExecute[v] = false
		d.result, d.stop = d.reduce.ReduceMap(v, d.result)
	} else {
		d.execute()
	}
	return nil
}

func (d *SchemaBreadthFirstVisitor[A]) execute() {
	for len(d.queue) > 0 {
		if d.stop {
			return
		}

		i := d.pop()
		if d.visited[i] {
			continue
		}
		d.visited[i] = true
		d.shouldExecute[i] = true
		i.AcceptSchema(d)
	}

	return
}

func (d *SchemaBreadthFirstVisitor[A]) pop() Schema {
	i := d.queue[0]
	d.queue = d.queue[1:]
	return i
}

func ReduceSchemaBreadthFirst[A any](r SchemaReducer[A], v Schema, init A) A {
	reducer := &SchemaBreadthFirstVisitor[A]{
		result:        init,
		reduce:        r,
		queue:         []Schema{v},
		visited:       make(map[Schema]bool),
		shouldExecute: make(map[Schema]bool),
	}

	_ = v.AcceptSchema(reducer)

	return reducer.result
}

// mkunion-extension:default_reducer
var _ SchemaReducer[any] = (*SchemaDefaultReduction[any])(nil)

type (
	SchemaDefaultReduction[A any] struct {
		PanicOnFallback      bool
		DefaultStopReduction bool
		OnNone               func(x *None, agg A) (result A, stop bool)
		OnBool               func(x *Bool, agg A) (result A, stop bool)
		OnNumber             func(x *Number, agg A) (result A, stop bool)
		OnString             func(x *String, agg A) (result A, stop bool)
		OnBinary             func(x *Binary, agg A) (result A, stop bool)
		OnList               func(x *List, agg A) (result A, stop bool)
		OnMap                func(x *Map, agg A) (result A, stop bool)
	}
)

func (t *SchemaDefaultReduction[A]) ReduceNone(x *None, agg A) (result A, stop bool) {
	if t.OnNone != nil {
		return t.OnNone(x, agg)
	}
	if t.PanicOnFallback {
		panic("no fallback allowed on undefined ReduceNone")
	}
	return agg, t.DefaultStopReduction
}

func (t *SchemaDefaultReduction[A]) ReduceBool(x *Bool, agg A) (result A, stop bool) {
	if t.OnBool != nil {
		return t.OnBool(x, agg)
	}
	if t.PanicOnFallback {
		panic("no fallback allowed on undefined ReduceBool")
	}
	return agg, t.DefaultStopReduction
}

func (t *SchemaDefaultReduction[A]) ReduceNumber(x *Number, agg A) (result A, stop bool) {
	if t.OnNumber != nil {
		return t.OnNumber(x, agg)
	}
	if t.PanicOnFallback {
		panic("no fallback allowed on undefined ReduceNumber")
	}
	return agg, t.DefaultStopReduction
}

func (t *SchemaDefaultReduction[A]) ReduceString(x *String, agg A) (result A, stop bool) {
	if t.OnString != nil {
		return t.OnString(x, agg)
	}
	if t.PanicOnFallback {
		panic("no fallback allowed on undefined ReduceString")
	}
	return agg, t.DefaultStopReduction
}

func (t *SchemaDefaultReduction[A]) ReduceBinary(x *Binary, agg A) (result A, stop bool) {
	if t.OnBinary != nil {
		return t.OnBinary(x, agg)
	}
	if t.PanicOnFallback {
		panic("no fallback allowed on undefined ReduceBinary")
	}
	return agg, t.DefaultStopReduction
}

func (t *SchemaDefaultReduction[A]) ReduceList(x *List, agg A) (result A, stop bool) {
	if t.OnList != nil {
		return t.OnList(x, agg)
	}
	if t.PanicOnFallback {
		panic("no fallback allowed on undefined ReduceList")
	}
	return agg, t.DefaultStopReduction
}

func (t *SchemaDefaultReduction[A]) ReduceMap(x *Map, agg A) (result A, stop bool) {
	if t.OnMap != nil {
		return t.OnMap(x, agg)
	}
	if t.PanicOnFallback {
		panic("no fallback allowed on undefined ReduceMap")
	}
	return agg, t.DefaultStopReduction
}

// mkunion-extension:default_visitor
type SchemaDefaultVisitor[A any] struct {
	Default  A
	OnNone   func(x *None) A
	OnBool   func(x *Bool) A
	OnNumber func(x *Number) A
	OnString func(x *String) A
	OnBinary func(x *Binary) A
	OnList   func(x *List) A
	OnMap    func(x *Map) A
}

func (t *SchemaDefaultVisitor[A]) VisitNone(v *None) any {
	if t.OnNone != nil {
		return t.OnNone(v)
	}
	return t.Default
}
func (t *SchemaDefaultVisitor[A]) VisitBool(v *Bool) any {
	if t.OnBool != nil {
		return t.OnBool(v)
	}
	return t.Default
}
func (t *SchemaDefaultVisitor[A]) VisitNumber(v *Number) any {
	if t.OnNumber != nil {
		return t.OnNumber(v)
	}
	return t.Default
}
func (t *SchemaDefaultVisitor[A]) VisitString(v *String) any {
	if t.OnString != nil {
		return t.OnString(v)
	}
	return t.Default
}
func (t *SchemaDefaultVisitor[A]) VisitBinary(v *Binary) any {
	if t.OnBinary != nil {
		return t.OnBinary(v)
	}
	return t.Default
}
func (t *SchemaDefaultVisitor[A]) VisitList(v *List) any {
	if t.OnList != nil {
		return t.OnList(v)
	}
	return t.Default
}
func (t *SchemaDefaultVisitor[A]) VisitMap(v *Map) any {
	if t.OnMap != nil {
		return t.OnMap(v)
	}
	return t.Default
}

// mkunion-extension:json
type SchemaUnionJSON struct {
	Type   string          `json:"$type,omitempty"`
	None   json.RawMessage `json:"github.com/widmogrod/mkunion/x/schema.None,omitempty"`
	Bool   json.RawMessage `json:"github.com/widmogrod/mkunion/x/schema.Bool,omitempty"`
	Number json.RawMessage `json:"github.com/widmogrod/mkunion/x/schema.Number,omitempty"`
	String json.RawMessage `json:"github.com/widmogrod/mkunion/x/schema.String,omitempty"`
	Binary json.RawMessage `json:"github.com/widmogrod/mkunion/x/schema.Binary,omitempty"`
	List   json.RawMessage `json:"github.com/widmogrod/mkunion/x/schema.List,omitempty"`
	Map    json.RawMessage `json:"github.com/widmogrod/mkunion/x/schema.Map,omitempty"`
}

func SchemaFromJSON(x []byte) (Schema, error) {
	var data SchemaUnionJSON
	err := json.Unmarshal(x, &data)
	if err != nil {
		return nil, err
	}

	switch data.Type {
	case "github.com/widmogrod/mkunion/x/schema.None":
		return NoneFromJSON(data.None)
	case "github.com/widmogrod/mkunion/x/schema.Bool":
		return BoolFromJSON(data.Bool)
	case "github.com/widmogrod/mkunion/x/schema.Number":
		return NumberFromJSON(data.Number)
	case "github.com/widmogrod/mkunion/x/schema.String":
		return StringFromJSON(data.String)
	case "github.com/widmogrod/mkunion/x/schema.Binary":
		return BinaryFromJSON(data.Binary)
	case "github.com/widmogrod/mkunion/x/schema.List":
		return ListFromJSON(data.List)
	case "github.com/widmogrod/mkunion/x/schema.Map":
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

	return nil, fmt.Errorf("unknown type %s", data.Type)
}

func SchemaToJSON(x Schema) ([]byte, error) {
	return MustMatchSchemaR2(
		x,
		func(x *None) ([]byte, error) {
			body, err := NoneToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(SchemaUnionJSON{
				Type: "github.com/widmogrod/mkunion/x/schema.None",
				None: body,
			})
		},
		func(x *Bool) ([]byte, error) {
			body, err := BoolToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(SchemaUnionJSON{
				Type: "github.com/widmogrod/mkunion/x/schema.Bool",
				Bool: body,
			})
		},
		func(x *Number) ([]byte, error) {
			body, err := NumberToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(SchemaUnionJSON{
				Type:   "github.com/widmogrod/mkunion/x/schema.Number",
				Number: body,
			})
		},
		func(x *String) ([]byte, error) {
			body, err := StringToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(SchemaUnionJSON{
				Type:   "github.com/widmogrod/mkunion/x/schema.String",
				String: body,
			})
		},
		func(x *Binary) ([]byte, error) {
			body, err := BinaryToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(SchemaUnionJSON{
				Type:   "github.com/widmogrod/mkunion/x/schema.Binary",
				Binary: body,
			})
		},
		func(x *List) ([]byte, error) {
			body, err := ListToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(SchemaUnionJSON{
				Type: "github.com/widmogrod/mkunion/x/schema.List",
				List: body,
			})
		},
		func(x *Map) ([]byte, error) {
			body, err := MapToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(SchemaUnionJSON{
				Type: "github.com/widmogrod/mkunion/x/schema.Map",
				Map:  body,
			})
		},
	)
}

func NoneFromJSON(x []byte) (*None, error) {
	var result *None = &None{}

	// if is Struct
	err := shared.JsonParseObject(x, func(key string, value []byte) error {
		switch key {
		}

		return fmt.Errorf("schema.NoneFromJSON: unknown key %s", key)
	})

	return result, err
}

func NoneToJSON(x *None) ([]byte, error) {
	return json.Marshal(map[string]json.RawMessage{})
}

func (self *None) MarshalJSON() ([]byte, error) {
	return NoneToJSON(self)
}

func (self *None) UnmarshalJSON(x []byte) error {
	n, err := NoneFromJSON(x)
	if err != nil {
		return err
	}
	*self = *n
	return nil
}

func BoolFromJSON(x []byte) (*Bool, error) {
	var result *Bool = &Bool{}

	// if is Struct
	err := shared.JsonParseObject(x, func(key string, value []byte) error {
		switch key {
		case "B":
			return json.Unmarshal(value, &result.B)
		}

		return fmt.Errorf("schema.BoolFromJSON: unknown key %s", key)
	})

	return result, err
}

func BoolToJSON(x *Bool) ([]byte, error) {
	field_B, err := json.Marshal(x.B)
	if err != nil {
		return nil, err
	}
	return json.Marshal(map[string]json.RawMessage{
		"B": field_B,
	})
}

func (self *Bool) MarshalJSON() ([]byte, error) {
	return BoolToJSON(self)
}

func (self *Bool) UnmarshalJSON(x []byte) error {
	n, err := BoolFromJSON(x)
	if err != nil {
		return err
	}
	*self = *n
	return nil
}

func NumberFromJSON(x []byte) (*Number, error) {
	var result *Number = &Number{}

	// if is Struct
	err := shared.JsonParseObject(x, func(key string, value []byte) error {
		switch key {
		case "N":
			return json.Unmarshal(value, &result.N)
		}

		return fmt.Errorf("schema.NumberFromJSON: unknown key %s", key)
	})

	return result, err
}

func NumberToJSON(x *Number) ([]byte, error) {
	field_N, err := json.Marshal(x.N)
	if err != nil {
		return nil, err
	}
	return json.Marshal(map[string]json.RawMessage{
		"N": field_N,
	})
}

func (self *Number) MarshalJSON() ([]byte, error) {
	return NumberToJSON(self)
}

func (self *Number) UnmarshalJSON(x []byte) error {
	n, err := NumberFromJSON(x)
	if err != nil {
		return err
	}
	*self = *n
	return nil
}

func StringFromJSON(x []byte) (*String, error) {
	var result *String = &String{}

	// if is Struct
	err := shared.JsonParseObject(x, func(key string, value []byte) error {
		switch key {
		case "S":
			return json.Unmarshal(value, &result.S)
		}

		return fmt.Errorf("schema.StringFromJSON: unknown key %s", key)
	})

	return result, err
}

func StringToJSON(x *String) ([]byte, error) {
	field_S, err := json.Marshal(x.S)
	if err != nil {
		return nil, err
	}
	return json.Marshal(map[string]json.RawMessage{
		"S": field_S,
	})
}

func (self *String) MarshalJSON() ([]byte, error) {
	return StringToJSON(self)
}

func (self *String) UnmarshalJSON(x []byte) error {
	n, err := StringFromJSON(x)
	if err != nil {
		return err
	}
	*self = *n
	return nil
}

func BinaryFromJSON(x []byte) (*Binary, error) {
	var result *Binary = &Binary{}

	// if is Struct
	err := shared.JsonParseObject(x, func(key string, value []byte) error {
		switch key {
		case "B":
			return json.Unmarshal(value, &result.B)
		}

		return fmt.Errorf("schema.BinaryFromJSON: unknown key %s", key)
	})

	return result, err
}

func BinaryToJSON(x *Binary) ([]byte, error) {
	field_B, err := json.Marshal(x.B)
	if err != nil {
		return nil, err
	}
	return json.Marshal(map[string]json.RawMessage{
		"B": field_B,
	})
}

func (self *Binary) MarshalJSON() ([]byte, error) {
	return BinaryToJSON(self)
}

func (self *Binary) UnmarshalJSON(x []byte) error {
	n, err := BinaryFromJSON(x)
	if err != nil {
		return err
	}
	*self = *n
	return nil
}

func ListFromJSON(x []byte) (*List, error) {
	var result *List = &List{}

	// if is Struct
	err := shared.JsonParseObject(x, func(key string, value []byte) error {
		switch key {
		case "Items":
			return json.Unmarshal(value, &result.Items)
		}

		return fmt.Errorf("schema.ListFromJSON: unknown key %s", key)
	})

	return result, err
}

func ListToJSON(x *List) ([]byte, error) {
	field_Items, err := json.Marshal(x.Items)
	if err != nil {
		return nil, err
	}
	return json.Marshal(map[string]json.RawMessage{
		"Items": field_Items,
	})
}

func (self *List) MarshalJSON() ([]byte, error) {
	return ListToJSON(self)
}

func (self *List) UnmarshalJSON(x []byte) error {
	n, err := ListFromJSON(x)
	if err != nil {
		return err
	}
	*self = *n
	return nil
}

func MapFromJSON(x []byte) (*Map, error) {
	var result *Map = &Map{}

	// if is Struct
	err := shared.JsonParseObject(x, func(key string, value []byte) error {
		switch key {
		case "Field":
			return json.Unmarshal(value, &result.Field)
		}

		return fmt.Errorf("schema.MapFromJSON: unknown key %s", key)
	})

	return result, err
}

func MapToJSON(x *Map) ([]byte, error) {
	field_Field, err := json.Marshal(x.Field)
	if err != nil {
		return nil, err
	}
	return json.Marshal(map[string]json.RawMessage{
		"Field": field_Field,
	})
}

func (self *Map) MarshalJSON() ([]byte, error) {
	return MapToJSON(self)
}

func (self *Map) UnmarshalJSON(x []byte) error {
	n, err := MapFromJSON(x)
	if err != nil {
		return err
	}
	*self = *n
	return nil
}
