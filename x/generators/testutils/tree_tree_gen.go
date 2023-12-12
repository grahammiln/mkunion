// Code generated by mkunion. DO NOT EDIT.
package testutils

import "github.com/widmogrod/mkunion/f"
import "github.com/widmogrod/mkunion/x/schema"
import "github.com/widmogrod/mkunion/x/shape"
import "github.com/widmogrod/mkunion/x/shared"
import "encoding/json"
import "fmt"

//mkunion-extension:visitor

type TreeVisitor interface {
	VisitBranch(v *Branch) any
	VisitLeaf(v *Leaf) any
	VisitK(v *K) any
}

type Tree interface {
	AcceptTree(g TreeVisitor) any
}

func (r *Branch) AcceptTree(v TreeVisitor) any { return v.VisitBranch(r) }
func (r *Leaf) AcceptTree(v TreeVisitor) any   { return v.VisitLeaf(r) }
func (r *K) AcceptTree(v TreeVisitor) any      { return v.VisitK(r) }

var (
	_ Tree = (*Branch)(nil)
	_ Tree = (*Leaf)(nil)
	_ Tree = (*K)(nil)
)

func MatchTree[TOut any](
	x Tree,
	f1 func(x *Branch) TOut,
	f2 func(x *Leaf) TOut,
	f3 func(x *K) TOut,
	df func(x Tree) TOut,
) TOut {
	return f.Match3(x, f1, f2, f3, df)
}

func MatchTreeR2[TOut1, TOut2 any](
	x Tree,
	f1 func(x *Branch) (TOut1, TOut2),
	f2 func(x *Leaf) (TOut1, TOut2),
	f3 func(x *K) (TOut1, TOut2),
	df func(x Tree) (TOut1, TOut2),
) (TOut1, TOut2) {
	return f.Match3R2(x, f1, f2, f3, df)
}

func MustMatchTree[TOut any](
	x Tree,
	f1 func(x *Branch) TOut,
	f2 func(x *Leaf) TOut,
	f3 func(x *K) TOut,
) TOut {
	return f.MustMatch3(x, f1, f2, f3)
}

func MustMatchTreeR0(
	x Tree,
	f1 func(x *Branch),
	f2 func(x *Leaf),
	f3 func(x *K),
) {
	f.MustMatch3R0(x, f1, f2, f3)
}

func MustMatchTreeR2[TOut1, TOut2 any](
	x Tree,
	f1 func(x *Branch) (TOut1, TOut2),
	f2 func(x *Leaf) (TOut1, TOut2),
	f3 func(x *K) (TOut1, TOut2),
) (TOut1, TOut2) {
	return f.MustMatch3R2(x, f1, f2, f3)
}

