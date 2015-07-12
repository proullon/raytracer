package back

import (
	"container/list"
	"fmt"
	// "log"
)

type Coord struct {
	x float64
	y float64
	z float64
}

type Pixel struct {
	X int
	Y int
	Color
}

func Trace(channel chan Pixel, width int, height int, objects *list.List, eye *Eye, lights *list.List) {
	traceParallel(channel, width, height, objects, eye, lights)
}

func fewfwfe(w int, h int, width int, height int, objects *list.List, eye *Eye, lights *list.List, res chan Pixel) {
	color := rgb(Ray(w, h, width, height, objects, eye, lights))

	pixel := Pixel{X:w, Y:h}
	pixel.R = color.R
	pixel.G = color.G
	pixel.B = color.B
	pixel.A = color.A

	// log.Printf("Sending pixel to websocket : %v\n", pixel)
	res <- pixel
}

func traceParallel(res chan Pixel, width int, height int, objects *list.List, eye *Eye, lights *list.List) (img [][]int) {
	// var color int

	img = make([][]int, height)
	for i := range img {
		img[i] = make([]int, width)
	}

	for w := 0; w < width; w++ {
		for h := 0; h < height; h++ {
			go fewfwfe(w, h, width, height, objects, eye, lights, res)
		}
	}

	return
}

func traceSequential(width int, height int, objects *list.List, eye *Eye, lights *list.List) (img [][]Color) {
	var color int

	img = make([][]Color, height)
	for i := range img {
		img[i] = make([]Color, width)
	}

	for w := 0; w < width; w++ {
		for h := 0; h < height; h++ {
			color = Ray(w, h, width, height, objects, eye, lights)
			img[h][w] = rgb(color)
			if color != 0 {
				fmt.Printf("x:%d y:%d = %d\n", h, w, color)
			}
		}
	}

	return
}