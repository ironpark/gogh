package gogh

import "github.com/ironpark/gogh/shape"

func (src *Img) Draw(s shape.Shape, x, y int) {
	switch v := s.(type) {
	case shape.Circle:
		drawCircle(src, v, x, y)
	case shape.Plus:
	case shape.Rect:
	}
}

func drawCircle(img *Img, circle shape.Circle, x, y int) {
	radius := circle.Radius

	x0, y0 := x, y
	//f := 1 - radius
	//ddF_x, ddF_y := 1, -2*radius
	x, y = radius, 0

	radiusError := 1 - x

	for x >= y {
		img.At(x0+x, y0+y).Set(0, 0, 0, 255)
		img.At(x0-x, y0+y).Set(0, 0, 0, 255)
		img.At(x0+x, y0-y).Set(0, 0, 0, 255)
		img.At(x0-x, y0-y).Set(0, 0, 0, 255)
		img.At(x0+y, y0+x).Set(0, 0, 0, 255)
		img.At(x0-y, y0+x).Set(0, 0, 0, 255)
		img.At(x0+y, y0-x).Set(0, 0, 0, 255)
		img.At(x0-y, y0-x).Set(0, 0, 0, 255)
		y++
		if radiusError < 0 {
			radiusError += 2*y + 1
		} else {
			x--
			radiusError += 2 * (y - x + 1)
		}
	}
}
