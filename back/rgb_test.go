package back

import (
	"testing"
)

// # Example of colors :
// # -BLUE		255
// # -RED		16711680
// # -WHITE	11145489
// # -BLACK	0
// # -GREEN	13056
// # -CYAN		3667332

func TestRGB(t *testing.T) {
	var c Color


	c = rgb(13056)
	if c.r != 0 || c.g != 255 || c.b != 0 || c.a != 0 {
		t.Errorf("Green (13056) failed r:%v g:%v b:%v a:%v\n", c.r, c.g, c.b, c.a)
	}

	c = rgb(11145489)
	if c.r != 255 || c.g != 0 || c.b != 0 || c.a != 0 {
		t.Errorf("White (11145489) failed r:%v g:%v b:%v a:%v\n", c.r, c.g, c.b, c.a)
	}

	c = rgb(16711680)
	if c.r != 255 || c.g != 0 || c.b != 0 || c.a != 0 {
		t.Errorf("Red (16711680) failed r:%v g:%v b:%v a:%v\n", c.r, c.g, c.b, c.a)
	}

	c = rgb(255)
	if c.r != 0 || c.g != 0 || c.b != 255 || c.a != 0 {
		t.Errorf("Blue (255) failed r:%v g:%v b:%v a:%v\n", c.r, c.g, c.b, c.a)
	}
}