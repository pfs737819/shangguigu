package main
import (
	"fmt"
	"math/rand"
	"sort"
)

type Student struct{
	Name string
	Age  int
	Score float64
}

type StuSlice []Student

func (p StuSlice) Len() int{
	return len(p)
}

func (p StuSlice) Less(i, j int) bool{
	return p[i].Score > p[j].Score  
}

func (p StuSlice) Swap(i, j int){
	p[i], p[j] = p[j], p[i] 
}

func main() {

	var StuSlices StuSlice;
	for i := 0;i < 10;i ++ {
		Students := Student{
			Name: fmt.Sprintf("Student-%d",rand.Intn(100)),
			Age: rand.Intn(100),
			Score: float64(rand.Intn(100)),
		}

		StuSlices = append(StuSlices,Students)
	}
	sort.Sort(StuSlices)
	for _,v := range StuSlices{
		fmt.Println(v)
	}

}