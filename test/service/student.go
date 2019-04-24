package service
import (
	"go_code/test/model"
)

type Student struct{
	studentInfo model.Person
}

func NewStudent() model.Person{
	Student := Student{}
	Student.studentInfo = model.NewPerson()
	return Student.studentInfo
}