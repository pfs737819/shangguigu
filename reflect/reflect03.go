package main
import (
	"fmt"
	"reflect"
)



func test(t interface{}) {
	rtype := reflect.ValueOf(t)
	
	// iv := rtype.Interface()

	// num2 := iv.(int)
	fmt.Printf("%v",rtype.Kind())
	rtype.Elem().SetInt(20)
	
}

func main() {
	num := 10
	test(&num)
	//fmt.Printf("iv%T,%d",num,num)
}