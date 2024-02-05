package shape

import (
	"fmt"
	"reflect"
	"strings"
)

func ToGoPkgName(x Shape) string {
	return MatchShapeR1(
		x,
		func(x *Any) string {
			return ""
		},
		func(x *RefName) string {
			return x.PkgName
		},
		func(x *PointerLike) string {
			return ToGoPkgName(x.Type)
		},
		func(x *AliasLike) string {
			return x.PkgName
		},
		func(x *PrimitiveLike) string {
			return ""
		},
		func(x *ListLike) string {
			return ""
		},
		func(x *MapLike) string {
			return ""
		},
		func(x *StructLike) string {
			return x.PkgName
		},
		func(x *UnionLike) string {
			return x.PkgName
		},
	)
}

func ToGoPkgImportName(x Shape) string {
	return MatchShapeR1(
		x,
		func(x *Any) string {
			return ""
		},
		func(x *RefName) string {
			return x.PkgImportName
		},
		func(x *PointerLike) string {
			return ToGoPkgImportName(x.Type)
		},
		func(x *AliasLike) string {
			return x.PkgImportName
		},
		func(x *PrimitiveLike) string {
			return ""
		},
		func(x *ListLike) string {
			return ""
		},
		func(x *MapLike) string {
			return ""
		},
		func(x *StructLike) string {
			return x.PkgImportName
		},
		func(x *UnionLike) string {
			return x.PkgImportName
		},
	)
}

type ToGoTypeNameOption string

func WithInstantiation() ToGoTypeNameOption {
	return "instantiate"
}

func WithPkgImportName() ToGoTypeNameOption {
	return usePkgImportName
}

const (
	rootPackage      = "root-package:"
	usePkgImportName = "usePkgImportName"
)

func WithRootPackage(pkgName string) ToGoTypeNameOption {
	result := rootPackage + pkgName
	return ToGoTypeNameOption(result)
}

func packageWrap(options []ToGoTypeNameOption, pkgName, pkgImportName, value string) string {
	if pkgName == "" {
		return value
	}

	useImportName := false
	for _, option := range options {
		if option == usePkgImportName {
			useImportName = true
		}
		if !strings.HasPrefix(string(option), rootPackage) {
			continue
		}

		rootPkgName := string(option)[len(rootPackage):]
		if pkgName == rootPkgName {
			return value
		}
	}

	if useImportName {
		return fmt.Sprintf("%s.%s", pkgImportName, value)
	}

	return fmt.Sprintf("%s.%s", pkgName, value)
}

func shouldInstantiate(options []ToGoTypeNameOption) bool {
	for _, x := range options {
		if x == "instantiate" {
			return true
		}
	}
	return false
}

// ToGoFullTypeNameFromReflect returns full type name with package name
// example:
//
//	string
//	github.com/widmogrod/mkunion/x/schema.Any
//	github.com/widmogrod/mkunion/x/projection.Record
func ToGoFullTypeNameFromReflect(x reflect.Type) string {
	if x.Kind() == reflect.Ptr {
		x = x.Elem()
	}

	if x.PkgPath() == "" {
		return x.String()
	}

	return fmt.Sprintf("%s.%s", x.PkgPath(), x.Name())
}

func ToGoTypeName(x Shape, options ...ToGoTypeNameOption) string {
	return MatchShapeR1(
		x,
		func(x *Any) string {
			return "any"
		},
		func(x *RefName) string {
			// this is when it's a parameter result or some other primitive type
			if x.PkgName == "" {
				return x.Name
			}

			result := packageWrap(options, x.PkgName, x.PkgImportName, x.Name)
			if len(x.Indexed) > 0 {
				names := make([]string, 0)
				for _, name := range x.Indexed {
					names = append(names, ToGoTypeName(name, options...))
				}

				result = fmt.Sprintf("%s[%s]", result, strings.Join(names, ","))
			}

			return result
		},
		func(x *PointerLike) string {
			return fmt.Sprintf("*%s", ToGoTypeName(x.Type, options...))
		},
		func(x *AliasLike) string {
			result := toGoNameWithTypeParams(x, options...)
			if result != "" {
				return result
			}

			return packageWrap(options, x.PkgName, x.PkgImportName, x.Name)
		},
		func(x *PrimitiveLike) string {
			return MatchPrimitiveKindR1(
				x.Kind,
				func(x *BooleanLike) string {
					return "bool"
				},
				func(x *StringLike) string {
					return "string"
				},
				func(x *NumberLike) string {
					return NumberKindToGoName(x.Kind)
				},
			)
		},
		func(x *ListLike) string {
			prefix := "[]"
			if x.ArrayLen != nil {
				prefix = fmt.Sprintf("[%d]", *x.ArrayLen)
			}
			return fmt.Sprintf("%s%s",
				prefix,
				ToGoTypeName(x.Element, options...),
			)
		},
		func(x *MapLike) string {
			return fmt.Sprintf("map[%s]%s",
				ToGoTypeName(x.Key, options...),
				ToGoTypeName(x.Val, options...),
			)
		},
		func(x *StructLike) string {
			result := toGoNameWithTypeParams(x, options...)
			if result != "" {
				return result
			}

			return packageWrap(options, x.PkgName, x.PkgImportName, x.Name)
		},
		func(x *UnionLike) string {
			result := toGoNameWithTypeParams(x, options...)
			if result != "" {
				return result
			}

			return packageWrap(options, x.PkgName, x.PkgImportName, x.Name)
		},
	)
}

