package testasset

import "time"

//go:generate go run ../../../cmd/mkunion/main.go --name=Example
//go:tag mkunion:"Example"
type (
	A struct {
		Name string `json:"name" desc:"Name of the person"`
	}
	B struct {
		Age int `json:"age"`
		A   *A
		T   *time.Time
	}
	C string
	D int64
	E float64
	F bool
	H map[string]Example
	I []Example
	J [2]string
	K A
	// L Example is not allowed, since Example is interface,
	// and interface cannot have methods implemented as Visitor pattern requires
	L = List
	//go:tag json:"m_list,omitempty"
	M List
	N time.Duration
	O ListOf[time.Duration]
	P ListOf2[ListOf[any], *ListOf2[int64, *time.Duration]]
)

// List is a list of elements
//
//go:tag serde:"json" json:"list,omitempty"
type List struct{}

//go:tag serde:"json" json:"list_of,omitempty"
type ListOf[T any] struct{}

// ListOf2 is a list of 2 elements
//
//go:tag serde:"json" json:"list_of_2,omitempty"
type ListOf2[T1, T2 any] struct{}
