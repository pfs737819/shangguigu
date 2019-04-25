package service
import (
	// "fmt"
	"go_code/shangguigu/CustomerManagement/customer"
)
type CustomerService struct {
	customers []customer.Customer
	customerNum int
}

func NewCustomerService()*CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	newCustomer := customer.NewCustomer(1,"张三","男",18,"17828005002","654498024@qq.com")
	// fmt.Printf("%p\n\n",newCustomer)
	customerService.customers = append(customerService.customers,newCustomer)
	return customerService
}

//返回客户数据
func(this *CustomerService) List() []customer.Customer {
	return this.customers
	// fmt.Println("\n\n---------------------------家庭收支系统----------------------------")
	// fmt.Println("-------------------------------家庭收支系统------------------------\n\n")
}

//添加客户数据
func(this *CustomerService) Add(customer customer.Customer) bool {

	this.customerNum ++
	customer.Id = this.customerNum
	this.customers = append(this.customers,customer)
	return true
}

//修改客户数据
func(this *CustomerService) Update(id int,customerInfo customer.Customer) bool {

	index,_ := this.FindById(id)
	if index == -1 {
		return false
	}

	this.customers[index].Name = customerInfo.Name
	this.customers[index].Gender = customerInfo.Gender
	this.customers[index].Age = customerInfo.Age
	this.customers[index].Phone = customerInfo.Phone
	this.customers[index].Email = customerInfo.Email
	return true
}

//根据id查找客户数据
func(this *CustomerService) FindById(id int) (int,customer.Customer) {
	index := -1
	var customer customer.Customer
	for i := 0;i < len(this.customers);i ++ {
		if this.customers[i].Id == id {
			index = i
			customer = this.customers[i]
			break
		}
	}
	return index,customer
}

//根据id删除客户数据
func(this *CustomerService) Del(id int) bool {
	index,_ := this.FindById(id)
	if index == -1 {
		return false
	}else{
		this.customers = append(this.customers[:index],this.customers[index+1:]...)
		return true
	}
}