func toGoNameWithTypeParams(x Shape, options ...ToGoTypeNameOption) string {
	var typeName string
	var pkgName string
	var pkgImportName string

	switch x := x.(type) {
	case *StructLike:
		typeName = x.Name
		pkgName = x.PkgName
		pkgImportName = x.PkgImportName

	case *UnionLike:
		typeName = x.Name
		pkgName = x.PkgName
		pkgImportName = x.PkgImportName

	case *AliasLike:
		typeName = x.Name
		pkgName = x.PkgName
		pkgImportName = x.PkgImportName

	default:
		return ""
	}

	if shouldInstantiate(options) {
		typeParams := ToGoTypeParamsTypes(x)
		if len(typeParams) > 0 {
			instantiatedTypeParams := make([]string, len(typeParams))
			for i, y := range typeParams {
				instantiatedTypeParams[i] = ToGoTypeName(y, options...)
			}

			result := fmt.Sprintf("%s[%s]",
				typeName,
				strings.Join(instantiatedTypeParams, ","),
			)

			result = packageWrap(options, pkgName, pkgImportName, result)
			return result
		}
	} else {
		typeParams := ToGoTypeParamsNames(x)
		if len(typeParams) > 0 {
			result := fmt.Sprintf("%s[%s]",
				typeName,
				strings.Join(typeParams, ","),
			)

			result = packageWrap(options, pkgName, pkgImportName, result)
			return result
		}
	}

	return ""
}

func ToGoTypeParamsNames(x Shape) []string {
	return MatchShapeR1(
		x,
		func(x *Any) []string {
			return nil
		},
		func(x *RefName) []string {
			return nil
		},
		func(x *PointerLike) []string {
			return ToGoTypeParamsNames(x.Type)
		},
		func(x *AliasLike) []string {
			typeParamsNames := make([]string, len(x.TypeParams))
			for i, y := range x.TypeParams {
				typeParamsNames[i] = y.Name
			}
			return typeParamsNames
		},
		func(x *PrimitiveLike) []string {
			return nil
		},
		func(x *ListLike) []string {
			return nil
		},
		func(x *MapLike) []string {
			return nil
		},
		func(x *StructLike) []string {
			typeParamsNames := make([]string, len(x.TypeParams))
			for i, y := range x.TypeParams {
				typeParamsNames[i] = y.Name
			}
			return typeParamsNames
		},
		func(x *UnionLike) []string {
			typeParamsNames := make([]string, len(x.TypeParams))
			for i, y := range x.TypeParams {
				typeParamsNames[i] = y.Name
			}
			return typeParamsNames
		},
	)
}

func ToGoTypeParamsTypes(x Shape) []Shape {
	return MatchShapeR1(
		x,
		func(x *Any) []Shape {
			return nil
		},
		func(x *RefName) []Shape {
			return nil
		},
		func(x *PointerLike) []Shape {
			return ToGoTypeParamsTypes(x.Type)
		},
		func(x *AliasLike) []Shape {
			typeParamsNames := make([]Shape, len(x.TypeParams))
			for i, y := range x.TypeParams {
				typeParamsNames[i] = y.Type
			}
			return typeParamsNames
		},
		func(x *PrimitiveLike) []Shape {
			return nil
		},
		func(x *ListLike) []Shape {
			return nil
		},
		func(x *MapLike) []Shape {
			return nil
		},
		func(x *StructLike) []Shape {
			typeParamsNames := make([]Shape, len(x.TypeParams))
			for i, y := range x.TypeParams {
				typeParamsNames[i] = y.Type
			}
			return typeParamsNames
		},
		func(x *UnionLike) []Shape {
			typeParamsNames := make([]Shape, len(x.TypeParams))
			for i, y := range x.TypeParams {
				typeParamsNames[i] = y.Type
			}
			return typeParamsNames
		},
	)
}

