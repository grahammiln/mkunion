// Code generated by mkunion. DO NOT EDIT.
package projection

import (
	"github.com/widmogrod/mkunion/x/shape"
)

func init() {
	shape.Register(DataShape())
	shape.Register(EitherShape())
	shape.Register(EventTimeShape())
	shape.Register(InMemoryJoinContextShape())
	shape.Register(LeftShape())
	shape.Register(PushAndPullInMemoryContextShape())
	shape.Register(RecordShape())
	shape.Register(RightShape())
	shape.Register(SimulateProblemShape())
	shape.Register(SnapshotStateShape())
	shape.Register(SnapshotStoreShape())
	shape.Register(WatermarkShape())
}

//shape:shape

func DataShape() shape.Shape {
	return &shape.UnionLike{
		Name:          "Data",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/projection",
		TypeParams: []shape.TypeParam{
			shape.TypeParam{
				Name: "A",
				Type: &shape.Any{},
			},
		},
		Variant: []shape.Shape{
			RecordShape(),
			WatermarkShape(),
		},
	}
}

func RecordShape() shape.Shape {
	return &shape.StructLike{
		Name:          "Record",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/projection",
		TypeParams: []shape.TypeParam{
			shape.TypeParam{
				Name: "A",
				Type: &shape.Any{},
			},
		},
		Fields: []*shape.FieldLike{
			{
				Name: "Key",
				Type: &shape.PrimitiveLike{Kind: &shape.StringLike{}},
			},
			{
				Name: "Data",
				Type: &shape.RefName{
					Name:          "A",
					PkgName:       "",
					PkgImportName: "",
				},
			},
			{
				Name: "EventTime",
				Type: &shape.RefName{
					Name:          "EventTime",
					PkgName:       "projection",
					PkgImportName: "github.com/widmogrod/mkunion/x/projection",
				},
			},
		},
	}
}

func WatermarkShape() shape.Shape {
	return &shape.StructLike{
		Name:          "Watermark",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/projection",
		TypeParams: []shape.TypeParam{
			shape.TypeParam{
				Name: "A",
				Type: &shape.Any{},
			},
		},
		Fields: []*shape.FieldLike{
			{
				Name: "Key",
				Type: &shape.PrimitiveLike{Kind: &shape.StringLike{}},
			},
			{
				Name: "EventTime",
				Type: &shape.RefName{
					Name:          "EventTime",
					PkgName:       "projection",
					PkgImportName: "github.com/widmogrod/mkunion/x/projection",
				},
			},
		},
	}
}

//shape:shape
func EventTimeShape() shape.Shape {
	return &shape.AliasLike{
		Name:          "EventTime",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/projection",
		IsAlias:       true,
		Type: &shape.PrimitiveLike{
			Kind: &shape.NumberLike{
				Kind: &shape.Int64{},
			},
		},
	}
}

//shape:shape

func EitherShape() shape.Shape {
	return &shape.UnionLike{
		Name:          "Either",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/projection",
		TypeParams: []shape.TypeParam{
			shape.TypeParam{
				Name: "A",
				Type: &shape.Any{},
			},
			shape.TypeParam{
				Name: "B",
				Type: &shape.Any{},
			},
		},
		Variant: []shape.Shape{
			LeftShape(),
			RightShape(),
		},
	}
}

func LeftShape() shape.Shape {
	return &shape.StructLike{
		Name:          "Left",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/projection",
		TypeParams: []shape.TypeParam{
			shape.TypeParam{
				Name: "A",
				Type: &shape.Any{},
			},
			shape.TypeParam{
				Name: "B",
				Type: &shape.Any{},
			},
		},
		Fields: []*shape.FieldLike{
			{
				Name: "Left",
				Type: &shape.RefName{
					Name:          "A",
					PkgName:       "",
					PkgImportName: "",
				},
			},
		},
	}
}

func RightShape() shape.Shape {
	return &shape.StructLike{
		Name:          "Right",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/projection",
		TypeParams: []shape.TypeParam{
			shape.TypeParam{
				Name: "A",
				Type: &shape.Any{},
			},
			shape.TypeParam{
				Name: "B",
				Type: &shape.Any{},
			},
		},
		Fields: []*shape.FieldLike{
			{
				Name: "Right",
				Type: &shape.RefName{
					Name:          "B",
					PkgName:       "",
					PkgImportName: "",
				},
			},
		},
	}
}

//shape:shape
func InMemoryJoinContextShape() shape.Shape {
	return &shape.StructLike{
		Name:          "InMemoryJoinContext",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/projection",
		TypeParams: []shape.TypeParam{
			shape.TypeParam{
				Name: "A",
				Type: &shape.Any{},
			},
			shape.TypeParam{
				Name: "B",
				Type: &shape.Any{},
			},
			shape.TypeParam{
				Name: "C",
				Type: &shape.Any{},
			},
		},
	}
}

//shape:shape
func PushAndPullInMemoryContextShape() shape.Shape {
	return &shape.StructLike{
		Name:          "PushAndPullInMemoryContext",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/projection",
		TypeParams: []shape.TypeParam{
			shape.TypeParam{
				Name: "A",
				Type: &shape.Any{},
			},
			shape.TypeParam{
				Name: "B",
				Type: &shape.Any{},
			},
		},
	}
}

//shape:shape
func SimulateProblemShape() shape.Shape {
	return &shape.StructLike{
		Name:          "SimulateProblem",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/projection",
		Fields: []*shape.FieldLike{
			{
				Name: "ErrorOnPullIn",
				Type: &shape.Any{},
			},
			{
				Name: "ErrorOnPushOut",
				Type: &shape.Any{},
			},
		},
	}
}

//shape:shape
func SnapshotStateShape() shape.Shape {
	return &shape.StructLike{
		Name:          "SnapshotState",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/projection",
		Fields: []*shape.FieldLike{
			{
				Name: "ID",
				Type: &shape.PrimitiveLike{Kind: &shape.StringLike{}},
			},
			{
				Name: "Offset",
				Type: &shape.PointerLike{
					Type: &shape.RefName{
						Name:          "Offset",
						PkgName:       "stream",
						PkgImportName: "github.com/widmogrod/mkunion/x/stream",
					},
				},
			},
			{
				Name: "Completed",
				Type: &shape.PrimitiveLike{Kind: &shape.BooleanLike{}},
			},
		},
	}
}

//shape:shape
func SnapshotStoreShape() shape.Shape {
	return &shape.StructLike{
		Name:          "SnapshotStore",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/projection",
	}
}
