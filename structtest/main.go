package main
import (
	"fmt"
)

type MethodUtils struct{

}

type Student struct {
	Name string
	Gender string
	Age int
	Id int
	Score float64
}

func (stu *Student) Say()string {
	rs := fmt.Sprintf("student信息 name=[%v] gender=[%v] age=[%v] id=[%v] score=[%v]",stu.Name,stu.Gender,
	stu.Age,stu.Id,stu.Score)
	return rs
}


func (t MethodUtils) Method1(m ,n int){
	for i := 0;i < m;i ++ {
		for i := 0;i < n;i ++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (t MethodUtils) area(m ,n int) int {
	return m*n
}

func (t MethodUtils) Rs(n int){
	for i := 1;i <= n;i ++ {
		for j := 1;j <= i;j ++ {
			fmt.Printf("%v x %v = %v\t",j,i,j*i)
		}
		fmt.Println()
	}	
}

func main() {
	stu := Student{
		Name:"张三",
		Gender:"男",
		Age:18,
		Id:1,
		Score:64.5,
	}
	fmt.Println(stu.Say())
}