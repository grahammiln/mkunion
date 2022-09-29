// Code generated by mkunion. DO NOT EDIT.
package example

type (
	CalcReducer[A any] interface {
		ReduceLit(x *Lit, agg A) (result A, stop bool)
		ReduceSum(x *Sum, agg A) (result A, stop bool)
		ReduceMul(x *Mul, agg A) (result A, stop bool)
	}
)

type CalcDepthFirstVisitor[A any] struct {
	stop   bool
	result A
	reduce CalcReducer[A]
}

var _ CalcVisitor = (*CalcDepthFirstVisitor[any])(nil)

func (d *CalcDepthFirstVisitor[A]) VisitLit(v *Lit) any {
	d.result, d.stop = d.reduce.ReduceLit(v, d.result)
	if d.stop {
		return nil
	}

	return nil
}

func (d *CalcDepthFirstVisitor[A]) VisitSum(v *Sum) any {
	d.result, d.stop = d.reduce.ReduceSum(v, d.result)
	if d.stop {
		return nil
	}
	if _ = v.Left.Accept(d); d.stop {
		return nil
	}
	if _ = v.Right.Accept(d); d.stop {
		return nil
	}

	return nil
}

func (d *CalcDepthFirstVisitor[A]) VisitMul(v *Mul) any {
	d.result, d.stop = d.reduce.ReduceMul(v, d.result)
	if d.stop {
		return nil
	}
	if _ = v.Left.Accept(d); d.stop {
		return nil
	}
	if _ = v.Right.Accept(d); d.stop {
		return nil
	}

	return nil
}

func ReduceCalc[A any](r CalcReducer[A], v Calc, init A) A {
	reducer := &CalcDepthFirstVisitor[A]{
		result: init,
		reduce: r,
	}

	_ = v.Accept(reducer)

	return reducer.result
}

var _ CalcReducer[any] = (*CalcDefaultReduction[any])(nil)

type (
	CalcDefaultReduction[A any] struct {
		PanicOnFallback      bool
		DefaultStopReduction bool
		OnLit                func(x *Lit, agg A) (result A, stop bool)
		OnSum                func(x *Sum, agg A) (result A, stop bool)
		OnMul                func(x *Mul, agg A) (result A, stop bool)
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
