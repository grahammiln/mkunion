// Code generated by mkunion. DO NOT EDIT.
package example

var _ TreeReducer[any] = (*TreeDefaultReduction[any])(nil)

type (
	TreeDefaultReduction[A any] struct {
		PanicOnFallback bool
		DefaultStopReduction bool
		OnBranch func(x *Branch, agg A) (result A, stop bool)
		OnLeaf func(x *Leaf, agg A) (result A, stop bool)
	}
)

func (t *TreeDefaultReduction[A]) ReduceBranch(x *Branch, agg A) (result A, stop bool) {
	if t.OnBranch != nil {
		return t.OnBranch(x, agg)
	}
	if t.PanicOnFallback {
		panic("no fallback allowed on undefined ReduceBranch")
	}
	return agg, t.DefaultStopReduction
}

func (t *TreeDefaultReduction[A]) ReduceLeaf(x *Leaf, agg A) (result A, stop bool) {
	if t.OnLeaf != nil {
		return t.OnLeaf(x, agg)
	}
	if t.PanicOnFallback {
		panic("no fallback allowed on undefined ReduceBranch")
	}
	return agg, t.DefaultStopReduction
}
