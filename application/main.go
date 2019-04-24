package main
import (
	"fmt"
)

type Hello struct{
	Name string
	Age  int
}


func main() {

var a []Hello 

var b = Hello{
	Name:"123",
	Age:10,
}
a = append(a,b)

fmt.Println(a)

}