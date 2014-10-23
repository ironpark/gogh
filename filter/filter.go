package filter

var (
	//Gaussian
	GaussianMask5x5 = [][]float32{
		{2, 4, 5, 4, 2},
		{4, 9, 12, 9, 4},
		{5, 12, 15, 12, 5},
		{4, 9, 12, 9, 4},
		{2, 4, 5, 4, 2},
	}
	GaussianMask3x3 = [][]float32{
		{1, 2, 1},
		{2, 4, 2},
		{1, 2, 1},
	}
	//Sobel 3x3,5x5
	SobelMask3x3X = [][]float64{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}
	SobelMask3x3Y = [][]float64{
		{-1, -2, -1},
		{0, 0, 0},
		{1, 2, 1},
	}
	SobelMask5x5X = [][]float64{
		{1, 2, 0, -2, -1},
		{4, 8, 0, -8, -4},
		{6, 12, 0, -12, -6},
		{4, 8, 0, -2, -4},
		{1, 2, 0, -2, -1},
	}
	SobelMask5x5Y = [][]float64{
		{-1, -4, -6, -4, -1},
		{-2, -8, -12, -8, -2},
		{0, 0, 0, 0, 0, 0},
		{2, 8, 12, 8, 2},
		{1, 4, 6, 4, 1},
	}
)

func GenBoxBlurMask(size int) [][]float32 {
	filter := make([][]float32, size)
	for i := 0; i < size; i++ {
		filter[i] = make([]float32, size)
		for j := 0; j < size; j++ {
			filter[i][j] = 1.0
		}
	}
	return filter
}
