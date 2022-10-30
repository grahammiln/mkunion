// Code generated by mkunion. DO NOT EDIT.
package example

var _ CalcReducer[any] = (*CalcDefaultReduction[any])(nil)

type (
	CalcDefaultReduction[A any] struct {
		PanicOnFallback bool
		DefaultStopReduction bool
		OnLit func(x *Lit, agg A) (result A, stop bool)
		OnSum func(x *Sum, agg A) (result A, stop bool)
		OnMul func(x *Mul, agg A) (result A, stop bool)
	}
)

func (t *CalcDefaultReduction[A]) ReduceLit(x *Lit, agg A) (result A, stop bool) {
	if t.OnLit != nil {
		return t.OnLit(x, agg)
	}
	if t.PanicOnFallback {
		panic("no fallback allowed on undefined ReduceBranch")
	}
	return agg, t.DefaultStopReduction
}

func (t *CalcDefaultReduction[A]) ReduceSum(x *Sum, agg A) (result A, stop bool) {
	if t.OnSum != nil {
		return t.OnSum(x, agg)
	}
	if t.PanicOnFallback {
		panic("no fallback allowed on undefined ReduceBranch")
	}
	return agg, t.DefaultStopReduction
}

func (t *CalcDefaultReduction[A]) ReduceMul(x *Mul, agg A) (result A, stop bool) {
	if t.OnMul != nil {
		return t.OnMul(x, agg)
	}
	if t.PanicOnFallback {
		panic("no fallback allowed on undefined ReduceBranch")
	}
	return agg, t.DefaultStopReduction
}
