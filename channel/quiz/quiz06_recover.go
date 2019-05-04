package main
import(
	"fmt"
	"time"
)

func Say() {
	for i := 0;i <= 10;i ++ {
		time.Sleep(time.Millisecond*100)
		fmt.Println("Say",i)
	}
}

func Do() {
	defer func(){
		err := recover()
		if err != nil {
			fmt.Println("Do发生错误",err)
		}
	}()
	var myMap map[int]string
	myMap[0] = "11111"
}

func main() {
	go Say()
	go Do()

	time.Sleep(time.Second*5)
}