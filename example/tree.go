package example

//go:generate go run ../cmd/mkunion/main.go

//go:tag mkunion:"Tree"
type (
	Branch[A any] struct{ L, R Tree[A] }
	Leaf[A any]   struct{ Value A }
)

func ReduceTree[A, B any](x Tree[A], f func(A, B) B, init B) B {
	return MatchTreeR1(
		x,
		func(x *Branch[A]) B {
			return ReduceTree(x.R, f, ReduceTree(x.L, f, init))
		}, func(x *Leaf[A]) B {
			return f(x.Value, init)
		},
	)
}

// Example of using Visitor pattern and interface, that was generated by mkunion.
var _ TreeVisitor[int] = (*sumVisitor)(nil)

type sumVisitor struct{}

func (s sumVisitor) VisitBranch(v *Branch[int]) any {
	return v.L.AcceptTree(s).(int) + v.R.AcceptTree(s).(int)
}

func (s sumVisitor) VisitLeaf(v *Leaf[int]) any {
	return v.Value
}

////go:generate go run ../cmd/mkunion/main.go match -name=MyTriesMatch
//type MyTriesMatch[T0, T1 Tree[A], A any] interface {
//	MatchLeafs(*Leaf[A], *Leaf[A])
//	MatchBranches(*Branch[A], any)
//	MatchMixed(any, any)
//}
