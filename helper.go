package gogh

import (
	"image"
	//_ "image/bmp"
	"fmt"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
)

func save(path string, src image.Image) {
	file, err := os.Create(path)

	if err != nil {
		log.Fatal(err)
	}
	png.Encode(file, src)
	defer file.Close()
}

func Load(path string) *Img {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	var pixels []uint8
	var imgType int = 100
	switch v := m.(type) {
	case *image.RGBA:
		imgType = RGBA
		pixels = v.Pix
	case *image.RGBA64:
		imgType = RGBA64
		pixels = v.Pix
	case *image.NRGBA:
		imgType = NRGBA
		pixels = v.Pix
	case *image.NRGBA64:
		imgType = NRGBA64
		pixels = v.Pix
	case *image.Gray:
		imgType = GRAY
		pixels = v.Pix
	case *image.Gray16:
		imgType = GRAY16
		pixels = v.Pix

	}
	fmt.Println("FUCK", imgType, m.Bounds().Max.X, m.Bounds().Max.Y)
	return &Img{pixels, imgType, m.Bounds().Max.X, m.Bounds().Max.Y, m.Bounds()}
}

func ImageToNRGBA(img image.Image) *image.NRGBA {
	bounds := img.Bounds()
	if bounds.Min.X == 0 && bounds.Min.Y == 0 {
		if src0, ok := img.(*image.NRGBA); ok {
			return src0
		}
	}

	rgba := image.NewNRGBA(bounds)
	model := rgba.ColorModel()
	for i := bounds.Min.Y; i < bounds.Max.X; i++ {
		for j := bounds.Min.X; j < bounds.Max.Y; j++ {
			rgba.Set(i, j, model.Convert(img.At(i, j)))
		}
	}
	return rgba
}

func clone(img *Img) *Img {
	pix := img.Pixels[:]
	return &Img{pix, img.ImageType, img.Width, img.Height, img.Bounds}
}