// mkunion-extension:reducer_dfs
type (
	TreeReducer[A any] interface {
		ReduceBranch(x *Branch, agg A) (result A, stop bool)
		ReduceLeaf(x *Leaf, agg A) (result A, stop bool)
		ReduceK(x *K, agg A) (result A, stop bool)
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
	if _ = v.Lit.AcceptTree(d); d.stop {
		return nil
	}
	for idx := range v.List {
		if _ = v.List[idx].AcceptTree(d); d.stop {
			return nil
		}
	}
	for idx, _ := range v.Map {
		if _ = v.Map[idx].AcceptTree(d); d.stop {
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

func (d *TreeDepthFirstVisitor[A]) VisitK(v *K) any {
	d.result, d.stop = d.reduce.ReduceK(v, d.result)
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

	_ = v.AcceptTree(reducer)

	return reducer.result
}

// mkunion-extension:reducer_bfs
var _ TreeVisitor = (*TreeBreadthFirstVisitor[any])(nil)

type TreeBreadthFirstVisitor[A any] struct {
	stop   bool
	result A
	reduce TreeReducer[A]

	queue         []Tree
	visited       map[Tree]bool
	shouldExecute map[Tree]bool
}

func (d *TreeBreadthFirstVisitor[A]) VisitBranch(v *Branch) any {
	d.queue = append(d.queue, v)
	d.queue = append(d.queue, v.Lit)
	for idx := range v.List {
		d.queue = append(d.queue, v.List[idx])
	}
	for idx, _ := range v.Map {
		d.queue = append(d.queue, v.Map[idx])
	}

	if d.shouldExecute[v] {
		d.shouldExecute[v] = false
		d.result, d.stop = d.reduce.ReduceBranch(v, d.result)
	} else {
		d.execute()
	}
	return nil
}

func (d *TreeBreadthFirstVisitor[A]) VisitLeaf(v *Leaf) any {
	d.queue = append(d.queue, v)

	if d.shouldExecute[v] {
		d.shouldExecute[v] = false
		d.result, d.stop = d.reduce.ReduceLeaf(v, d.result)
	} else {
		d.execute()
	}
	return nil
}

func (d *TreeBreadthFirstVisitor[A]) VisitK(v *K) any {
	d.queue = append(d.queue, v)

	if d.shouldExecute[v] {
		d.shouldExecute[v] = false
		d.result, d.stop = d.reduce.ReduceK(v, d.result)
	} else {
		d.execute()
	}
	return nil
}

func (d *TreeBreadthFirstVisitor[A]) execute() {
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
		i.AcceptTree(d)
	}

	return
}

func (d *TreeBreadthFirstVisitor[A]) pop() Tree {
	i := d.queue[0]
	d.queue = d.queue[1:]
	return i
}

func ReduceTreeBreadthFirst[A any](r TreeReducer[A], v Tree, init A) A {
	reducer := &TreeBreadthFirstVisitor[A]{
		result:        init,
		reduce:        r,
		queue:         []Tree{v},
		visited:       make(map[Tree]bool),
		shouldExecute: make(map[Tree]bool),
	}

	_ = v.AcceptTree(reducer)

	return reducer.result
}

// mkunion-extension:default_reducer
var _ TreeReducer[any] = (*TreeDefaultReduction[any])(nil)

type (
	TreeDefaultReduction[A any] struct {
		PanicOnFallback      bool
		DefaultStopReduction bool
		OnBranch             func(x *Branch, agg A) (result A, stop bool)
		OnLeaf               func(x *Leaf, agg A) (result A, stop bool)
		OnK                  func(x *K, agg A) (result A, stop bool)
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
		panic("no fallback allowed on undefined ReduceLeaf")
	}
	return agg, t.DefaultStopReduction
}

func (t *TreeDefaultReduction[A]) ReduceK(x *K, agg A) (result A, stop bool) {
	if t.OnK != nil {
		return t.OnK(x, agg)
	}
	if t.PanicOnFallback {
		panic("no fallback allowed on undefined ReduceK")
	}
	return agg, t.DefaultStopReduction
}

// mkunion-extension:default_visitor
type TreeDefaultVisitor[A any] struct {
	Default  A
	OnBranch func(x *Branch) A
	OnLeaf   func(x *Leaf) A
	OnK      func(x *K) A
}

func (t *TreeDefaultVisitor[A]) VisitBranch(v *Branch) any {
	if t.OnBranch != nil {
		return t.OnBranch(v)
	}
	return t.Default
}
func (t *TreeDefaultVisitor[A]) VisitLeaf(v *Leaf) any {
	if t.OnLeaf != nil {
		return t.OnLeaf(v)
	}
	return t.Default
}
func (t *TreeDefaultVisitor[A]) VisitK(v *K) any {
	if t.OnK != nil {
		return t.OnK(v)
	}
	return t.Default
}

// mkunion-extension:schema
func init() {
	schema.RegisterUnionTypes(TreeSchemaDef())
}

func TreeSchemaDef() *schema.UnionVariants[Tree] {
	return schema.MustDefineUnion[Tree](
		new(Branch),
		new(Leaf),
		new(K),
	)
}

// mkunion-extension:shape
func TreeShape() shape.Shape {
	return &shape.UnionLike{
		Name:          "Tree",
		PkgName:       "testutils",
		PkgImportName: "github.com/widmogrod/mkunion/x/generators/testutils",
		Variant: []shape.Shape{
			BranchShape(),
			LeafShape(),
			KShape(),
		},
	}
}

func BranchShape() shape.Shape {
	return &shape.StructLike{
		Name:          "Branch",
		PkgName:       "testutils",
		PkgImportName: "github.com/widmogrod/mkunion/x/generators/testutils",
		Fields: []*shape.FieldLike{
			{
				Name: "Lit",
				Type: &shape.RefName{
					Name:          "Tree",
					PkgName:       "testutils",
					PkgImportName: "github.com/widmogrod/mkunion/x/generators/testutils",
				},
			},
			{
				Name: "List",
				Type: &shape.ListLike{
					Element: &shape.RefName{
						Name:          "Tree",
						PkgName:       "testutils",
						PkgImportName: "github.com/widmogrod/mkunion/x/generators/testutils",
					},
					ElementIsPointer: false,
				},
			},
			{
				Name: "Map",
				Type: &shape.MapLike{
					Key:          &shape.StringLike{},
					KeyIsPointer: false,
					Val: &shape.RefName{
						Name:          "Tree",
						PkgName:       "testutils",
						PkgImportName: "github.com/widmogrod/mkunion/x/generators/testutils",
					},
					ValIsPointer: false,
				},
			},
		},
	}
}

func LeafShape() shape.Shape {
	return &shape.StructLike{
		Name:          "Leaf",
		PkgName:       "testutils",
		PkgImportName: "github.com/widmogrod/mkunion/x/generators/testutils",
		Fields: []*shape.FieldLike{
			{
				Name: "Value",
				Type: &shape.NumberLike{
					Kind: &shape.Int64{},
				},
			},
		},
	}
}

func KShape() shape.Shape {
	return &shape.StringLike{
		Named: &shape.Named{
			Name:          "K",
			PkgName:       "testutils",
			PkgImportName: "github.com/widmogrod/mkunion/x/generators/testutils",
		},
	}
}

// mkunion-extension:json
type TreeUnionJSON struct {
	Type   string          `json:"$type,omitempty"`
	Branch json.RawMessage `json:"testutils.Branch,omitempty"`
	Leaf   json.RawMessage `json:"testutils.Leaf,omitempty"`
	K      json.RawMessage `json:"testutils.K,omitempty"`
}

func TreeFromJSON(x []byte) (Tree, error) {
	var data TreeUnionJSON
	err := json.Unmarshal(x, &data)
	if err != nil {
		return nil, err
	}

	switch data.Type {
	case "testutils.Branch":
		return BranchFromJSON(data.Branch)
	case "testutils.Leaf":
		return LeafFromJSON(data.Leaf)
	case "testutils.K":
		return KFromJSON(data.K)
	}

	if data.Branch != nil {
		return BranchFromJSON(data.Branch)
	} else if data.Leaf != nil {
		return LeafFromJSON(data.Leaf)
	} else if data.K != nil {
		return KFromJSON(data.K)
	}

	return nil, fmt.Errorf("unknown type %s", data.Type)
}

func TreeToJSON(x Tree) ([]byte, error) {
	if x == nil {
		return nil, nil
	}
	return MustMatchTreeR2(
		x,
		func(x *Branch) ([]byte, error) {
			body, err := BranchToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(TreeUnionJSON{
				Type:   "testutils.Branch",
				Branch: body,
			})
		},
		func(x *Leaf) ([]byte, error) {
			body, err := LeafToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(TreeUnionJSON{
				Type: "testutils.Leaf",
				Leaf: body,
			})
		},
		func(x *K) ([]byte, error) {
			body, err := KToJSON(x)
			if err != nil {
				return nil, err
			}

			return json.Marshal(TreeUnionJSON{
				Type: "testutils.K",
				K:    body,
			})
		},
	)
}

func BranchFromJSON(x []byte) (*Branch, error) {
	var result *Branch = new(Branch)
	// if is Struct
	err := shared.JSONParseObject(x, func(key string, value []byte) error {
		switch key {
		case "Lit":
			res, err := TreeFromJSON(value)
			if err != nil {
				return fmt.Errorf("testutils._FromJSON: field Tree %w", err)
			}
			result.Lit = res
			return nil
		case "List":
			res, err := shared.JSONToListWithDeserializer(value, result.List, TreeFromJSON)
			if err != nil {
				return fmt.Errorf("testutils._FromJSON: field Tree %w", err)
			}
			result.List = res
			return nil
		case "Map":
			res, err := shared.JSONToMapWithDeserializer(value, result.Map, TreeFromJSON)
			if err != nil {
				return fmt.Errorf("testutils._FromJSON: field Tree %w", err)
			}
			result.Map = res
			return nil
		}

		return fmt.Errorf("testutils.BranchFromJSON: unknown key %s", key)
	})

	return result, err
}

func BranchToJSON(x *Branch) ([]byte, error) {
	field_Lit, err := TreeToJSON(x.Lit)
	if err != nil {
		return nil, err
	}
	field_List, err := shared.JSONListFromSerializer(x.List, TreeToJSON)
	if err != nil {
		return nil, err
	}
	field_Map, err := shared.JSONMapFromSerializer(x.Map, TreeToJSON)
	if err != nil {
		return nil, err
	}
	return json.Marshal(map[string]json.RawMessage{
		"Lit":  field_Lit,
		"List": field_List,
		"Map":  field_Map,
	})
}
func (self *Branch) MarshalJSON() ([]byte, error) {
	return BranchToJSON(self)
}

func (self *Branch) UnmarshalJSON(x []byte) error {
	n, err := BranchFromJSON(x)
	if err != nil {
		return err
	}
	*self = *n
	return nil
}

func LeafFromJSON(x []byte) (*Leaf, error) {
	var result *Leaf = new(Leaf)
	// if is Struct
	err := shared.JSONParseObject(x, func(key string, value []byte) error {
		switch key {
		case "Value":
			return json.Unmarshal(value, &result.Value)
		}

		return fmt.Errorf("testutils.LeafFromJSON: unknown key %s", key)
	})

	return result, err
}

func LeafToJSON(x *Leaf) ([]byte, error) {
	field_Value, err := json.Marshal(x.Value)
	if err != nil {
		return nil, err
	}
	return json.Marshal(map[string]json.RawMessage{
		"Value": field_Value,
	})
}
func (self *Leaf) MarshalJSON() ([]byte, error) {
	return LeafToJSON(self)
}

func (self *Leaf) UnmarshalJSON(x []byte) error {
	n, err := LeafFromJSON(x)
	if err != nil {
		return err
	}
	*self = *n
	return nil
}

func KFromJSON(x []byte) (*K, error) {
	var result *K = new(K)
	err := json.Unmarshal(x, result)

	return result, err
}

func KToJSON(x *K) ([]byte, error) {
	return json.Marshal(x)
}
