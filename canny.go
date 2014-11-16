package gogh

import (
	"math"
)

var (
	gaussianMask5x5 = [][]float32{
		{2, 4, 5, 4, 2},
		{4, 9, 12, 9, 4},
		{5, 12, 15, 12, 5},
		{4, 9, 12, 9, 4},
		{2, 4, 5, 4, 2},
	}
	gaussianMask3x3 = [][]float32{
		{1, 2, 1},
		{2, 4, 2},
		{1, 2, 1},
	}
	sobel_x = [][]float64{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}
	sobel_y = [][]float64{
		{-1, -2, -1},
		{0, 0, 0},
		{1, 2, 1},
	}
)

func gray(src *Img, x, y int) float64 {
	bounds := src.Bounds
	if x < 0 {
		x = 0
	}
	if x > bounds.Max.X {
		x = bounds.Max.X
	}
	if y < 0 {
		y = 0
	}
	if y > bounds.Max.Y {
		y = bounds.Max.Y
	}
	return float64(src.At(x, y).Gray())
}
func (src *Img) Canny(th_high, th_low int) *Img {
	w := src.Width
	h := src.Height

	G := src.Filter(gaussianMask5x5)

	//gradx := G.Clone().Filter(sobel_x)
	//grady := G.Clone().Filter(sobel_y)

	mag := make([][]int, w)
	dir := make([][]int, w)

	dxMat := make([][]int, w)
	dyMat := make([][]int, w)

	gnh := make([][]int, w)
	gnl := make([][]int, w)

	for i := range mag {
		mag[i] = make([]int, h)
		dir[i] = make([]int, h)
		dxMat[i] = make([]int, h)
		dyMat[i] = make([]int, h)
		gnh[i] = make([]int, h)
		gnl[i] = make([]int, h)
	}
	/* Determine edge directions and gradient strengths */
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			a1, a2, a3 := gray(G, x-1, y-1), gray(G, x, y-1), gray(G, x+1, y-1)
			b1, b3 := gray(G, x-1, y), gray(G, x+1, y-1)
			c1, c2, c3 := gray(G, x-1, y+1), gray(G, x, y+1), gray(G, x+1, y+1)
			// -1 0 1
			// -2 0 2
			// -1 0 1
			Gx := a1*-1 + a3 + b1*-2 + b3*2 + c1*-1 + c3

			// -1 -2 -1
			//  0  0  0
			//  1  2  1
			Gy := a1*-1 + a2*-2 + a3*-1 + c1 + c2*2 + c3
			dxMat[x][y] = int(Gx)
			dyMat[x][y] = int(Gy)
			m := math.Abs(Gx) + math.Abs(Gy)
			//if m < 255 {
			mag[x][y] = int(m)
			//} else {
			//	mag[x][y] = 255
			//}
		}
	}
	const fbit = 10
	const tan225 = 424
	const tan675 = 2472

	for y := 1; y < h-1; y++ {
		for x := 1; x < w-1; x++ {
			direction := 0
			src.At(x, y).Set(0, 0, 0)
			if mag[x][y] > th_low {
				dx := dxMat[x][y]
				dy := dyMat[x][y]
				if dx != 0 {
					slope := (dy << fbit) / dx
					if slope > 0 {
						if slope < tan225 {
							direction = 0
						} else if slope < tan675 {
							direction = 1
						} else {
							direction = 2
						}
					} else {
						slope = slope * -1
						if slope > tan675 {
							direction = 2
						} else if slope > tan225 {
							direction = 3
						} else {
							direction = 0
						}
					}
				} else {
					direction = 2
				}
				bMaxima := true
				// perform non-maxima suppression
				switch direction {
				case 0:
					if mag[x][y] < mag[x-1][y] || mag[x][y] < mag[x+1][y] {
						bMaxima = false
					}
				case 1:
					if mag[x][y] < mag[x+1][y+1] || mag[x][y] < mag[x-1][y-1] {
						bMaxima = false
					}
				case 2:
					if mag[x][y] < mag[x][y+1] || mag[x][y] < mag[x-1][y] {
						bMaxima = false
					}
				default: // 3
					if mag[x][y] < mag[x-1][y+1] || mag[x][y] < mag[x+1][y-1] {
						bMaxima = false
					}
				}
				if bMaxima {
					if mag[x][y] > th_high {
						src.At(x, y).Set(255, 255, 255)
						gnh[x][y] = mag[x][y]
						gnl[x][y] = 0
					} else {
						//src.At(x, y).Set(100, 100, 100)
						gnl[x][y] = mag[x][y]
					}
				}

			}

		}

	}
	temp := make([][]int, w)

	for i := range temp {
		temp[i] = make([]int, h)
	}
	for y := 1; y < h-1; y++ {
		for x := 1; x < w-1; x++ {
			if gnl[x][y] != 0 {

				//remove noise
				b1, b2, b3 := gnl[x-1][y-1], gnl[x][y-1], gnl[x+1][y-1]
				b4, b5 := gnl[x-1][y], gnl[x+1][y]
				b6, b7, b8 := gnl[x-1][y+1], gnl[x][y+1], gnl[x+1][y+1]

				if b1 == 0 && b2 == 0 && b3 == 0 &&
					b4 == 0 && b5 == 0 &&
					b6 == 0 && b7 == 0 && b8 == 0 {
					gnl[x][y] = 0
					continue
				}

				a1, a2, a3 := gnh[x-1][y-1], gnh[x][y-1], gnh[x+1][y-1]
				a4, a5 := gnh[x-1][y], gnh[x+1][y]
				a6, a7, a8 := gnh[x-1][y+1], gnh[x][y+1], gnh[x+1][y+1]

				if a1 != 0 || a2 != 0 || a3 != 0 ||
					a4 != 0 || a5 != 0 ||
					a6 != 0 || a7 != 0 || a8 != 0 {
					gnl[x][y] = 0
					gnh[x][y] = 1
					src.At(x, y).Set(255, 255, 255)
				}

			}
		}
		for x := w - 1; x > 0; x-- {
			if gnl[x][y] != 0 {
				a1, a2, a3 := gnh[x-1][y-1], gnh[x][y-1], gnh[x+1][y-1]
				a4, a5 := gnh[x-1][y], gnh[x+1][y]
				a6, a7, a8 := gnh[x-1][y+1], gnh[x][y+1], gnh[x+1][y+1]
				if a1 != 0 || a2 != 0 || a3 != 0 ||
					a4 != 0 || a5 != 0 ||
					a6 != 0 || a7 != 0 || a8 != 0 {
					gnl[x][y] = 0
					gnh[x][y] = 1
					src.At(x, y).Set(255, 255, 255)
				}

			}
		}
	}
	for y := h - 1; y > 1; y-- {
		for x := 1; x < w-1; x++ {
			if gnl[x][y] != 0 {
				a1, a2, a3 := gnh[x-1][y-1], gnh[x][y-1], gnh[x+1][y-1]
				a4, a5 := gnh[x-1][y], gnh[x+1][y]
				a6, a7, a8 := gnh[x-1][y+1], gnh[x][y+1], gnh[x+1][y+1]

				if a1 != 0 || a2 != 0 || a3 != 0 ||
					a4 != 0 || a5 != 0 ||
					a6 != 0 || a7 != 0 || a8 != 0 {
					gnl[x][y] = 0
					gnh[x][y] = 1
					src.At(x, y).Set(255, 255, 255)
				}

			}
		}
		for x := w - 1; x > 0; x-- {
			if gnl[x][y] != 0 {
				a1, a2, a3 := gnh[x-1][y-1], gnh[x][y-1], gnh[x+1][y-1]
				a4, a5 := gnh[x-1][y], gnh[x+1][y]
				a6, a7, a8 := gnh[x-1][y+1], gnh[x][y+1], gnh[x+1][y+1]
				if a1 != 0 || a2 != 0 || a3 != 0 ||
					a4 != 0 || a5 != 0 ||
					a6 != 0 || a7 != 0 || a8 != 0 {
					gnl[x][y] = 0
					gnh[x][y] = 1
					src.At(x, y).Set(255, 255, 255)
				}

			}
		}
	}

	return src
}

func traverse(src *Img, gnh, gnl [][]int, i, j int) {
	x := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	y := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	for k := 0; k < 8; k++ {
		if gnh[i+x[k]][j+y[k]] == 0 && gnl[i+x[k]][j+y[k]] != 0 {
			src.At(i+x[k], j+y[k]).Set(255, 255, 255)
			traverse(src, gnh, gnl, i+x[k], j+y[k])
		}
	}
}
