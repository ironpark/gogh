gogh
====

image processing package write in go

##How To Install?
```
go get github.com/ironpark/gogh
```
##How To Update?
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
var (
	sobel = [][]float32{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}
)

func main() {
	fmt.Println("Hello World!")
	img := gogh.Load("4.jpg")
	img.Filter(sobel).Histogram().Cumulative()
}
```