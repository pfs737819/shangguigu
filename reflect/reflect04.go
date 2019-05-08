package main
import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name string `json:"name"`
	Age int
	Skill string
}

func (m Monster) Print() {
	fmt.Println("-------Start-------")
	fmt.Println("...................")
	fmt.Println("--------End--------")
}

func (m Monster) GetSum(n1, n2 int) int { 
	return n1+n2
}

func TestStruct(Itf interface{}) {
	reType := reflect.TypeOf(Itf)
	reValue := reflect.ValueOf(Itf)
	kd := reValue.Kind()
	
	fmt.Printf("reflect.TypeOf %v\n",reType)

	if kd != reflect.Struct {
		fmt.Println("expect Struct")
		return
	}

	numField := reValue.NumField()
	fmt.Printf("Struct has Field %v\n",numField)

	for i := 0;i < numField;i ++ {
		fmt.Printf("Filed %d:的值为=%v\n",i,reValue.Field(i))
		tagVal := reType.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Filed %d:的值为=%v tag的值为=%v\n",i,reValue.Field(i),tagVal)
		}
	}

	//返回该对象下的方法数量
	numMethod := reValue.NumMethod()
	fmt.Printf("Struct has Method %v\n",numMethod)

	var params []reflect.Value
	n1 := reflect.ValueOf(10)
	n2 := reflect.ValueOf(1)
	params = append(params,n1)
	params = append(params,n2)
	res := reValue.Method(0).Call(params)

	fmt.Printf("Method res %v\n",res[0])
}

func main() {
	monster := Monster{
		Name:"孙悟空",
		Age:1000,
		Skill:"七十二变",
	}

	TestStruct(monster)
}