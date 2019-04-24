package model

type Person struct{
	Name string
	age  int
	salary float64
}

func NewPerson() Person{
	return Person{
		Name : "张三",
		age : 18,
		salary : 2000,

	}
}

func (p *Person)QueryInfo() string {
	str := "test"
	return str
}

func (p *Person)SetAge(age int){
	
}