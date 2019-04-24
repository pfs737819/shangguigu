package customer
import (
	"fmt"
)

type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

//使用工厂模式返回一个Customer实例
func NewCustomer(Id int,Name string,Gender string,Age int,Phone string,
	Email string)(Customer){
	return Customer{
		Id : Id,
		Name : Name,
		Gender : Gender,
		Age : Age,
		Phone : Phone,
		Email : Email,
	}
}

//使用工厂模式返回一个Customer实例
func AddCustomer(Name string,Gender string,Age int,Phone string,
	Email string)(Customer){
	return Customer{
		Name : Name,
		Gender : Gender,
		Age : Age,
		Phone : Phone,
		Email : Email,
	}
}


//返回用户信息
func (this Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v",this.Id,this.Name,this.Gender,
	this.Age,this.Phone,this.Email)
	return info
}