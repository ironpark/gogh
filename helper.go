package gogh

import (
	"image"
	//_ "image/bmp"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
)

func Save(path string, src image.Image) {
	file, err := os.Create(path)

	if err != nil {
		log.Fatal(err)
	}
	png.Encode(file, src)
	defer file.Close()
}
func Load(path string) *Mat {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	return NewMat(ImageToRGBA(m))
}

func ImageToRGBA(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	model := rgba.ColorModel()
	for i := 0; i < bounds.Max.X; i++ {
		for j := 0; j < bounds.Max.Y; j++ {
			rgba.Set(i, j, model.Convert(img.At(i, j)))
		}
	}
	return rgba
}

func clone(img *Mat) *Mat {
	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	for i := 0; i < bounds.Max.X; i++ {
		for j := 0; j < bounds.Max.Y; j++ {
			rgba.Set(i, j, img.At(i, j).color)
		}
	}
	return NewMat(rgba)
}
