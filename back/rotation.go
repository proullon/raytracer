package back

import (
	"math"
)

func radian(x float64) float64 {
	return math.Pi * x / 180
}

func rotate(baseline *Coord, rx float64, ry float64, rz float64) {
	rotate_x(baseline, radian(rx));
 	rotate_y(baseline, radian(ry));
 	rotate_z(baseline, radian(rz));
}

func rotate_x(baseline *Coord, a float64) {
	var tmp Coord

	tmp = *baseline
	sina := math.Sin(a)
	cosa := math.Cos(a)

	// baseline.x unchanged
	baseline.y = cosa * tmp.y - sina * tmp.z
	baseline.z = sina *tmp.y + cosa * tmp.z
}

func rotate_y(baseline *Coord, a float64) {
	var tmp Coord

	tmp = *baseline
	sina := math.Sin(a)
	cosa := math.Cos(a)

	baseline.x = cosa * tmp.x + sina * tmp.z
	// baseline.y unchanged
	baseline.z = -sina * tmp.x + cosa * tmp.z
}

func rotate_z(baseline *Coord, a float64) {
	var tmp Coord

	tmp = *baseline
	sina := math.Sin(a)
	cosa := math.Cos(a)

	baseline.x = cosa * tmp.x - sina * tmp.y
	baseline.y = sina * tmp.x + cosa * tmp.y
	// baseline.z unchanged
}
