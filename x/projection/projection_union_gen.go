// Code generated by mkunion. DO NOT EDIT.
package projection

type DataVisitor[A any] interface {
	VisitRecord(v *Record[A]) any
	VisitWatermark(v *Watermark[A]) any
}

type Data[A any] interface {
	AcceptData(g DataVisitor[A]) any
}

var (
	_ Data[any] = (*Record[any])(nil)
	_ Data[any] = (*Watermark[any])(nil)
)

func (r *Record[A]) AcceptData(v DataVisitor[A]) any    { return v.VisitRecord(r) }
func (r *Watermark[A]) AcceptData(v DataVisitor[A]) any { return v.VisitWatermark(r) }

func MatchDataR3[A any, T0, T1, T2 any](
	x Data[A],
	f1 func(x *Record[A]) (T0, T1, T2),
	f2 func(x *Watermark[A]) (T0, T1, T2),
) (T0, T1, T2) {
	switch v := x.(type) {
	case *Record[A]:
		return f1(v)
	case *Watermark[A]:
		return f2(v)
	}
	var result1 T0
	var result2 T1
	var result3 T2
	return result1, result2, result3
}

func MatchDataR2[A any, T0, T1 any](
	x Data[A],
	f1 func(x *Record[A]) (T0, T1),
	f2 func(x *Watermark[A]) (T0, T1),
) (T0, T1) {
	switch v := x.(type) {
	case *Record[A]:
		return f1(v)
	case *Watermark[A]:
		return f2(v)
	}
	var result1 T0
	var result2 T1
	return result1, result2
}

func MatchDataR1[A any, T0 any](
	x Data[A],
	f1 func(x *Record[A]) T0,
	f2 func(x *Watermark[A]) T0,
) T0 {
	switch v := x.(type) {
	case *Record[A]:
		return f1(v)
	case *Watermark[A]:
		return f2(v)
	}
	var result1 T0
	return result1
}

func MatchDataR0[A any](
	x Data[A],
	f1 func(x *Record[A]),
	f2 func(x *Watermark[A]),
) {
	switch v := x.(type) {
	case *Record[A]:
		f1(v)
	case *Watermark[A]:
		f2(v)
	}
}

type EitherVisitor[A any, B any] interface {
	VisitLeft(v *Left[A, B]) any
	VisitRight(v *Right[A, B]) any
}

type Either[A any, B any] interface {
	AcceptEither(g EitherVisitor[A, B]) any
}

var (
	_ Either[any, any] = (*Left[any, any])(nil)
	_ Either[any, any] = (*Right[any, any])(nil)
)

func (r *Left[A, B]) AcceptEither(v EitherVisitor[A, B]) any  { return v.VisitLeft(r) }
func (r *Right[A, B]) AcceptEither(v EitherVisitor[A, B]) any { return v.VisitRight(r) }

func MatchEitherR3[A any, B any, T0, T1, T2 any](
	x Either[A, B],
	f1 func(x *Left[A, B]) (T0, T1, T2),
	f2 func(x *Right[A, B]) (T0, T1, T2),
) (T0, T1, T2) {
	switch v := x.(type) {
	case *Left[A, B]:
		return f1(v)
	case *Right[A, B]:
		return f2(v)
	}
	var result1 T0
	var result2 T1
	var result3 T2
	return result1, result2, result3
}

func MatchEitherR2[A any, B any, T0, T1 any](
	x Either[A, B],
	f1 func(x *Left[A, B]) (T0, T1),
	f2 func(x *Right[A, B]) (T0, T1),
) (T0, T1) {
	switch v := x.(type) {
	case *Left[A, B]:
		return f1(v)
	case *Right[A, B]:
		return f2(v)
	}
	var result1 T0
	var result2 T1
	return result1, result2
}

func MatchEitherR1[A any, B any, T0 any](
	x Either[A, B],
	f1 func(x *Left[A, B]) T0,
	f2 func(x *Right[A, B]) T0,
) T0 {
	switch v := x.(type) {
	case *Left[A, B]:
		return f1(v)
	case *Right[A, B]:
		return f2(v)
	}
	var result1 T0
	return result1
}

func MatchEitherR0[A any, B any](
	x Either[A, B],
	f1 func(x *Left[A, B]),
	f2 func(x *Right[A, B]),
) {
	switch v := x.(type) {
	case *Left[A, B]:
		f1(v)
	case *Right[A, B]:
		f2(v)
	}
}
