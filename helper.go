package gogh

import (
	"image"
	//_ "image/bmp"
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
	return &Img{ImageToNRGBA(m)}
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
	bounds := img.Bounds()
	rgba := image.NewNRGBA(bounds)
	for i := 0; i < bounds.Max.X; i++ {
		for j := 0; j < bounds.Max.Y; j++ {
			rgba.Set(i, j, img.At(i, j).color)
		}
	}
	return &Img{rgba}
}
