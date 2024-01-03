package testutil

import "github.com/widmogrod/mkunion/x/schema"

//go:generate go run ../../../../cmd/mkunion/main.go serde

//go:tag serde:"json"
type SampleStruct struct {
	ID      string
	Age     int
	Friends []SampleStruct
	Tree    Treeish
	Visible bool
}

//go:generate go run ../../../../cmd/mkunion/main.go -name=Treeish
type (
	Branch struct {
		Name        string
		Left, Right Treeish
	}
	Leaf struct {
		Value schema.Schema
	}
)
