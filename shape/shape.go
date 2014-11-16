package shape

type Shape interface {
	Width() int
	Height() int
}

type Circle struct {
	Radius int
}

func (src Circle) Width() int {
	return src.Radius * 2
}
func (src Circle) Height() int {
	return src.Radius * 2
}

type Rect struct {
	W int
	H int
}

func (src Rect) Width() int {
	return src.W
}
func (src Rect) Height() int {
	return src.H
}

type Plus struct {
	Size int
}

func (src Plus) Width() int {
	return src.Size*2 + 1
}
func (src Plus) Height() int {
	return src.Size*2 + 1
}
