package mkunion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneration(t *testing.T) {
	g := NewVisitorGenerator("Vehicle", []string{"Plane", "Car", "Boat"}, NewHelper(WithPackageName("visitor")))

	result, err := g.Generate()
	assert.NoError(t, err)
	assert.Equal(t, `// Code generated by mkunion. DO NOT EDIT.
package visitor

import "github.com/widmogrod/mkunion/f"


type VehicleVisitor interface {
	VisitPlane(v *Plane) any
	VisitCar(v *Car) any
	VisitBoat(v *Boat) any
}

type Vehicle interface {
	AcceptVehicle(g VehicleVisitor) any
}

func (r *Plane) AcceptVehicle(v VehicleVisitor) any { return v.VisitPlane(r) }
func (r *Car) AcceptVehicle(v VehicleVisitor) any { return v.VisitCar(r) }
func (r *Boat) AcceptVehicle(v VehicleVisitor) any { return v.VisitBoat(r) }

var (
	_ Vehicle = (*Plane)(nil)
	_ Vehicle = (*Car)(nil)
	_ Vehicle = (*Boat)(nil)
)

func MatchVehicle[TOut any](
	x Vehicle,
	f1 func(x *Plane) TOut,
	f2 func(x *Car) TOut,
	f3 func(x *Boat) TOut,
	df func(x Vehicle) TOut,
) TOut {
	return f.Match3(x, f1, f2, f3, df)
}

func MatchVehicleR2[TOut1, TOut2 any](
	x Vehicle,
	f1 func(x *Plane) (TOut1, TOut2),
	f2 func(x *Car) (TOut1, TOut2),
	f3 func(x *Boat) (TOut1, TOut2),
	df func(x Vehicle) (TOut1, TOut2),
) (TOut1, TOut2) {
	return f.Match3R2(x, f1, f2, f3, df)
}

func MustMatchVehicle[TOut any](
	x Vehicle,
	f1 func(x *Plane) TOut,
	f2 func(x *Car) TOut,
	f3 func(x *Boat) TOut,
) TOut {
	return f.MustMatch3(x, f1, f2, f3)
}

func MustMatchVehicleR0(
	x Vehicle,
	f1 func(x *Plane),
	f2 func(x *Car),
	f3 func(x *Boat),
) {
	f.MustMatch3R0(x, f1, f2, f3)
}

func MustMatchVehicleR2[TOut1, TOut2 any](
	x Vehicle,
	f1 func(x *Plane) (TOut1, TOut2),
	f2 func(x *Car) (TOut1, TOut2),
	f3 func(x *Boat) (TOut1, TOut2),
) (TOut1, TOut2) {
	return f.MustMatch3R2(x, f1, f2, f3)
}`, string(result))
}
