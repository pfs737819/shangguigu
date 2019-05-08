package main
import (
	"fmt"
)



func test(args... interface{}) {
	fmt.Println(args)
}

func main() {
	test(1,"2")
}