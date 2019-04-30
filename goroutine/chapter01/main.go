package main
import (
	"fmt"
	"time"
)

func test() {
	for i := 0;i < 10;i ++ {
		fmt.Printf("test() %d\n",i)
		time.Sleep(time.Second)
	}
}

func main(){
	go test()

	for i := 0;i < 10;i ++ {
		fmt.Printf("main() %d\n",i)
		time.Sleep(time.Second)
	}

}