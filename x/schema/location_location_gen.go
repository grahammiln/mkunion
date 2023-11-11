// Code generated by mkunion. DO NOT EDIT.
package schema

import "github.com/widmogrod/mkunion/f"

//mkunion-extension:visitor

type LocationVisitor interface {
	VisitLocationField(v *LocationField) any
	VisitLocationIndex(v *LocationIndex) any
	VisitLocationAnything(v *LocationAnything) any
}

type Location interface {
	AcceptLocation(g LocationVisitor) any
}

func (r *LocationField) AcceptLocation(v LocationVisitor) any    { return v.VisitLocationField(r) }
func (r *LocationIndex) AcceptLocation(v LocationVisitor) any    { return v.VisitLocationIndex(r) }
func (r *LocationAnything) AcceptLocation(v LocationVisitor) any { return v.VisitLocationAnything(r) }

var (
	_ Location = (*LocationField)(nil)
	_ Location = (*LocationIndex)(nil)
	_ Location = (*LocationAnything)(nil)
)

func MatchLocation[TOut any](
	x Location,
	f1 func(x *LocationField) TOut,
	f2 func(x *LocationIndex) TOut,
	f3 func(x *LocationAnything) TOut,
	df func(x Location) TOut,
) TOut {
	return f.Match3(x, f1, f2, f3, df)
}

func MatchLocationR2[TOut1, TOut2 any](
	x Location,
	f1 func(x *LocationField) (TOut1, TOut2),
	f2 func(x *LocationIndex) (TOut1, TOut2),
	f3 func(x *LocationAnything) (TOut1, TOut2),
	df func(x Location) (TOut1, TOut2),
) (TOut1, TOut2) {
	return f.Match3R2(x, f1, f2, f3, df)
}

func MustMatchLocation[TOut any](
	x Location,
	f1 func(x *LocationField) TOut,
	f2 func(x *LocationIndex) TOut,
	f3 func(x *LocationAnything) TOut,
) TOut {
	return f.MustMatch3(x, f1, f2, f3)
}

func MustMatchLocationR0(
	x Location,
	f1 func(x *LocationField),
	f2 func(x *LocationIndex),
	f3 func(x *LocationAnything),
) {
	f.MustMatch3R0(x, f1, f2, f3)
}

func MustMatchLocationR2[TOut1, TOut2 any](
	x Location,
	f1 func(x *LocationField) (TOut1, TOut2),
	f2 func(x *LocationIndex) (TOut1, TOut2),
	f3 func(x *LocationAnything) (TOut1, TOut2),
) (TOut1, TOut2) {
	return f.MustMatch3R2(x, f1, f2, f3)
}

// mkunion-extension:reducer_dfs
type (
	LocationReducer[A any] interface {
		ReduceLocationField(x *LocationField, agg A) (result A, stop bool)
		ReduceLocationIndex(x *LocationIndex, agg A) (result A, stop bool)
		ReduceLocationAnything(x *LocationAnything, agg A) (result A, stop bool)
	}
)

type LocationDepthFirstVisitor[A any] struct {
	stop   bool
	result A
	reduce LocationReducer[A]
}

var _ LocationVisitor = (*LocationDepthFirstVisitor[any])(nil)

func (d *LocationDepthFirstVisitor[A]) VisitLocationField(v *LocationField) any {
	d.result, d.stop = d.reduce.ReduceLocationField(v, d.result)
	if d.stop {
		return nil
	}

	return nil
}

func (d *LocationDepthFirstVisitor[A]) VisitLocationIndex(v *LocationIndex) any {
	d.result, d.stop = d.reduce.ReduceLocationIndex(v, d.result)
	if d.stop {
		return nil
	}

	return nil
}

func (d *LocationDepthFirstVisitor[A]) VisitLocationAnything(v *LocationAnything) any {
	d.result, d.stop = d.reduce.ReduceLocationAnything(v, d.result)
	if d.stop {
		return nil
	}

	return nil
}

func ReduceLocationDepthFirst[A any](r LocationReducer[A], v Location, init A) A {
	reducer := &LocationDepthFirstVisitor[A]{
		result: init,
		reduce: r,
	}

	_ = v.AcceptLocation(reducer)

	return reducer.result
}

// mkunion-extension:reducer_bfs
var _ LocationVisitor = (*LocationBreadthFirstVisitor[any])(nil)

