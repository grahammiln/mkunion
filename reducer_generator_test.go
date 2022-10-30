package mkunion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTravers(t *testing.T) {
	g := ReducerGenerator{
		Name:        "Tree",
		PackageName: "visitor",
		Types:       []string{"Branch", "Leaf"},
		Branches: map[string][]Branching{
			"Branch": {
				{Lit: PtrStr("Lit")},
				{List: PtrStr("List")},
				{Map: PtrStr("Map")},
			},
		},
	}

	result, err := g.Generate()
	assert.NoError(t, err)
	assert.Equal(t, `// Code generated by mkunion. DO NOT EDIT.
package visitor

type (
	TreeReducer[A any] interface {
		ReduceBranch(x *Branch, agg A) (result A, stop bool)
		ReduceLeaf(x *Leaf, agg A) (result A, stop bool)
	}
)

type TreeDepthFirstVisitor[A any] struct {
	stop   bool
	result A
	reduce TreeReducer[A]
}

var _ TreeVisitor = (*TreeDepthFirstVisitor[any])(nil)

func (d *TreeDepthFirstVisitor[A]) VisitBranch(v *Branch) any {
	d.result, d.stop = d.reduce.ReduceBranch(v, d.result)
	if d.stop {
		return nil
	}
	if _ = v.Lit.Accept(d); d.stop {
		return nil
	}
	for idx := range v.List {
		if _ = v.List[idx].Accept(d); d.stop {
			return nil
		}
	}
	for idx, _ := range v.Map {
		if _ = v.Map[idx].Accept(d); d.stop {
			return nil
		}
	}

	return nil
}

func (d *TreeDepthFirstVisitor[A]) VisitLeaf(v *Leaf) any {
	d.result, d.stop = d.reduce.ReduceLeaf(v, d.result)
	if d.stop {
		return nil
	}

	return nil
}

func ReduceTreeDepthFirst[A any](r TreeReducer[A], v Tree, init A) A {
	reducer := &TreeDepthFirstVisitor[A]{
		result: init,
		reduce: r,
	}

	_ = v.Accept(reducer)

	return reducer.result
}

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
`, string(result))
}
