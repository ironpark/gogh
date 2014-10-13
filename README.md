gogh
====

image processing package write in go

```go
import (
	"fmt"
	"github.com/ironpark/gogh"
)


func main() {
	src := gogh.Load("some.jpg")
	src.Binarization(50, false).Save("Binarization.png")
	fmt.Println("histogram",src.Histogram())
}
```
##pixel control

```go
import (
	"fmt"
	"github.com/ironpark/gogh"
)


func main() {
		src := gogh.Load("some.jpg")
		//src.At(x,y).Set(r,g,b)
		src.At(1,2).Set(0,0,0)
		
		//src.At(x,y).Set(r,g,b)
		fmt.Println(src.At(1,2).Gray())
		fmt.Println(src.At(1,2).RGBA())
}
```
