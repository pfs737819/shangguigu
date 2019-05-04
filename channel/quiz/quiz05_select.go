package main
import(
	"fmt"
	// "time"
)

func main() {
	intChan := make(chan int,50)
	StringChan := make(chan string,50)


	go func(){
		for i := 0;i < 100;i ++ {
		intChan <- i
		}
	}()


	go func(){
		for i := 0;i < 100;i ++ {
			StringChan <- fmt.Sprintf("hello%v",i)
		}
	}()

	
	for {
		select {
			case v := <- intChan :
				fmt.Printf("intChan取数据 %v\n",v)
			case v := <- StringChan :
				fmt.Printf("StringChan取数据 %v\n",v)
			default :
				fmt.Printf("没有数据了\n")
				return

		}
	}
	
}