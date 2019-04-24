package main
import (
	"fmt"
	"go_code/extend/model"
)



func main() {
	P := &(model.Pupil{})
	P.Name = "张三"
	P.Age = 18
	P.Score = 78.8
	P.SetScore(60)
	P.ShowInfo()
	P.Testing()
	fmt.Println(*P)



	// var g = model.Graduate.Student{
	// 	Name : "战三",
	// 	Age : 28,
	
	// }
	// g.SetScore(60)
	// g.ShowInfo()
	// fmt.Println(g)
}