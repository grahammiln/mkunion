package schema

import (
	"encoding/json"
	"fmt"
	"github.com/widmogrod/mkunion/x/shared"
	"reflect"
)

var none = &None{}

func MkNone() *None {
	return none
}

func MkBool(b bool) *Bool {
	return (*Bool)(&b)
}

func MkInt(x int) *Number {
	v := float64(x)
	return (*Number)(&v)
}

func MkFloat(x float64) *Number {
	return (*Number)(&x)
}

func MkBinary(b []byte) *Binary {
	return &Binary{B: b}
}

func MkString(s string) *String {
	return (*String)(&s)
}

func MkList(items ...Schema) *List {
	result := make(List, len(items))
	copy(result, items)
	return &result
}
func MkMap(fields ...Field) *Map {
	var result = make(Map)
	for _, field := range fields {
		result[field.Name] = field.Value
	}
	return &result
}

func MkField(name string, value Schema) Field {
	return Field{
		Name:  name,
		Value: value,
	}
}

type (
	TypeListDefinition interface {
		NewListBuilder() ListBuilder
	}
	TypeMapDefinition interface {
		NewMapBuilder() MapBuilder
	}
)

type (
	ListBuilder interface {
		Append(value any) error
		Build() any
	}

	MapBuilder interface {
		Set(key string, value any) error
		Build() any
	}

	// mapBuilderCanProcessRawMapSchema returns marks special class of MapBuilder that they can work with raw Schema value,
	// and don't need go value that was decoded using default schemaToGo.
	// in technical terms, it disables recursive call to schemaToGo
	mapBuilderCanProcessRawMapSchema interface {
		BuildFromMapSchema(x *Map) (any, error)
	}
)

//go:generate go run ../../cmd/mkunion/main.go -name=Schema -skip-extension=schema,shape
type (
	None   struct{}
	Bool   bool
	Number float64
	String string
	Binary struct{ B []byte }
	List   []Schema
	Map    map[string]Schema
)

var _ json.Unmarshaler = (*Map)(nil)

func (x *Map) UnmarshalJSON(bytes []byte) error {
	*x = make(Map)
	return shared.JSONParseObject(bytes, func(key string, value []byte) error {
		val, err := SchemaFromJSON(value)
		if err != nil {
			return fmt.Errorf("schema.Map.UnmarshalJSON: %w", err)
		}

		(*x)[key] = val
		return nil
	})
}

type (
	Marshaler interface {
		MarshalSchema() (*Map, error)
	}

	Unmarshaler interface {
		UnmarshalSchema(x *Map) error
	}
)

type Field struct {
	Name  string
	Value Schema
}

func (f *Field) MarshalJSON() ([]byte, error) {
	field_Name, err := json.Marshal(f.Name)
	if err != nil {
		return nil, fmt.Errorf("schema.Field.MarshalJSON: Name; %w", err)
	}

	field_Value, err := SchemaToJSON(f.Value)
	if err != nil {
		return nil, fmt.Errorf("schema.Field.MarshalJSON: Value; %w", err)
	}

	return json.Marshal(map[string]json.RawMessage{
		"Name":  field_Name,
		"Value": field_Value,
	})
}

func (f *Field) UnmarshalJSON(bytes []byte) error {
	return shared.JSONParseObject(bytes, func(key string, value []byte) error {
		switch key {
		case "Name":
			return json.Unmarshal(value, &f.Name)
		case "Value":
			field, err := SchemaFromJSON(value)
			if err != nil {
				return fmt.Errorf("schema.Field.UnmarshalJSON: Value; %w", err)
			}
			f.Value = field
		}
		return nil
	})
}

var _ json.Unmarshaler = (*Field)(nil)
var _ json.Marshaler = (*Field)(nil)

type UnionInformationRule interface {
	UnionType() reflect.Type
	VariantsTypes() []reflect.Type
	IsUnionOrUnionType(t reflect.Type) bool
}

func UseStruct(t any) TypeMapDefinition {
	// Optimisation: When struct has its own definition how to populate it from schema,
	// we can use it instead of costly StructDefinition, that is based on reflection.
	if from, ok := t.(Unmarshaler); ok {
		// here is assumption that t is pointer to struct
		tType := reflect.ValueOf(from).Type().Elem()
		return UseSelfUnmarshallingStruct(func() Unmarshaler {
			// that's why here we create new emtpy type using reflection
			return reflect.New(tType).Interface().(Unmarshaler)
		})
	}

	return UseReflectionUnmarshallingStruct(t)
}

func UseTypeDef(definition TypeMapDefinition) TypeMapDefinition {
	return definition
}

func WhenPath(path []string, setter TypeMapDefinition) *WhenField[struct{}] {
	return &WhenField[struct{}]{
		path:       path,
		typeMapDef: setter,
	}
}

type RuleMatcher interface {
	MapDefFor(x *Map, path []string, config *goConfig) (TypeMapDefinition, bool)
	SchemaToUnionType(x any, schema Schema, config *goConfig) (Schema, bool)
}
