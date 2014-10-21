package gogh

import (
	"image"
	//_ "image/bmp"
	_ "image/jpeg"
	_ "image/png"
)

func (src *Img) Grayscale() *Img {
	gray := src.Clone()
	bounds := img.Bounds()
	for i := 0; i < bounds.Max.X; i++ {
		for j := 0; j < bounds.Max.Y; j++ {
			r, g, b := gray.At(x, y).RGBA()
			color := (r + g + b) / 3
			gray.At(x, y).Set(color, color, color)
		}
	}
	return gray
}
func (src *Img) Binarization(T int, reverse bool) *Img {
	dst := src.Clone()
	bounds := src.Bounds()
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			pixel := dst.At(x, y)
			if reverse {
				if T < pixel.Gray() {
					pixel.Set(0, 0, 0)
				} else {
					pixel.Set(255, 255, 255)
				}
			} else {
				if T > pixel.Gray() {
					pixel.Set(0, 0, 0)
				} else {
					pixel.Set(255, 255, 255)
				}
			}
		}
	}
	return dst
}
