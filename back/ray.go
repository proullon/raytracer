package back

import (
	"container/list"
)

func choose_color(objects *list.List) (color int) {
	var lower, k float64
	var obj Object

	lower = 100000
	for e := objects.Front(); e != nil; e = e.Next() {
		obj = e.Value.(Object)

		k = obj.K()
		if k > 0 && k < lower {
			lower = k
			color = obj.Color()
		}
	}

	return
}

func findIntersection(baseline Coord, eye *Eye, obj Object) {
	rotate(&baseline, obj.RX(), obj.RY(), obj.RZ())

	baseline.x -= obj.X();
    baseline.y -= obj.Y();
    baseline.z -= obj.Z();

    obj.Intersection(baseline, eye)
}

func Ray(x int, y int, width int, heigth int, objects *list.List, eye *Eye, lights *list.List) (color int) {

	baseline := Coord{x:100.0, y:float64((width / 2) - x), z:float64((heigth / 2) - y)} 

	for e := objects.Front(); e != nil; e = e.Next() {
		// do something with e.Value
		findIntersection(baseline, eye, e.Value.(Object))
	}

	return choose_color(objects)
}