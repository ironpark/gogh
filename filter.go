package gogh

import (
	"github.com/ironpark/gogh/filter"
	"math"
)

const (
	BLUR_BOX = 0
)

func (src *Img) Filter(kernel [][]float32) *Img {
	size := len(kernel)
	//kernel center point
	kernelCP := int(float32(size) / float32(2))
	normalization := false
	normal := float32(0)
	for _, Es := range kernel {
		for _, E := range Es {
			normal += E
		}
	}

	if normal > 1 {
		normalization = true
	}

	v := clone(src)
	bounds := src.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			//Convolution
			cPixel := float32(0)
			for ky, Es := range kernel {
				for kx, E := range Es {
					xe1 := kx - kernelCP
					ye2 := ky - kernelCP

					selx := x + xe1
					sely := y + ye2

					if selx < 0 {
						selx = 0
					}
					if selx > (bounds.Max.X - 1) {
						selx = (bounds.Max.X - 1)
					}
					if sely < 0 {
						sely = 0
					}
					if sely > (bounds.Max.Y - 1) {
						sely = (bounds.Max.Y - 1)
					}
					gray := src.At(selx, sely).Gray()
					cPixel += float32(gray) * E
				}
			}
			//normalization
			if normalization {
				cPixel = cPixel / normal
			}
			c := int(math.Abs(float64(cPixel)))
			if 255 < c {
				c = 255
			}
			//set pixel

			v.At(x, y).Set(c, c, c)
		}
	}
	src = v
	return v
}

func (src *Img) Blur(size, blurtype int) *Img {
	switch blurtype {
	case BLUR_BOX:
		return src.Filter(filter.GenBoxBlurMask(size))
	}
	//unknown filter
	return nil
}
