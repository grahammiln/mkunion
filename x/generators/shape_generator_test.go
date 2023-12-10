package generators

import (
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/widmogrod/mkunion/x/shape"
	"testing"
)

func TestShapeGenerator(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	inferred, err := shape.InferFromFile("testutils/tree.go")
	assert.NoError(t, err)

	g := NewShapeGenerator(
		inferred.RetrieveUnion("Tree"),
		NewHelper(WithPackageName("testutils")),
	)

	result, err := g.Generate()
	assert.NoError(t, err)
	assert.Equal(t, `// Code generated by mkunion. DO NOT EDIT.
package testutils

import "github.com/widmogrod/mkunion/x/shape"

func TreeShape() shape.Shape {
	return &shape.UnionLike{
		Name: "Tree",
		PkgName: "testutils",
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
		Name: "Branch",
		PkgName: "testutils",
		PkgImportName: "github.com/widmogrod/mkunion/x/generators/testutils",
		Fields: []*shape.FieldLike{
			{
				Name: "Lit",
				Type: &shape.RefName{
					Name: "Tree",
					PkgName: "testutils",
					PkgImportName: "github.com/widmogrod/mkunion/x/generators/testutils",
				},
			},
			{
				Name: "List",
				Type: &shape.ListLike{
					Element: &shape.RefName{
						Name: "Tree",
						PkgName: "testutils",
						PkgImportName: "github.com/widmogrod/mkunion/x/generators/testutils",
					},
					ElementIsPointer: false,
				},
			},
			{
				Name: "Map",
				Type: &shape.MapLike{
					Key: &shape.StringLike{},
					KeyIsPointer: false,
					Val: &shape.RefName{
						Name: "Tree",
						PkgName: "testutils",
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
		Name: "Leaf",
		PkgName: "testutils",
		PkgImportName: "github.com/widmogrod/mkunion/x/generators/testutils",
		Fields: []*shape.FieldLike{
			{
				Name: "Value",
				Type: &shape.NumberLike{},
			},
		},
	}
}

func KShape() shape.Shape {
	return &shape.StringLike{
		Named: &shape.Named{
			Name: "K",
			PkgName: "testutils",
			PkgImportName: "github.com/widmogrod/mkunion/x/generators/testutils",
		},
	}
}
`, string(result))
}
