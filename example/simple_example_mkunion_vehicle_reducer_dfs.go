// Code generated by mkunion. DO NOT EDIT.
package example

type (
	VehicleReducer[A any] interface {
		ReduceCar(x *Car, agg A) (result A, stop bool)
		ReducePlane(x *Plane, agg A) (result A, stop bool)
		ReduceBoat(x *Boat, agg A) (result A, stop bool)
	}
)

type VehicleDepthFirstVisitor[A any] struct {
	stop   bool
	result A
	reduce VehicleReducer[A]
}

var _ VehicleVisitor = (*VehicleDepthFirstVisitor[any])(nil)

func (d *VehicleDepthFirstVisitor[A]) VisitCar(v *Car) any {
	d.result, d.stop = d.reduce.ReduceCar(v, d.result)
	if d.stop {
		return nil
	}

	return nil
}

func (d *VehicleDepthFirstVisitor[A]) VisitPlane(v *Plane) any {
	d.result, d.stop = d.reduce.ReducePlane(v, d.result)
	if d.stop {
		return nil
	}

	return nil
}

func (d *VehicleDepthFirstVisitor[A]) VisitBoat(v *Boat) any {
	d.result, d.stop = d.reduce.ReduceBoat(v, d.result)
	if d.stop {
		return nil
	}

	return nil
}

func ReduceVehicleDepthFirst[A any](r VehicleReducer[A], v Vehicle, init A) A {
	reducer := &VehicleDepthFirstVisitor[A]{
		result: init,
		reduce: r,
	}

	_ = v.Accept(reducer)

	return reducer.result
}
