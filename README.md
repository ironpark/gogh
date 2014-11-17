gogh
====

image processing package write in go
for Computer Vision (like OpenCV)

##How To Install/Update?
```
go get -u github.com/ironpark/gogh
```
##How To Use?
### Simple Examples
**Binarization/Histogram**
```go
import (
	"fmt"
	"github.com/ironpark/gogh"
)


func main() {
	src := gogh.Load("some.jpg")
	//method chaining pattern!
	fmt.Println("histogram",src.Histogram().Array()) //Nomal histogram
	fmt.Println("histogram",src.Histogram().Cumulative()) //Cumulative histogram
	
	//Binarization
	src.Binarization(50, false).Save("Binarization.png")
}
```
**Pixel Get/Set**
```go
import (
	"fmt"
	"github.com/ironpark/gogh"
)


func main() {
		src := gogh.Load("some.jpg")
		//src.At(x,y).Set(r,g,b)
		src.At(1,2).Set(0,0,0)
		
		fmt.Println(src.At(1,2).Gray())
		fmt.Println(src.At(1,2).RGBA())
}
```
**Pixel Loop**

if you want loop all pixels
you can use double for loop

```go
	src := gogh.Load("some.jpg")
	bounds := src.Bounds()
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			fmt.Println(c.At(x, y).Gray()
		}
	}
```

it is same fuction
more than simple

```go
	src.Loop(func(x, y int) {
		fmt.Println(src.At(x, y).Gray())
	})
```

**Blur**
```go
import (
	"fmt"
	"github.com/ironpark/gogh"
)

func main() {
	img := gogh.Load("some.jpg")
	img.Blur(3, gogh.BLUR_BOX).Filter(sobel).Save("5.png")
}
```
**Convolution Filter**
```go
import (
	"fmt"
	"github.com/ironpark/gogh"
	"github.com/ironpark/gogh/mask"
)
var (
	sobel = [][]float32{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}
)

func main() {
	img := gogh.Load("some.jpg")
	
	img.Filter(sobel).Save("Sobel1.png")
	//same code
	img.Filter(mask.SobelMask3x3X).Save("Sobel2.png")
	
	//Box Blur
	img.Filter(mask.GenBoxBlurMask(3)).Save("BoxBlur1.png")
	//or
	img.Blur(3, gogh.BLUR_BOX).Save("BoxBlur2.png")
	
}
```