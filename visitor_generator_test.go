package mkunion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneration(t *testing.T) {
	g := VisitorGenerator{
		Header:      Header,
		Name:        "Vehicle",
		Types:       []string{"Plane", "Car", "Boat"},
		PackageName: "visitor",
	}

	result, err := g.Generate()
	assert.NoError(t, err)
	assert.Equal(t, `// Code generated by mkunion. DO NOT EDIT.
package visitor

import (
	"github.com/widmogrod/mkunion/f"
)


type VehicleVisitor interface {
	VisitPlane(v *Plane) any
	VisitCar(v *Car) any
	VisitBoat(v *Boat) any
}

type Vehicle interface {
	Accept(g VehicleVisitor) any
}

func (r *Plane) Accept(v VehicleVisitor) any { return v.VisitPlane(r) }
func (r *Car) Accept(v VehicleVisitor) any { return v.VisitCar(r) }
func (r *Boat) Accept(v VehicleVisitor) any { return v.VisitBoat(r) }

var (
	_ Vehicle = (*Plane)(nil)
	_ Vehicle = (*Car)(nil)
	_ Vehicle = (*Boat)(nil)
)

type VehicleOneOf struct {
	Plane *Plane `+"`json:\",omitempty\"`"+`
	Car *Car `+"`json:\",omitempty\"`"+`
	Boat *Boat `+"`json:\",omitempty\"`"+`
}

func (r *VehicleOneOf) Accept(v VehicleVisitor) any {
	switch {
	case r.Plane != nil:
		return v.VisitPlane(r.Plane)
	case r.Car != nil:
		return v.VisitCar(r.Car)
	case r.Boat != nil:
		return v.VisitBoat(r.Boat)
	default:
		panic("unexpected")
	}
}

var _ Vehicle = (*VehicleOneOf)(nil)

type mapVehicleToOneOf struct{}

func (t *mapVehicleToOneOf) VisitPlane(v *Plane) any { return &VehicleOneOf{Plane: v} }
func (t *mapVehicleToOneOf) VisitCar(v *Car) any { return &VehicleOneOf{Car: v} }
func (t *mapVehicleToOneOf) VisitBoat(v *Boat) any { return &VehicleOneOf{Boat: v} }

var defaultMapVehicleToOneOf VehicleVisitor = &mapVehicleToOneOf{}

func MapVehicleToOneOf(v Vehicle) *VehicleOneOf {
	return v.Accept(defaultMapVehicleToOneOf).(*VehicleOneOf)
}

func MustMatchVehicle[TOut any](
	x Vehicle,
	f1 func(x *Plane) TOut,
	f2 func(x *Car) TOut,
	f3 func(x *Boat) TOut,
) TOut {
	return f.MustMatch3(x, f1, f2, f3)
}`, string(result))
}
