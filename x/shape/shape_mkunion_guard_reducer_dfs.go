// Code generated by mkunion. DO NOT EDIT.
package shape

type (
	GuardReducer[A any] interface {
		ReduceRegexp(x *Regexp, agg A) (result A, stop bool)
		ReduceBetween(x *Between, agg A) (result A, stop bool)
		ReduceAndGuard(x *AndGuard, agg A) (result A, stop bool)
		ReduceOrGuard(x *OrGuard, agg A) (result A, stop bool)
	}
)

type GuardDepthFirstVisitor[A any] struct {
	stop   bool
	result A
	reduce GuardReducer[A]
}

var _ GuardVisitor = (*GuardDepthFirstVisitor[any])(nil)

func (d *GuardDepthFirstVisitor[A]) VisitRegexp(v *Regexp) any {
	d.result, d.stop = d.reduce.ReduceRegexp(v, d.result)
	if d.stop {
		return nil
	}

	return nil
}

func (d *GuardDepthFirstVisitor[A]) VisitBetween(v *Between) any {
	d.result, d.stop = d.reduce.ReduceBetween(v, d.result)
	if d.stop {
		return nil
	}

	return nil
}

func (d *GuardDepthFirstVisitor[A]) VisitAndGuard(v *AndGuard) any {
	d.result, d.stop = d.reduce.ReduceAndGuard(v, d.result)
	if d.stop {
		return nil
	}
	for idx := range v.L {
		if _ = v.L[idx].AcceptGuard(d); d.stop {
			return nil
		}
	}

	return nil
}

func (d *GuardDepthFirstVisitor[A]) VisitOrGuard(v *OrGuard) any {
	d.result, d.stop = d.reduce.ReduceOrGuard(v, d.result)
	if d.stop {
		return nil
	}
	for idx := range v.L {
		if _ = v.L[idx].AcceptGuard(d); d.stop {
			return nil
		}
	}

	return nil
}

func ReduceGuardDepthFirst[A any](r GuardReducer[A], v Guard, init A) A {
	reducer := &GuardDepthFirstVisitor[A]{
		result: init,
		reduce: r,
	}

	_ = v.AcceptGuard(reducer)

	return reducer.result
}
