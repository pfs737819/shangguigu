package main
import (
	"fmt"
	"math/rand"
)

type Person struct {
	Name string
	Age  int
	Adderss string
}



func main(){
	personChan := make(chan Person,10) 
	for i := 1;i <= 10;i ++ {
		num := rand.Intn(1000)
	
		person := Person{
			Name:fmt.Sprintf("张三%v",num),
			Age:num,
			Adderss:fmt.Sprintf("上海%v",num),
		}

		personChan <- person
	}

	
	// for  {
	// 	v ,ok := <- personChan
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(v,ok)
	// }

	count := 10
	for v := range personChan {
		count -- 
		if count == 0{
			close(personChan)
		}
		fmt.Println(v)
	}

	
}