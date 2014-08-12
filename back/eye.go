package back

type Eye struct {
	Coord
}

func NewEye(x float64, y float64, z float64) *Eye {
	e := &Eye{}

	e.x = x
	e.y = y
	e.z = z

	return e
}