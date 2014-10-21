package gogh

import (
	//"image"
	"fmt"
)

type histogram struct {
	array      []int
	cumulative []int
	src        *Img
	high       int
	low        int
}

func (src *Img) Histogram() *histogram {
	bounds := src.Bounds()
	histo := make([]int, 256)

	high := 0
	low := 255
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := src.At(x, y) //.Gray()
			gray := pixel.Gray()
			histo[gray]++

			if high < gray {
				high = gray
			}
			if low > gray {
				low = gray
			}
		}
	}
	return &histogram{histo, nil, src, high, low}
}

//Histogram Array
func (src *histogram) Array() []int {
	return src.array
}

//Cumulative Histogram Array
func (src *histogram) Cumulative() []int {
	if src.cumulative != nil {
		return src.cumulative
	} else {
		src.cumulative = make([]int, 256)
		cumulate := 0
		for i, volume := range src.array {
			cumulate += volume
			src.cumulative[i] = cumulate
		}
		return src.cumulative
	}
}

//Cumulative Histogram graph Img
func (src *histogram) CumulativeGraph() {
}

func (src *histogram) Graph() []int {
	if src.cumulative != nil {
		return src.cumulative
	} else {
		src.cumulative = make([]int, 256)
		cumulate := 0
		for i, volume := range src.array {
			cumulate += volume
			src.cumulative[i] = cumulate
		}
		return src.cumulative
	}
}

//Contrast Stretching
func (histo *histogram) Stretching() *Img {
	//high - low
	hml := (histo.high - histo.low)
	bounds := histo.src.Bounds()
	fmt.Println(histo.low, histo.high)
	/*
		new pixel = (oldpixel - low/high - low)*255
	*/
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := histo.src.At(x, y)
			gray := pixel.Gray()
			gray = int(float32(gray-histo.low) / float32(hml) * 255)
			//gray = (gray - histo.low) / hml * 255

			pixel.Set(gray, gray, gray)
		}
	}
	//히스토그램 변경사항을 구조체에 반영하는 코드 필요 .
	return histo.src
}
