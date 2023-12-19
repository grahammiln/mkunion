package testutil

import "github.com/widmogrod/mkunion/x/schema"

//go:generate go run ../../../cmd/mkunion/main.go -name=Example -skip-extension=shape,schema
type (
	ExampleOne struct {
		OneValue string
	}
	ExampleTwo struct {
		TwoData schema.Schema
		TwoNext Example
	}
	ExampleTree struct {
		Items   []Example
		Schemas []schema.Schema
	}
)
