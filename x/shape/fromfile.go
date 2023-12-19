package shape

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/widmogrod/mkunion/x/shared"
	"go/ast"
	"go/parser"
	"go/token"
	"golang.org/x/mod/modfile"
	"io"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

func InferFromFile(filename string) (*InferredInfo, error) {
	if !path.IsAbs(filename) {
		cwd, _ := os.Getwd()
		filename = path.Join(cwd, filename)
	}

	result := &InferredInfo{
		pkgImportName:        tryToPkgImportName(filename),
		possibleVariantTypes: map[string][]string{},
		possibleTaggedTypes:  map[string]map[string]Tag{},
		shapes:               make(map[string]Shape),
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	ast.Walk(result, f)
	return result, nil
}

// tryToPkgImportName contains import name of the package
func tryToPkgImportName(filename string) string {
	var toadd []string
	for {
		filename = path.Dir(filename)
		if filename == "." || filename == "/" {
			log.Infof("shape.InferFromFile: could not find go.mod file in %s, returning empty pkg name", filename)
			return ""
		}

		modpath := path.Join(filename, "go.mod")
		if _, err := os.Stat(modpath); err == nil {
			f, err := os.Open(modpath)
			defer f.Close()

			data, err := io.ReadAll(f)
			if err != nil {
				log.Infof("shape.InferFromFile: could not read go.mod file in %s, returning empty pkg name. %s", filename, err.Error())
				return ""
			}

			parsed, err := modfile.Parse(modpath, data, nil)
			if err != nil {
				log.Infof("shape.InferFromFile: could not parse go.mod file in %s, returning empty pkg name. %s", filename, err.Error())
				return ""
			}

			if parsed.Module == nil {
				log.Infof("shape.InferFromFile: could not find module name in go.mod file in %s, returning empty pkg name", filename)
				return ""
			}

			return path.Join(append([]string{parsed.Module.Mod.Path}, toadd...)...)
		}

		toadd = append([]string{path.Base(filename)}, toadd...)
	}
}

var (
	matchGoGenerateExtractUnionName = regexp.MustCompile(`go:generate .* -{1,2}name=*\s*(\w+)`)
)

type InferredInfo struct {
	packageName                string
	pkgImportName              string
	possibleVariantTypes       map[string][]string
	shapes                     map[string]Shape
	packageNameToPackageImport map[string]string
	currentType                string
	possibleTaggedTypes        map[string]map[string]Tag
}

func (f *InferredInfo) PackageName() string {
	return f.packageName
}

func (f *InferredInfo) PackageImportName() string {
	return f.pkgImportName
}

func (f *InferredInfo) RetrieveUnions() []*UnionLike {
	var result []*UnionLike
	for _, shape := range f.RetrieveShapes() {
		if unionShape, ok := shape.(*UnionLike); ok {
			result = append(result, unionShape)
		}
	}

	return result
}

func (f *InferredInfo) RetrieveStruct(name string) *StructLike {
	result, ok := f.shapes[name].(*StructLike)
	if !ok {
		return nil
	}

	return result
}

func (f *InferredInfo) RetrieveUnion(name string) *UnionLike {
	var variants []Shape
	for _, variant := range f.possibleVariantTypes[name] {
		variants = append(variants, f.shapes[variant])
	}

	if len(variants) == 0 {
		return nil
	}

	return &UnionLike{
		Name:          name,
		PkgName:       f.packageName,
		PkgImportName: f.pkgImportName,
		Variant:       variants,
		Tags:          f.possibleTaggedTypes[name],
	}
}

func (f *InferredInfo) RetrieveShapes() []Shape {
	shapes := make(map[string]Shape)
	for name, shape := range f.shapes {
		shapes[name] = shape
	}

	var result = make([]Shape, 0)
	for unionName, variantsNames := range f.possibleVariantTypes {
		union := f.RetrieveUnion(unionName)
		if union == nil {
			continue
		}

		result = append(result, union)

		delete(shapes, unionName)
		for _, variantName := range variantsNames {
			delete(shapes, variantName)
		}
	}

	for _, shape := range shapes {
		result = append(result, shape)
	}

	return result
}

func (f *InferredInfo) RetrieveStructs() []*StructLike {
	var result []*StructLike
	for _, shape := range f.RetrieveShapes() {
		if structShape, ok := shape.(*StructLike); ok {
			result = append(result, structShape)
		}
	}

	return result
}

func (f *InferredInfo) RetrieveShapeNamedAs(name string) Shape {
	return f.shapes[name]
}

func (f *InferredInfo) RetrieveShapesTaggedAs(tagName string) []Shape {
	var result []Shape
	for _, shape := range f.RetrieveShapes() {
		tags := Tags(shape)
		if _, ok := tags[tagName]; ok {
			result = append(result, shape)
		}
	}

	return result
}

func (f *InferredInfo) Visit(n ast.Node) ast.Visitor {
	opt := f.optionAST()
	switch t := n.(type) {
	case *ast.GenDecl:
		if t.Tok != token.TYPE {
			return f
		}

		// detect declaration of union type
		// either as comment
		// //go:generate mkunion -name=Example
		// //go:tag mkunion:"Example"
		tags := ExtractDocumentTags(t.Doc)

		// detect single declaration of type with a comment block
		// // some comment
		// type A struct {}
		if t.Lparen == 0 && t.Rparen == 0 && len(t.Specs) == 1 {
			switch s := t.Specs[0].(type) {
			case *ast.TypeSpec:
				// extract individual tags for each of variant
				f.possibleTaggedTypes[s.Name.Name] = tags
				return f
			}
		}

		// when there are more than one spec block,
		// it means that we are dealing with union (by convention)

		// register tags for specific type inside block:
		// type (
		//   ...
		// )
		for _, spec := range t.Specs {
			switch s := spec.(type) {
			case *ast.TypeSpec:
				// extract individual tags for each of variant
				f.possibleTaggedTypes[s.Name.Name] = ExtractDocumentTags(s.Doc)
			}
		}

		unionName := ""
		if unionTag, ok := tags["mkunion"]; ok {
			unionName = unionTag.Value
		} else {
			comment := shared.Comment(t.Doc)
			names := matchGoGenerateExtractUnionName.FindStringSubmatch(comment)
			if len(names) < 2 {
				return f
			}
			unionName = names[1]
		}

		// register union specific tags
		f.possibleTaggedTypes[unionName] = tags

		// start capturing possible variants
		if _, ok := f.possibleVariantTypes[unionName]; !ok {
			f.possibleVariantTypes[unionName] = make([]string, 0)
		}

		for _, spec := range t.Specs {
			switch s := spec.(type) {
			case *ast.TypeSpec:
				// register possible variant for union
				// NOTE: this is only convention that unions must be declared as type group specification:
				// type (
				// 	Variant2 struct {}
				//	Variant2 int
				//)
				f.possibleVariantTypes[unionName] = append(f.possibleVariantTypes[unionName], s.Name.Name)
			}
		}
		return f

	case *ast.FuncDecl:
		return nil

	case *ast.File:
		if t.Name != nil {
			f.packageName = t.Name.String()
		}

		f.packageNameToPackageImport = map[string]string{
			f.packageName: f.pkgImportName,
		}
		for _, imp := range t.Imports {
			if imp.Name != nil {
				f.packageNameToPackageImport[imp.Name.String()] = strings.Trim(imp.Path.Value, "\"")
			} else {
				f.packageNameToPackageImport[path.Base(strings.Trim(imp.Path.Value, "\""))] = strings.Trim(imp.Path.Value, "\"")
			}
		}

	case *ast.TypeSpec:
		f.currentType = t.Name.Name

		// Detect named literal types like:
		// type A string
		// type B int
		// type C bool
		switch next := t.Type.(type) {
		case *ast.Ident:
			switch next.Name {
			case "string":
				f.shapes[f.currentType] = &AliasLike{
					Name:          f.currentType,
					PkgName:       f.packageName,
					PkgImportName: f.pkgImportName,
					IsAlias:       IsAlias(t),
					Type:          &StringLike{},
					Tags:          f.possibleTaggedTypes[f.currentType],
				}

			case "int", "int8", "int16", "int32", "int64",
				"uint", "uint8", "uint16", "uint32", "uint64",
				"float64", "float32", "byte", "rune":
				f.shapes[f.currentType] = &AliasLike{
					Name:          f.currentType,
					PkgName:       f.packageName,
					PkgImportName: f.pkgImportName,
					IsAlias:       IsAlias(t),
					Type: &NumberLike{
						Kind: TypeStringToNumberKindMap[next.Name],
					},
					Tags: f.possibleTaggedTypes[f.currentType],
				}

			case "bool":
				f.shapes[f.currentType] = &AliasLike{
					Name:          f.currentType,
					PkgName:       f.packageName,
					PkgImportName: f.pkgImportName,
					IsAlias:       IsAlias(t),
					Type:          &BooleanLike{},
					Tags:          f.possibleTaggedTypes[f.currentType],
				}

			default:
				// alias type from the same package
				// example
				//  type A ListOf
				f.shapes[f.currentType] = &AliasLike{
					Name:          f.currentType,
					PkgName:       f.packageName,
					PkgImportName: f.pkgImportName,
					IsAlias:       IsAlias(t),
					Type: &RefName{
						Name:          next.Name,
						PkgName:       f.packageName,
						PkgImportName: f.pkgImportName,
					},
					Tags: f.possibleTaggedTypes[f.currentType],
				}
			}

		case *ast.SelectorExpr:
			// alias type from other packages
			// example:
			//  type A time.Time
			f.shapes[f.currentType] = &AliasLike{
				Name:          f.currentType,
				PkgName:       f.packageName,
				PkgImportName: f.pkgImportName,
				IsAlias:       IsAlias(t),
				Type:          f.selectExrToShape(next),
				Tags:          f.possibleTaggedTypes[f.currentType],
			}

		case *ast.IndexExpr:
			// alias of type that has one type params initialized
			// example:
			//  type A ListOf[any]
			//  type B ListOf[ListOf[any]]
			f.shapes[f.currentType] = &AliasLike{
				Name:          f.currentType,
				PkgName:       f.packageName,
				PkgImportName: f.pkgImportName,
				IsAlias:       IsAlias(t),
				Type:          FromAST(next, opt...),
				Tags:          f.possibleTaggedTypes[f.currentType],
			}

		case *ast.IndexListExpr:
			// alias of type that has two type params initialized
			// example:
			//  type A ListOf2[string, int]
			//  type B ListOf2[ListOf2[string, int], ListOf2[string, int]]
			f.shapes[f.currentType] = &AliasLike{
				Name:          f.currentType,
				PkgName:       f.packageName,
				PkgImportName: f.pkgImportName,
				IsAlias:       IsAlias(t),
				Type:          FromAST(next, opt...),
				Tags:          f.possibleTaggedTypes[f.currentType],
			}

		case *ast.MapType:
			// example:
			//  type A map[string]string
			//  type B map[string]ListOf2[string, int]
			f.shapes[f.currentType] = &AliasLike{
				Name:          f.currentType,
				PkgName:       f.packageName,
				PkgImportName: f.pkgImportName,
				IsAlias:       IsAlias(t),
				Type: &MapLike{
					Key:          FromAST(next.Key, opt...),
					Val:          FromAST(next.Value, opt...),
					KeyIsPointer: IsStarExpr(next.Key),
					ValIsPointer: IsStarExpr(next.Value),
				},
				Tags: f.possibleTaggedTypes[f.currentType],
			}

		case *ast.ArrayType:
			f.shapes[f.currentType] = &AliasLike{
				Name:          f.currentType,
				PkgName:       f.packageName,
				PkgImportName: f.pkgImportName,
				IsAlias:       IsAlias(t),
				Type: &ListLike{
					Element:          FromAST(next.Elt, opt...),
					ElementIsPointer: IsStarExpr(next.Elt),
					ArrayLen:         tryGetArrayLen(next.Len),
				},
				Tags: f.possibleTaggedTypes[f.currentType],
			}

		case *ast.StructType:
			f.shapes[f.currentType] = &StructLike{
				Name:          f.currentType,
				PkgName:       f.packageName,
				PkgImportName: f.pkgImportName,
				TypeParams:    f.extractTypeParams(t.TypeParams),
				Tags:          f.possibleTaggedTypes[f.currentType],
			}
		}

	case *ast.StructType:
		if !t.Struct.IsValid() {
			break
		}

		structShape, ok := f.shapes[f.currentType].(*StructLike)
		if !ok {
			log.Warnf("shape.InferFromFile: could not cast %s to StructLike", f.currentType)
			return f
		}

		for _, field := range t.Fields.List {
			// this happens when field is embedded in struct
			// something like `type A struct { B }`
			if len(field.Names) == 0 {
				switch typ := field.Type.(type) {
				case *ast.Ident:
					structShape.Fields = append(structShape.Fields, &FieldLike{
						Name: typ.Name,
						Type: FromAST(typ, opt...),
					})
					break
				default:
					log.Warnf("shape.InferFromFile: unknown ast type embedded in struct: %T\n", typ)
					continue
				}
			}

			for _, fieldName := range field.Names {
				if !fieldName.IsExported() {
					continue
				}

				var typ Shape
				switch ttt := field.Type.(type) {
				// selectors in struct, means that we are using type from other package
				case *ast.SelectorExpr:
					typ = f.selectExrToShape(ttt)
				// this is reference to other struct in the same package or other package
				case *ast.StarExpr:
					if selector, ok := ttt.X.(*ast.SelectorExpr); ok {
						typ = f.selectExrToShape(selector)
						switch x := typ.(type) {
						case *RefName:
							x.IsPointer = true
						}
					} else {
						typ = FromAST(ttt, opt...)
					}

				case *ast.IndexExpr, *ast.Ident, *ast.ArrayType, *ast.MapType, *ast.StructType:
					typ = FromAST(ttt, opt...)

				default:
					log.Warnf("shape.InferFromFile: unknown ast type in  %s.%s: %T\n", f.currentType, fieldName.Name, ttt)
					typ = &Any{}
				}

				typ = CleanTypeThatAreOvershadowByTypeParam(typ, structShape.TypeParams)

				tag := ""
				if field.Tag != nil {
					tag = field.Tag.Value
				}

				tags := ExtractTags(tag)
				desc := TagsToDesc(tags)
				guard := TagsToGuard(tags)

				structShape.Fields = append(structShape.Fields, &FieldLike{
					Name:      fieldName.Name,
					Type:      typ,
					Desc:      desc,
					Guard:     guard,
					Tags:      tags,
					IsPointer: IsStarExpr(field.Type),
				})
			}
		}

		f.shapes[f.currentType] = structShape
		log.Infof("shape.InferFromFile: struct %s: %s\n", f.currentType, ToStr(structShape))
	}

	return f
}

func CleanTypeThatAreOvershadowByTypeParam(typ Shape, params []TypeParam) Shape {
	return MustMatchShape(
		typ,
		func(x *Any) Shape {
			return x
		},
		func(x *RefName) Shape {
			if nameExistsInParams(x.Name, params) {
				x.PkgName = ""
				x.PkgImportName = ""
			}

			for i, name := range x.Indexed {
				x.Indexed[i] = CleanTypeThatAreOvershadowByTypeParam(name, params)
			}

			return x
		},
		func(x *AliasLike) Shape {
			if nameExistsInParams(x.Name, params) {
				x.PkgName = ""
				x.PkgImportName = ""
			}

			return x
		},
		func(x *BooleanLike) Shape {
			return x
		},
		func(x *StringLike) Shape {
			return x
		},
		func(x *NumberLike) Shape {
			return x

		},
		func(x *ListLike) Shape {
			x.Element = CleanTypeThatAreOvershadowByTypeParam(x.Element, params)
			return x
		},
		func(x *MapLike) Shape {
			x.Key = CleanTypeThatAreOvershadowByTypeParam(x.Key, params)
			x.Val = CleanTypeThatAreOvershadowByTypeParam(x.Val, params)
			return x
		},
		func(x *StructLike) Shape {
			if nameExistsInParams(x.Name, params) {
				x.PkgName = ""
				x.PkgImportName = ""
			}

			for _, field := range x.Fields {
				field.Type = CleanTypeThatAreOvershadowByTypeParam(field.Type, params)
			}
			return x
		},
		func(x *UnionLike) Shape {
			if nameExistsInParams(x.Name, params) {
				x.PkgName = ""
				x.PkgImportName = ""
			}

			for _, variant := range x.Variant {
				variant = CleanTypeThatAreOvershadowByTypeParam(variant, params)
			}
			return x
		},
	)
}

func IndexWith(y Shape, indexed []Shape) Shape {
	x, ok := y.(*StructLike)
	if !ok {
		return y
	}

	if len(x.TypeParams) != len(indexed) ||
		len(x.TypeParams) == 0 {
		return x
	}

	params := map[string]*RefName{}
	for i, param := range x.TypeParams {
		typ := indexed[i]

		ref, ok := typ.(*RefName)
		if !ok {
			panic(fmt.Errorf("IndexWith: expected indexed type to be *RefName, got %T", typ))
		}

		params[param.Name] = ref
	}

	return InstantiateTypeThatAreOvershadowByTypeParam(x, params)
}

func InstantiateTypeThatAreOvershadowByTypeParam(typ Shape, replacement map[string]*RefName) Shape {
	return MustMatchShape(
		typ,
		func(x *Any) Shape {
			return x
		},
		func(x *RefName) Shape {
			if result, err := replacement[x.Name]; err {
				return result
			}

			result := &RefName{
				Name:          x.Name,
				PkgName:       x.PkgName,
				PkgImportName: x.PkgImportName,
				IsPointer:     x.IsPointer,
			}
			for _, name := range x.Indexed {
				result.Indexed = append(result.Indexed, InstantiateTypeThatAreOvershadowByTypeParam(name, replacement))
			}

			return result
		},
		func(x *AliasLike) Shape {
			result := &AliasLike{
				Name:          x.Name,
				PkgName:       x.PkgName,
				PkgImportName: x.PkgImportName,
				IsAlias:       x.IsAlias,
				Type:          InstantiateTypeThatAreOvershadowByTypeParam(x.Type, replacement),
				Tags:          x.Tags,
			}
			return result
		},
		func(x *BooleanLike) Shape {
			return x
		},
		func(x *StringLike) Shape {
			return x
		},
		func(x *NumberLike) Shape {
			return x
		},
		func(x *ListLike) Shape {
			result := &ListLike{
				Element:          InstantiateTypeThatAreOvershadowByTypeParam(x.Element, replacement),
				ElementIsPointer: x.ElementIsPointer,
			}
			return result
		},
		func(x *MapLike) Shape {
			result := &MapLike{
				Key:          InstantiateTypeThatAreOvershadowByTypeParam(x.Key, replacement),
				KeyIsPointer: x.KeyIsPointer,
				Val:          InstantiateTypeThatAreOvershadowByTypeParam(x.Val, replacement),
				ValIsPointer: x.ValIsPointer,
			}
			return result
		},
		func(x *StructLike) Shape {
			result := &StructLike{
				Name:          x.Name,
				PkgName:       x.PkgName,
				PkgImportName: x.PkgImportName,
				TypeParams:    x.TypeParams,
				Tags:          x.Tags,
			}

			for _, field := range x.Fields {
				result.Fields = append(result.Fields, &FieldLike{
					Name:      field.Name,
					Type:      InstantiateTypeThatAreOvershadowByTypeParam(field.Type, replacement),
					Desc:      field.Desc,
					Guard:     field.Guard,
					IsPointer: field.IsPointer,
					Tags:      field.Tags,
				})
			}
			return result
		},
		func(x *UnionLike) Shape {
			result := &UnionLike{
				Name:          x.Name,
				PkgName:       x.PkgName,
				PkgImportName: x.PkgImportName,
				Tags:          x.Tags,
			}
			for _, variant := range x.Variant {
				result.Variant = append(result.Variant, InstantiateTypeThatAreOvershadowByTypeParam(variant, replacement))
			}
			return result
		},
	)
}

func nameExistsInParams(name string, params []TypeParam) bool {
	for _, param := range params {
		if param.Name == name {
			return true
		}
	}

	return false
}

func (f *InferredInfo) optionAST() []FromASTOption {
	return []FromASTOption{
		InjectPkgName(f.packageName),
		InjectPkgImportName(f.packageNameToPackageImport),
	}
}

func IsAlias(t *ast.TypeSpec) bool {
	return t.Assign != 0
}

func (f *InferredInfo) selectExrToShape(ttt *ast.SelectorExpr) Shape {
	if ident, ok := ttt.X.(*ast.Ident); ok {
		pkgName := ident.Name
		return FromAST(ttt.Sel, InjectPkgName(pkgName), InjectPkgImportName(f.packageNameToPackageImport))
	}

	return FromAST(ttt, f.optionAST()...)
}

func (f *InferredInfo) extractTypeParams(params *ast.FieldList) []TypeParam {
	if params == nil {
		return nil
	}

	var result []TypeParam
	for _, param := range params.List {
		typ := FromAST(param.Type,
			InjectPkgImportName(f.packageNameToPackageImport),
			InjectPkgName(f.packageName),
		)

		if len(param.Names) == 0 {
			result = append(result, TypeParam{
				Name: param.Type.(*ast.Ident).Name,
				Type: typ,
			})
			continue
		}

		for _, name := range param.Names {
			result = append(result, TypeParam{
				Name: name.Name,
				Type: typ,
			})
		}
	}

	return result
}

func tryGetArrayLen(expr ast.Expr) *int {
	if expr == nil {
		return nil
	}

	switch x := expr.(type) {
	case *ast.BasicLit:
		if x.Kind == token.INT {
			n, _ := strconv.Atoi(x.Value)
			return &n
		}
	}

	return nil
}
