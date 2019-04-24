package main
import (
	"fmt"
	"go_code/factory/model"
)
func main() {
	stu := model.NewStudent("李四",65.5)
	fmt.Println(stu.Name,stu.GetScore())
}