// gogh project gogh.go
package gogh

import (
	//"fmt"
	"image"
	"image/color"
)

var (
	GrayArray = make([]color.Color, 256)
)

func NewMat(src image.Image) *Mat {
	dst := ImageToRGBA(src)
	return &Mat{dst}
}

type Mat struct {
	src *image.RGBA
}

type Pixel struct {
	src   *image.RGBA
	X     int
	Y     int
	color color.Color
}

func (src *Mat) At(x, y int) *Pixel {
	return &Pixel{src.src, x, y, src.src.At(x, y)}
}

func (src *Pixel) RGBA() (int, int, int, int) {
	r, g, b, a := src.color.RGBA()
	return int(r >> 8), int(g >> 8), int(b >> 8), int(a >> 8)
}

func (src *Mat) Save(path string) {
	Save(path, src.src)
}

func (src *Mat) Clone() *Mat {
	return clone(src)
}

func (src *Pixel) Gray() int {
	gray, _, _, _ := src.color.RGBA()
	return int(gray >> 8)
}

func (src *Pixel) Set(r, g, b int) {
	src.src.Set(src.X, src.Y, color.Color(color.RGBA{uint8(r), uint8(g), uint8(b), uint8(255)}))
}

func (src *Mat) Bounds() image.Rectangle {
	return src.src.Bounds()
}

func (src *Mat) Pixels() []uint8 {
	return src.src.Pix
}
