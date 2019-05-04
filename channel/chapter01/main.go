package main
import (
	"fmt"
)

func main(){

	var chanInt chan int
	chanInt = make(chan int,3)
	chanInt <- 10
	chanInt <- 11
	chanInt <- 12

	num1 := <- chanInt
	num2 := <- chanInt
	num3 := <- chanInt

	fmt.Printf("chanInt容量 %v,num1 %v,num2 %v,num3 %v",len(chanInt),num1,num2,num3)

}