package main
import (
	"fmt"
	// "time"
)

func Write(chanInt chan<- int) {
	chanInt <- 1
	chanInt <- 2
	
}

func Read(chanInt <-chan int) {
	fmt.Println(<-chanInt)
	fmt.Println(<-chanInt)
}

func main(){
	chanInt := make(chan int,2)

	Write(chanInt)
	Read(chanInt)
	
}