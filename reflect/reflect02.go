package main
import (
	"fmt"
	"reflect"
	// "encoding/json"
)


type Student struct {
	Name string
	Age  int
}

func test(t interface{}) {
	rtype := reflect.ValueOf(t)
	ttype := reflect.TypeOf(t)
	fmt.Println(ttype.Name())
	iv := rtype.Interface()
	stu, ok := iv.(Student)
	if !ok {
		fmt.Println(ok)
	}
	// num2 := iv.(int)
	fmt.Printf("%T,%v",iv,stu.Name)
}

func main() {
	student := Student{
		Name:"张三",
		Age:20,
	}
	test(student)
}