func ExtractTypeParams(x Shape) []TypeParam {
	return MatchShapeR1(
		x,
		func(x *Any) []TypeParam {
			return nil
		},
		func(x *RefName) []TypeParam {
			return nil
		},
		func(x *PointerLike) []TypeParam {
			return ExtractTypeParams(x.Type)
		},
		func(x *AliasLike) []TypeParam {
			return x.TypeParams
		},
		func(x *PrimitiveLike) []TypeParam {
			return nil
		},
		func(x *ListLike) []TypeParam {
			return nil
		},
		func(x *MapLike) []TypeParam {
			return nil
		},
		func(x *StructLike) []TypeParam {
			return x.TypeParams
		},
		func(x *UnionLike) []TypeParam {
			return x.TypeParams
		},
	)
}

func ExtractPkgImportNames(x Shape) map[string]string {
	return MatchShapeR1(
		x,
		func(y *Any) map[string]string {
			return nil
		},
		func(y *RefName) map[string]string {
			result := make(map[string]string)
			if y.PkgName != "" {
				result[y.PkgName] = y.PkgImportName
			}

			for _, x := range y.Indexed {
				result = joinMaps(result, ExtractPkgImportNames(x))
			}

			return result
		},
		func(x *PointerLike) map[string]string {
			result := make(map[string]string)
			result = joinMaps(result, ExtractPkgImportNames(x.Type))
			return result
		},
		func(x *AliasLike) map[string]string {
			result := make(map[string]string)
			if x.PkgName != "" {
				result[x.PkgName] = x.PkgImportName
			}

			result = joinMaps(result, ExtractPkgImportNames(x.Type))
			for _, y := range x.TypeParams {
				result = joinMaps(result, ExtractPkgImportNames(y.Type))
			}

			return result
		},
		func(x *PrimitiveLike) map[string]string {
			return nil
		},
		func(x *ListLike) map[string]string {
			result := make(map[string]string)
			result = joinMaps(result, ExtractPkgImportNames(x.Element))
			return result
		},
		func(x *MapLike) map[string]string {
			result := make(map[string]string)
			result = joinMaps(result, ExtractPkgImportNames(x.Key))
			result = joinMaps(result, ExtractPkgImportNames(x.Val))
			return result
		},
		func(x *StructLike) map[string]string {
			result := make(map[string]string)
			if x.PkgName != "" {
				result[x.PkgName] = x.PkgImportName
			}

			for _, y := range x.TypeParams {
				result = joinMaps(result, ExtractPkgImportNames(y.Type))
			}

			for _, y := range x.Fields {
				result = joinMaps(result, ExtractPkgImportNames(y.Type))
			}

			return result

		},
		func(x *UnionLike) map[string]string {
			result := make(map[string]string)
			if x.PkgName != "" {
				result[x.PkgName] = x.PkgImportName
			}

			for _, y := range x.Variant {
				result = joinMaps(result, ExtractPkgImportNames(y))
			}

			return result
		},
	)
}

func joinMaps(x map[string]string, maps ...map[string]string) map[string]string {
	for _, m := range maps {
		for k, v := range m {
			x[k] = v
		}
	}
	return x
}

func Name(x Shape) string {
	return MatchShapeR1(
		x,
		func(x *Any) string {
			return ""
		},
		func(x *RefName) string {
			return x.Name
		},
		func(x *PointerLike) string {
			return Name(x.Type)
		},
		func(x *AliasLike) string {
			return x.Name
		},
		func(x *PrimitiveLike) string {
			return ""
		},
		func(x *ListLike) string {
			return ""
		},
		func(x *MapLike) string {
			return ""
		},
		func(x *StructLike) string {
			return x.Name
		},
		func(x *UnionLike) string {
			return x.Name
		},
	)
}

func NameToPrimitiveShape(name string) Shape {
	switch name {
	case "any":
		return &Any{}
	case "string":
		return &PrimitiveLike{Kind: &StringLike{}}
	case "bool":
		return &PrimitiveLike{Kind: &BooleanLike{}}
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float64", "float32",
		"byte", "rune":
		return &PrimitiveLike{
			Kind: &NumberLike{
				Kind: TypeStringToNumberKindMap[name],
			},
		}
	}

	return nil
}
