package main
import (
	"fmt"
	_"time"
	"sync"
)

var Wg sync.WaitGroup

func Write(numChan chan int,num int) {
	for i := 1;i <= num;i ++ {
		numChan <- i
	}
	close(numChan)
}

func Read(numChan chan int,resChan chan map[int]int) {
	for {
		num, ok := <- numChan
		if !ok {
			break
		}
		sum := (1+num)*num/2
		m := make(map[int]int)
		m[num] = sum

		resChan <- m
	}

	Wg.Done()
	
}






func main(){
	Wg.Add(8)
	numChan := make(chan int,2000)
	resChan := make(chan map[int]int,2000)
	count := 8
	 
	go Write(numChan,2000)


	for i := 0;i < count;i ++{
		go Read(numChan,resChan)
	}
	Wg.Wait()

	close(resChan)

	for v := range resChan{
		fmt.Println(v)
	}
	// for {
	// 	v,ok := <- resChan
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(v)
	// }
	
}
