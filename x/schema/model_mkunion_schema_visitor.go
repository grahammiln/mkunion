// Code generated by mkunion. DO NOT EDIT.
package schema

import (
	"github.com/widmogrod/mkunion/f"
)

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
