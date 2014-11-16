// gogh project gogh.go
package gogh

import (
	"image"
)

const (
	GRAY    = 0
	GRAY16  = 1
	NRGBA   = 3
	NRGBA64 = 4
	RGBA    = 5
	RGBA64  = 6
)

func NewImg(rect image.Rectangle, T int) *Img {
	var im []uint8
	switch T {
	case GRAY:
		im = image.NewGray(rect).Pix
	case GRAY16:
		im = image.NewGray16(rect).Pix
	case NRGBA:
		im = image.NewNRGBA(rect).Pix
	case NRGBA64:
		im = image.NewNRGBA64(rect).Pix
	case RGBA64:
		im = image.NewRGBA(rect).Pix
	case RGBA:
		im = image.NewRGBA64(rect).Pix
	default:
		return nil
	}

	return &Img{im, T, rect.Max.X, rect.Max.Y, rect}
}

type Img struct {
	Pixels    []uint8
	ImageType int
	Width     int
	Height    int
	Bounds    image.Rectangle
}

type Pixel struct {
	src *uint8
	X   int
	Y   int
}

func (img *Img) At(x, y int) *Pixel {
	return &Pixel{&img.Pixels[img.Width*y+x], x, y}
}

//TEMP!!! YOU MUST FIX!!
func (src *Pixel) RGBA() (int, int, int, int) {
	r := src.src
	g := src.src
	b := src.src
	a := src.src
	return int(*r), int(*g), int(*b), int(*a)
}

func (src *Img) Save(path string) {
	var im image.Image
	switch src.ImageType {
	case GRAY:
		t := image.NewGray(src.Bounds)
		t.Pix = src.Pixels
		im = t
	case GRAY16:
		t := image.NewGray16(src.Bounds)
		t.Pix = src.Pixels
		im = t
	case NRGBA:
		t := image.NewNRGBA(src.Bounds)
		t.Pix = src.Pixels
		im = t
	case NRGBA64:
		t := image.NewNRGBA64(src.Bounds)
		t.Pix = src.Pixels
		im = t
	case RGBA64:
		t := image.NewRGBA(src.Bounds)
		t.Pix = src.Pixels
		im = t
	case RGBA:
		t := image.NewRGBA64(src.Bounds)
		t.Pix = src.Pixels
		im = t
	}

	save(path, im)
}

func (src *Img) Clone() *Img {
	return clone(src)
}

func (src *Pixel) Gray() int {
	return int(*src.src)
}

//MUST FIX!!!!!!!!!!!
func (src *Pixel) Set(r, g, b int) {
	*src.src = uint8(r)
	*src.src = uint8(g)
	*src.src = uint8(b)
}

func (src *Img) Loop(some func(int, int, *Pixel)) {
	bounds := src.Bounds
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			some(x, y, src.At(x, y))
		}
	}
}
