package generators

import (
	"fmt"
	"github.com/widmogrod/mkunion/x/shape"
)

func PtrStr(x string) *string {
	return &x
}

type Branching struct {
	Lit  *string
	List *string
	Map  *string
}

func AdaptUnionToOldVersionOfGenerator(union shape.UnionLike) ([]string, map[string][]Branching) {
	types := []string{}
	branches := map[string][]Branching{}

	for _, v := range union.Variant {
		typeName := shape.MustMatchShape(
			v,
			func(x *shape.Any) string {
				panic(fmt.Errorf("generators.AdaptUnionToOldVersionOfGenerator: %T not supported", x))
			},
			func(x *shape.RefName) string {
				return x.Name
			},
			func(x *shape.BooleanLike) string {
				if shape.IsNamed(x) {
					return x.Named.Name
				}

				panic(fmt.Errorf("generators.AdaptUnionToOldVersionOfGenerator: expects only named shape, but given %#v", x))
			},
			func(x *shape.StringLike) string {
				if shape.IsNamed(x) {
					return x.Named.Name
				}

				panic(fmt.Errorf("generators.AdaptUnionToOldVersionOfGenerator: expects only named shape, but given %#v", x))
			},
			func(x *shape.NumberLike) string {
				if shape.IsNamed(x) {
					return x.Named.Name
				}

				panic(fmt.Errorf("generators.AdaptUnionToOldVersionOfGenerator: expects only named shape, but given %#v", x))
			},
			func(x *shape.ListLike) string {
				if shape.IsNamed(x) {
					return x.Named.Name
				}

				panic(fmt.Errorf("generators.AdaptUnionToOldVersionOfGenerator: expects only named shape, but given %#v", x))
			},
			func(x *shape.MapLike) string {
				if shape.IsNamed(x) {
					return x.Named.Name
				}

				panic(fmt.Errorf("generators.AdaptUnionToOldVersionOfGenerator: expects only named shape, but given %#v", x))
			},
			func(x *shape.StructLike) string {
				return x.Name
			},
			func(x *shape.UnionLike) string {
				return x.Name
			},
		)

		types = append(types, typeName)
		branches[typeName] = []Branching{}

		branches[typeName] = shape.MustMatchShape(
			v,
			func(x *shape.Any) []Branching {
				return []Branching{}
			},
			func(x *shape.RefName) []Branching {
				return []Branching{}
			},
			func(x *shape.BooleanLike) []Branching {
				return []Branching{}
			},
			func(x *shape.StringLike) []Branching {
				return []Branching{}
			},
			func(x *shape.NumberLike) []Branching {
				return []Branching{}
			},
			func(x *shape.ListLike) []Branching {
				return []Branching{}
			},
			func(x *shape.MapLike) []Branching {
				return []Branching{}
			},
			func(x *shape.StructLike) []Branching {
				result := []Branching{}
				for _, field := range x.Fields {
					switch y := field.Type.(type) {
					case *shape.RefName:
						if y.PkgImportName == union.PkgImportName &&
							y.Name == union.Name {
							result = append(result, Branching{
								Lit: PtrStr(field.Name),
							})
						}
					case *shape.ListLike:
						ref, ok := y.Element.(*shape.RefName)
						if ok &&
							ref.PkgImportName == union.PkgImportName &&
							ref.Name == union.Name {
							result = append(result, Branching{
								List: PtrStr(field.Name),
							})
						}
					case *shape.MapLike:
						ref, ok := y.Val.(*shape.RefName)
						if ok &&
							ref.PkgImportName == union.PkgImportName &&
							ref.Name == union.Name {
							result = append(result, Branching{
								Map: PtrStr(field.Name),
							})
						}
					}
				}

				return result
			},
			func(x *shape.UnionLike) []Branching {
				return []Branching{}
			},
		)
	}
	return types, branches
}
