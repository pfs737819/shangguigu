package model
import (
	"fmt"
)


type Student struct {
	Name string
	Age  int
	Score float64
}

type Pupil struct {
	Student
}

func (p *Student) ShowInfo() {
	fmt.Printf("学生姓名 : %v 年龄 : %v 成绩 : %v \n",p.Name,p.Age,p.Score)
}

func (p *Student) SetScore(Score float64) {
	p.Score = 68
}

func (p *Student) Testing() {
	fmt.Println("学生正在考试")
}


type Graduate struct {
	Student
}
