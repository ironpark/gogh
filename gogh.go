// gogh project gogh.go
package gogh

import (
	//"fmt"
	"image"
	"log"
)

const (
	GRAY    = 0
	GRAY16  = 1
	NRGBA   = 3
	NRGBA64 = 4
	RGBA    = 5
	RGBA64  = 6
)

func Color(colors ...uint8) []uint8 {
	return colors
}
func NewImg(rect image.Rectangle, T int) *Img {
	var (
		im     []uint8
		stride int = 1
	)

	switch T {
	case GRAY:
		im = image.NewGray(rect).Pix
		stride = 1
	case GRAY16:
		im = image.NewGray16(rect).Pix
		stride = 2
	case NRGBA:
		im = image.NewNRGBA(rect).Pix
		stride = 4
	case NRGBA64:
		im = image.NewNRGBA64(rect).Pix
		stride = 8
	case RGBA64:
		im = image.NewRGBA(rect).Pix
		stride = 8
	case RGBA:
		im = image.NewRGBA64(rect).Pix
		stride = 8
	default:
		return nil
	}

	return &Img{im, T, rect.Max.X, rect.Max.Y, rect, stride}
}

type Img struct {
	Pixels    []uint8
	ImageType int
	Width     int
	Height    int
	Bounds    image.Rectangle
	stride    int
}

type Pixel struct {
	src       []*uint8
	X         int
	Y         int
	ImageType int
}

func (img *Img) At(x, y int) *Pixel {
	//if out pixel!!
	if !(image.Point{x, y}.In(img.Bounds)) {
		log.Fatal("out of pixel")
		return nil
	}

	switch img.ImageType {

	case NRGBA, RGBA:
		index := img.Width*y*4 + x*4
		//fmt.Println(img.Pixels[index+x], img.Pixels[index+x+1], img.Pixels[index+x+2], img.Pixels[index+x+3])
		return &Pixel{[]*uint8{
			&img.Pixels[index+0],
			&img.Pixels[index+1],
			&img.Pixels[index+2],
			&img.Pixels[index+3],
		}, x, y, img.ImageType}
	case GRAY:
		return &Pixel{[]*uint8{
			&img.Pixels[img.Width*y+x],
		}, x, y, img.ImageType}
	default:
		return &Pixel{[]*uint8{
			&img.Pixels[img.Width*y+x],
		}, x, y, img.ImageType}
	}
}

//TEMP!!! YOU MUST FIX!!
func (src *Pixel) RGBA() (int, int, int, int) {
	var r, g, b, a uint8
	if GRAY == src.ImageType {
		r = *src.src[0]
		g = *src.src[0]
		b = *src.src[0]
		a = 255
	} else {
		r = *src.src[0]
		g = *src.src[1]
		b = *src.src[2]
		a = *src.src[3]
	}
	return int(r), int(g), int(b), int(a)
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
	return int(*src.src[0])
}

//Slow?
func (src *Pixel) Set(color ...interface{}) {
	if len(color) == 1 {
		switch v := color[0].(type) {
		case []uint8:
			if len(color) != len(src.src) {
				log.Fatal("color type different")
			}
			for i := range src.src {
				*src.src[i] = v[i]
			}
		case uint8:
			*src.src[0] = v
		case int:
			*src.src[0] = uint8(v)
		}
	} else {

		if src.ImageType == GRAY {
			switch v := color[0].(type) {
			case uint8:
				*src.src[0] = v
			case int:
				*src.src[0] = uint8(v)
			}
		} else {
			for i, item := range src.src {
				if len(color) > i {
					if u, e := color[i].(int); e {
						*item = uint8(u)
					} else if u, e := color[i].(uint8); e {
						*item = u
					} else {
						log.Fatal("argument is not uint8 or int")
					}
				}
			}
		}
	}
}

func (src *Img) Loop(some func(int, int, *Pixel)) {
	bounds := src.Bounds
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			some(x, y, src.At(x, y))
		}
	}
}
