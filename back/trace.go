package back

import (
	"fmt"
	"container/list"
)

type Coord struct {
	x float64
	y float64
	z float64
}

type Light struct {
	Coord
	color int
}

func Trace(width int, height int, objects *list.List, eye *Eye, lights *list.List) (img [][]int) {
	var color int

	img = make([][]int, height)
	for i := range img {
		img[i] = make([]int, width)
	}

	for w := 0; w < width; w++ {
		for h := 0; h < height; h++ {
			color = Ray(w, h, width, height, objects, eye, lights)
			img[h][w] = color
			if color != 0 {
			fmt.Printf("%d:%d = %d\n", w, h, color)				
			}
		}
	}

	return
}