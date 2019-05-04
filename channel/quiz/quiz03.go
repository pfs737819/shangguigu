package main
import (
	"fmt"
	// "time"
)

func main(){

	intChan := make(chan int,1000)
	resChan := make(chan int,4000)
	stsChan := make(chan bool,4)

	go func(){
		for i := 2;i <= 1000;i ++ {
			intChan <- i
		} 
		close(intChan)
	}()

	

	for j := 0;j < 4;j ++ {
		
		go func(){
			for {
				flag := true
				num, ok := <- intChan
				if !ok {
					break //无法取出数据时退出循环
				}
	
				for i := 2;i < num;i ++ {
					if num % i == 0 {
						flag = false
						break
					}
				}
	
				if flag {
					resChan <- num
				}
			}
	
			// stsChan <- true
			// fmt.Printf("协程%d已退出\n",i)
			// time.Sleep(time.Millisecond*10)
			// for num := range intChan {
			// 	flag := true
				
			// 	for i := 2;i < num;i ++ {
			// 		if num % i == 0 {
			// 			flag = false
			// 			break
			// 		}
			// 	}
	
			// 	if flag {
			// 		resChan <- num
			// 	}
			// }
			fmt.Printf("协程%d已退出\n",j)
			stsChan <- true
			
		}()
	}

	go func(){
		for i := 0;i < 4;i ++ {
			<- stsChan
		}
		close(resChan)
		close(stsChan)
	}()

	for v := range resChan {
		fmt.Println(v)
	}
	// for i := 0;i < 4;i ++ {
	// 	fmt.Printf("正在读取第%d次,值为%v...\n",i,<-intChan)
	// }
	// num := 9
	// sum := (1+num)*num/2
	// fmt.Println(sum)
	// m := make(map[int]int,9)
	// resChan := make(chan map[int]int,9)
}
