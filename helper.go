package gogh

import (
	"image"
	//_ "image/bmp"
	//"fmt"
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
	var stride int = 1
	switch v := m.(type) {
	case *image.RGBA:
		imgType = RGBA
		pixels = v.Pix
		stride = 4
	case *image.RGBA64:
		imgType = RGBA64
		pixels = v.Pix
		stride = 8
	case *image.NRGBA:
		imgType = NRGBA
		pixels = v.Pix
		stride = 4
	case *image.NRGBA64:
		imgType = NRGBA64
		pixels = v.Pix
		stride = 8
	case *image.Gray:
		imgType = GRAY
		pixels = v.Pix
		stride = 1
	case *image.Gray16:
		imgType = GRAY16
		pixels = v.Pix
		stride = 2

	}
	return &Img{pixels, imgType, m.Bounds().Max.X, m.Bounds().Max.Y, m.Bounds(), stride}
}

func clone(img *Img) *Img {
	pix := img.Pixels[:]
	return &Img{pix, img.ImageType, img.Width, img.Height, img.Bounds, img.stride}
}
