package main
import (
	"fmt"
	"reflect"
)

func test01(itf interface{}) {
	ref := reflect.ValueOf(itf)
	kind := ref.Kind()
	getType := ref.Type()

	vInterface := ref.Interface()
	value, ok := vInterface.(float64)
	// ref.SetFloat(1.3)
	if !ok {
		fmt.Println("no  ok")
	}
	if kind == reflect.Float64 {
		fmt.Println("  ok")
	}else{
		fmt.Println("no  ok")
	}

	fmt.Println("kind",kind)
	fmt.Println("kind2",reflect.Float64)
	fmt.Println("getType",getType)
	fmt.Println("断言",value)
	fmt.Printf("类型\\T %T \n",vInterface)
	
}

func main() {
	var v float64 = 1.2
	test01(v)
	// a := reflect.ValueOf(&v)
	// a = a.Elem()
	// a.SetFloat(1.3)
	// fmt.Println(v)
}