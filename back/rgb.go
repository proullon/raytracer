package back

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func rgb(color int) (c Color) {
	c.R = uint8(color)
	c.G = uint8(color >> 8)
	c.B = uint8(color >> 16)
	c.A = uint8(color >> 24)
	return
}