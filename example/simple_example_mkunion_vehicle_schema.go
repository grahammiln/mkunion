// Code generated by mkunion. DO NOT EDIT.
package example

import (
	"github.com/widmogrod/mkunion/x/schema"
)

func init() {
	schema.RegisterTransformations(VehicleSchemaTransformations())
	schema.RegisterRules(VehicleSchemaRules())
}

func VehicleSchemaTransformations() []schema.TransformFunc {
	return []schema.TransformFunc{
		schema.WrapStruct(&Car{}, "Car"),
		schema.WrapStruct(&Plane{}, "Plane"),
		schema.WrapStruct(&Boat{}, "Boat"),
	}
}

func VehicleSchemaRules() []schema.RuleMatcher {
	return []schema.RuleMatcher{
		schema.UnwrapStruct(&Car{}, "Car"),
		schema.UnwrapStruct(&Plane{}, "Plane"),
		schema.UnwrapStruct(&Boat{}, "Boat"),
	}
}
