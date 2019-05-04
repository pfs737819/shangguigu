package main
import (
	"fmt"
	"reflect"
	// "encoding/json"
)

func test(t interface{}) {
	rtype := reflect.ValueOf(t)
	
	iv := rtype.Interface()

	// num2 := iv.(int)
	fmt.Printf("iv%T,%d",iv,iv)
}

func main() {
	num := 10
	test(num)
}