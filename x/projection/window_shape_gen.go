// Code generated by mkunion. DO NOT EDIT.
package projection

import (
	"github.com/widmogrod/mkunion/x/shape"
)

func init() {
	shape.Register(FixedWindowShape())
	shape.Register(SessionWindowShape())
	shape.Register(SlidingWindowShape())
	shape.Register(WindowDescriptionShape())
	shape.Register(WindowIDShape())
	shape.Register(WindowInMemoryStoreShape())
	shape.Register(WindowKeyShape())
	shape.Register(WindowRecordShape())
	shape.Register(WindowShape())
	shape.Register(WindowSnapshotStateShape())
}

//shape:shape

func WindowDescriptionShape() shape.Shape {
	return &shape.UnionLike{
		Name:          "WindowDescription",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/projection",
		Variant: []shape.Shape{
			SessionWindowShape(),
			SlidingWindowShape(),
			FixedWindowShape(),
		},
	}
}

func SessionWindowShape() shape.Shape {
	return &shape.StructLike{
		Name:          "SessionWindow",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/projection",
		Fields: []*shape.FieldLike{
			{
				Name: "GapDuration",
				Type: &shape.RefName{
					Name:          "Duration",
					PkgName:       "time",
					PkgImportName: "time",
				},
			},
		},
	}
}

func SlidingWindowShape() shape.Shape {
	return &shape.StructLike{
		Name:          "SlidingWindow",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/projection",
		Fields: []*shape.FieldLike{
			{
				Name: "Width",
				Type: &shape.RefName{
					Name:          "Duration",
					PkgName:       "time",
					PkgImportName: "time",
				},
			},
			{
				Name: "Period",
				Type: &shape.RefName{
					Name:          "Duration",
					PkgName:       "time",
					PkgImportName: "time",
				},
			},
		},
	}
}

func FixedWindowShape() shape.Shape {
	return &shape.StructLike{
		Name:          "FixedWindow",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/projection",
		Fields: []*shape.FieldLike{
			{
				Name: "Width",
				Type: &shape.RefName{
					Name:          "Duration",
					PkgName:       "time",
					PkgImportName: "time",
				},
			},
		},
	}
}

//shape:shape
func WindowShape() shape.Shape {
	return &shape.StructLike{
		Name:          "Window",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/projection",
		Fields: []*shape.FieldLike{
			{
				Name: "Start",
				Type: &shape.PrimitiveLike{
					Kind: &shape.NumberLike{
						Kind: &shape.Int64{},
					},
				},
			},
			{
				Name: "End",
				Type: &shape.PrimitiveLike{
					Kind: &shape.NumberLike{
						Kind: &shape.Int64{},
					},
				},
			},
		},
	}
}

//shape:shape
func WindowIDShape() shape.Shape {
	return &shape.AliasLike{
		Name:          "WindowID",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/projection",
		IsAlias:       true,
		Type:          &shape.PrimitiveLike{Kind: &shape.StringLike{}},
	}
}

//shape:shape
func WindowInMemoryStoreShape() shape.Shape {
	return &shape.StructLike{
		Name:          "WindowInMemoryStore",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/projection",
		TypeParams: []shape.TypeParam{
			shape.TypeParam{
				Name: "A",
				Type: &shape.Any{},
			},
		},
	}
}

//shape:shape
func WindowKeyShape() shape.Shape {
	return &shape.AliasLike{
		Name:          "WindowKey",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/projection",
		IsAlias:       true,
		Type:          &shape.PrimitiveLike{Kind: &shape.StringLike{}},
	}
}

//shape:shape
func WindowRecordShape() shape.Shape {
	return &shape.StructLike{
		Name:          "WindowRecord",
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
				Type: &shape.RefName{
					Name:          "WindowKey",
					PkgName:       "projection",
					PkgImportName: "github.com/widmogrod/mkunion/x/projection",
				},
			},
			{
				Name: "Window",
				Type: &shape.PointerLike{
					Type: &shape.RefName{
						Name:          "Window",
						PkgName:       "projection",
						PkgImportName: "github.com/widmogrod/mkunion/x/projection",
					},
				},
			},
			{
				Name: "Record",
				Type: &shape.RefName{
					Name:          "A",
					PkgName:       "",
					PkgImportName: "",
				},
			},
		},
	}
}

//shape:shape
func WindowSnapshotStateShape() shape.Shape {
	return &shape.StructLike{
		Name:          "WindowSnapshotState",
		PkgName:       "projection",
		PkgImportName: "github.com/widmogrod/mkunion/x/projection",
	}
}