type LocationBreadthFirstVisitor[A any] struct {
	stop   bool
	result A
	reduce LocationReducer[A]

	queue         []Location
	visited       map[Location]bool
	shouldExecute map[Location]bool
}

func (d *LocationBreadthFirstVisitor[A]) VisitLocationField(v *LocationField) any {
	d.queue = append(d.queue, v)

	if d.shouldExecute[v] {
		d.shouldExecute[v] = false
		d.result, d.stop = d.reduce.ReduceLocationField(v, d.result)
	} else {
		d.execute()
	}
	return nil
}

func (d *LocationBreadthFirstVisitor[A]) VisitLocationIndex(v *LocationIndex) any {
	d.queue = append(d.queue, v)

	if d.shouldExecute[v] {
		d.shouldExecute[v] = false
		d.result, d.stop = d.reduce.ReduceLocationIndex(v, d.result)
	} else {
		d.execute()
	}
	return nil
}

func (d *LocationBreadthFirstVisitor[A]) VisitLocationAnything(v *LocationAnything) any {
	d.queue = append(d.queue, v)

	if d.shouldExecute[v] {
		d.shouldExecute[v] = false
		d.result, d.stop = d.reduce.ReduceLocationAnything(v, d.result)
	} else {
		d.execute()
	}
	return nil
}

func (d *LocationBreadthFirstVisitor[A]) execute() {
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
		i.AcceptLocation(d)
	}

	return
}

func (d *LocationBreadthFirstVisitor[A]) pop() Location {
	i := d.queue[0]
	d.queue = d.queue[1:]
	return i
}

func ReduceLocationBreadthFirst[A any](r LocationReducer[A], v Location, init A) A {
	reducer := &LocationBreadthFirstVisitor[A]{
		result:        init,
		reduce:        r,
		queue:         []Location{v},
		visited:       make(map[Location]bool),
		shouldExecute: make(map[Location]bool),
	}

	_ = v.AcceptLocation(reducer)

	return reducer.result
}

// mkunion-extension:default_reducer
var _ LocationReducer[any] = (*LocationDefaultReduction[any])(nil)

type (
	LocationDefaultReduction[A any] struct {
		PanicOnFallback      bool
		DefaultStopReduction bool
		OnLocationField      func(x *LocationField, agg A) (result A, stop bool)
		OnLocationIndex      func(x *LocationIndex, agg A) (result A, stop bool)
		OnLocationAnything   func(x *LocationAnything, agg A) (result A, stop bool)
	}
)

func (t *LocationDefaultReduction[A]) ReduceLocationField(x *LocationField, agg A) (result A, stop bool) {
	if t.OnLocationField != nil {
		return t.OnLocationField(x, agg)
	}
	if t.PanicOnFallback {
		panic("no fallback allowed on undefined ReduceBranch")
	}
	return agg, t.DefaultStopReduction
}

func (t *LocationDefaultReduction[A]) ReduceLocationIndex(x *LocationIndex, agg A) (result A, stop bool) {
	if t.OnLocationIndex != nil {
		return t.OnLocationIndex(x, agg)
	}
	if t.PanicOnFallback {
		panic("no fallback allowed on undefined ReduceBranch")
	}
	return agg, t.DefaultStopReduction
}

func (t *LocationDefaultReduction[A]) ReduceLocationAnything(x *LocationAnything, agg A) (result A, stop bool) {
	if t.OnLocationAnything != nil {
		return t.OnLocationAnything(x, agg)
	}
	if t.PanicOnFallback {
		panic("no fallback allowed on undefined ReduceBranch")
	}
	return agg, t.DefaultStopReduction
}

// mkunion-extension:default_visitor
type LocationDefaultVisitor[A any] struct {
	Default            A
	OnLocationField    func(x *LocationField) A
	OnLocationIndex    func(x *LocationIndex) A
	OnLocationAnything func(x *LocationAnything) A
}

func (t *LocationDefaultVisitor[A]) VisitLocationField(v *LocationField) any {
	if t.OnLocationField != nil {
		return t.OnLocationField(v)
	}
	return t.Default
}
func (t *LocationDefaultVisitor[A]) VisitLocationIndex(v *LocationIndex) any {
	if t.OnLocationIndex != nil {
		return t.OnLocationIndex(v)
	}
	return t.Default
}
func (t *LocationDefaultVisitor[A]) VisitLocationAnything(v *LocationAnything) any {
	if t.OnLocationAnything != nil {
		return t.OnLocationAnything(v)
	}
	return t.Default
}
