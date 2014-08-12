package back

type Object interface {
	Intersection(baseline Coord, eye *Eye)
	K() float64
	Color() int

	X() float64
	Y() float64
	Z() float64

	RX() float64
	RY() float64
	RZ() float64
}

type BaseObject struct {
	Coord
	rx float64
	ry float64
	rz float64
	k float64
	color int

  // int		id;
  // int		cte;
}

func (bo *BaseObject) RX() float64 {
	return bo.rx
}

func (bo *BaseObject) RY() float64 {
	return bo.ry
}

func (bo *BaseObject) RZ() float64 {
	return bo.rz
}

func (bo *BaseObject) X() float64 {
	return bo.x
}

func (bo *BaseObject) Y() float64 {
	return bo.y
}

func (bo *BaseObject) Z() float64 {
	return bo.z
}

func (bo *BaseObject) K() float64 {
	return bo.k
}

func (bo *BaseObject) Color() int {
	return bo.color
}
