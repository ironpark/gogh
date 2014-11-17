package gogh

import (
	"github.com/ironpark/gogh/shape"
)

func (src *Img) Draw(s shape.Shape, x, y int, color interface{}) {
	switch v := s.(type) {
	case shape.Circle:
		drawCircle(src, v, x, y, color)
	case shape.Plus:
	case shape.Rect:
	}
}

func drawCircle(img *Img, circle shape.Circle, x, y int, color interface{}) {
	radius := circle.Radius

	x0, y0 := x, y
	//f := 1 - radius
	//ddF_x, ddF_y := 1, -2*radius
	x, y = radius, 0

	radiusError := 1 - x

	if gray, e := color.(int); e {
		for x >= y {
			img.At(x0+x, y0+y).Set(gray, gray, gray, 255)
			img.At(x0-x, y0+y).Set(gray, gray, gray, 255)
			img.At(x0+x, y0-y).Set(gray, gray, gray, 255)
			img.At(x0-x, y0-y).Set(gray, gray, gray, 255)
			img.At(x0+y, y0+x).Set(gray, gray, gray, 255)
			img.At(x0-y, y0+x).Set(gray, gray, gray, 255)
			img.At(x0+y, y0-x).Set(gray, gray, gray, 255)
			img.At(x0-y, y0-x).Set(gray, gray, gray, 255)
			y++
			if radiusError < 0 {
				radiusError += 2*y + 1
			} else {
				x--
				radiusError += 2 * (y - x + 1)
			}
		}
	} else {
		if c, e := color.([]int); e {
			switch len(c) {
			case 1:
				for x >= y {
					img.At(x0+x, y0+y).Set(c[0], c[0], c[0], 255)
					img.At(x0-x, y0+y).Set(c[0], c[0], c[0], 255)
					img.At(x0+x, y0-y).Set(c[0], c[0], c[0], 255)
					img.At(x0-x, y0-y).Set(c[0], c[0], c[0], 255)
					img.At(x0+y, y0+x).Set(c[0], c[0], c[0], 255)
					img.At(x0-y, y0+x).Set(c[0], c[0], c[0], 255)
					img.At(x0+y, y0-x).Set(c[0], c[0], c[0], 255)
					img.At(x0-y, y0-x).Set(c[0], c[0], c[0], 255)
					y++
					if radiusError < 0 {
						radiusError += 2*y + 1
					} else {
						x--
						radiusError += 2 * (y - x + 1)
					}
				}
			case 2:
				for x >= y {
					img.At(x0+x, y0+y).Set(c[0], c[1], 0, 255)
					img.At(x0-x, y0+y).Set(c[0], c[1], 0, 255)
					img.At(x0+x, y0-y).Set(c[0], c[1], 0, 255)
					img.At(x0-x, y0-y).Set(c[0], c[1], 0, 255)
					img.At(x0+y, y0+x).Set(c[0], c[1], 0, 255)
					img.At(x0-y, y0+x).Set(c[0], c[1], 0, 255)
					img.At(x0+y, y0-x).Set(c[0], c[1], 0, 255)
					img.At(x0-y, y0-x).Set(c[0], c[1], 0, 255)
					y++
					if radiusError < 0 {
						radiusError += 2*y + 1
					} else {
						x--
						radiusError += 2 * (y - x + 1)
					}
				}
			case 3:
				for x >= y {
					img.At(x0+x, y0+y).Set(c[0], c[1], c[2], 255)
					img.At(x0-x, y0+y).Set(c[0], c[1], c[2], 255)
					img.At(x0+x, y0-y).Set(c[0], c[1], c[2], 255)
					img.At(x0-x, y0-y).Set(c[0], c[1], c[2], 255)
					img.At(x0+y, y0+x).Set(c[0], c[1], c[2], 255)
					img.At(x0-y, y0+x).Set(c[0], c[1], c[2], 255)
					img.At(x0+y, y0-x).Set(c[0], c[1], c[2], 255)
					img.At(x0-y, y0-x).Set(c[0], c[1], c[2], 255)
					y++
					if radiusError < 0 {
						radiusError += 2*y + 1
					} else {
						x--
						radiusError += 2 * (y - x + 1)
					}
				}

			case 4:
				for x >= y {
					img.At(x0+x, y0+y).Set(c[0], c[1], c[2], c[3])
					img.At(x0-x, y0+y).Set(c[0], c[1], c[2], c[3])
					img.At(x0+x, y0-y).Set(c[0], c[1], c[2], c[3])
					img.At(x0-x, y0-y).Set(c[0], c[1], c[2], c[3])
					img.At(x0+y, y0+x).Set(c[0], c[1], c[2], c[3])
					img.At(x0-y, y0+x).Set(c[0], c[1], c[2], c[3])
					img.At(x0+y, y0-x).Set(c[0], c[1], c[2], c[3])
					img.At(x0-y, y0-x).Set(c[0], c[1], c[2], c[3])
					y++
					if radiusError < 0 {
						radiusError += 2*y + 1
					} else {
						x--
						radiusError += 2 * (y - x + 1)
					}
				}
			default:
			}
		}
	}

}
