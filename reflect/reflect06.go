package main
import (
	"reflect"
	"fmt"
)

type Cal struct {
	Num1 int
	Num2 int
}

func (c Cal)GetSub(num1,num2 int) int {
	return num1 + num2
}

func ReflectFunction(c interface{}) {
	calVal := reflect.ValueOf(c)
	calType := reflect.TypeOf(c)
	fmt.Println(calVal,calType)
	numField := calVal.NumField()
	for i := 0;i < numField;i ++ {
		fmt.Printf("Field %d 的值为 %v\n",i,calVal.Field(i))
	}

	parameter := []reflect.Value{reflect.ValueOf(2),reflect.ValueOf(3)}
	rs := calVal.MethodByName("GetSub").Call(parameter)
	fmt.Printf("GetSub 的值为 %v\n",rs[0].Int())
}

func main(){
	cal := Cal{
		Num1:1,
		Num2:2,
	}

	ReflectFunction(cal)
}