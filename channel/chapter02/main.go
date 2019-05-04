package main
import (
	"fmt"
)

type Cat struct {
	Name string
	Age  int
}

func main(){

	chanInt := make(chan interface{},3)
	cat := Cat{
		Name:"大脸猫",
		Age:3,
	}
	chanInt <- 10
	chanInt <- 11
	chanInt <- cat

	<- chanInt
	<- chanInt
	newCat := <- chanInt

	catCat := newCat.(Cat)//类型断言

	fmt.Printf("newCat类型 %T,newCat %v",newCat,catCat.Name)

}