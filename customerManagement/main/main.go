package main
import (
	"fmt"
	"go_code/shangguigu/CustomerManagement/service"
	"go_code/shangguigu/CustomerManagement/customer"
)

type customerView struct {
	key string
	loop bool
	customerService *service.CustomerService
}

//显示客户信息列表
func (this *customerView) list() {

	customers := this.customerService.List()
	
	if len(customers) > 0{
		fmt.Println("\n\n---------------------------客户信息列表----------------------------")
		fmt.Println("\n用户ID\t姓名\t性别\t年龄\t手机号码\t邮箱地址")
		for i := 0;i < len(customers);i ++ {
			fmt.Println(customers[i].GetInfo())
		}
		fmt.Println("\n---------------------------客户信息列表----------------------------")
	}else{
		fmt.Println("\n\n暂无用户信息")
	}
	
}

//添加客户
func (this *customerView) add() {
	fmt.Println("\n\n---------------------------添加客户----------------------------")
	fmt.Println("请输入用户姓名")
	name := ""
	fmt.Scanln(&name)

	fmt.Println("请输入用户性别")
	Gender := ""
	fmt.Scanln(&Gender)

	fmt.Println("请输入用户年龄")
	Age := 0
	fmt.Scanln(&Age)

	fmt.Println("请输入用户手机号码")
	Phone := ""
	fmt.Scanln(&Phone)

	fmt.Println("请输入用户邮箱地址")
	Email := ""
	fmt.Scanln(&Email)

	addCustomer := customer.AddCustomer(name,Gender,Age,Phone,Email)
	addCustomerService := this.customerService.Add(addCustomer)
	if addCustomerService {
		fmt.Println("\n\n添加客户成功")
	}else{
		fmt.Println("\n\n添加客户失败")
	}
}

//修改客户
func (this *customerView) update() {
	fmt.Println("\n\n---------------------------修改客户信息----------------------------")
	fmt.Print("请输入待修改的用户编号(-1:退出)")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}

	index,customerInfo := this.customerService.FindById(id)
	if index == -1 {
		fmt.Print("\n该用户不存在")
		return
	}
	fmt.Printf("姓名(%v) :",customerInfo.Name)
	name := ""
	fmt.Scanln(&name)
	if name != ""{
		customerInfo.Name = name
	}

	fmt.Printf("性别(%v) :",customerInfo.Gender)
	gender := ""
	fmt.Scanln(&gender)
	if gender != ""{
		customerInfo.Gender = gender
	}

	fmt.Printf("年龄(%v) :",customerInfo.Age)
	var age int = -1
	fmt.Scanln(&age)
	if age != -1 {
		customerInfo.Age = age
	}

	fmt.Printf("电话(%v) :",customerInfo.Phone)
	phone := ""
	fmt.Scanln(&phone)
	if phone != "" {
		customerInfo.Phone = phone
	}

	fmt.Printf("邮箱(%v) :",customerInfo.Email)
	email := ""
	fmt.Scanln(&email)
	if email != "" {
		customerInfo.Email = email
	}

	result := this.customerService.Update(id,customerInfo)
	if result {
		fmt.Println("\n\n修改客户成功")
	}else{
		fmt.Println("\n\n修改客户失败")
	}
}


//删除客户
func (this *customerView) del() {

	fmt.Println("\n\n---------------------------删除客户----------------------------")
	fmt.Println("请输入待删除的用户编号(-1:退出)")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}

	fmt.Println("\n\n确认是否删除请输入 y|n")
	choice := ""
	for{
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" || choice == "Y" || choice == "N" {
			break
		}
		fmt.Println("输入有误! 请输入y/n...")
	}

	if choice == "n"{
		return
	}
	
	if this.customerService.Del(id) {
		fmt.Println("\n\n删除客户成功")
	}else{
		fmt.Println("\n\n删除客户失败")
	}

	

	fmt.Println("\n\n---------------------------删除客户----------------------------")

	
}

//退出系统
func (this *customerView) exit() {
	fmt.Println("\n\n确认是否退出 y|n")
	choice := ""
	for{
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" || choice == "Y" || choice == "N" {
			break
		}
		fmt.Println("输入有误! 请输入y/n...")
	}

	if choice == "n" || choice == "N"{
		return
	}
	
	this.loop = false
}

func (this *customerView) mainMenu(){
	for{

		fmt.Println("\n\n---------------------------客户信息管理系统----------------------------")
		fmt.Println("                           1、添加客户                                 ")
		fmt.Println("                           2、修改客户                                 ")
		fmt.Println("                           3、删除客户                                 ")
		fmt.Println("                           4、客户列表                                 ")
		fmt.Println("                           5、退出系统                                 ")
		fmt.Println("-----------------------------------------------------------------------\n\n")
	
		fmt.Println("请选择1-5")
		fmt.Scanln(&(this.key))
		switch this.key {
			case "1":
				this.add()
			case "2":	
				this.update()
			case "3":
				this.del()
			case "4":	
				this.list()
			case "5":	
				this.exit()
			default :
				fmt.Println("请输入正确的选项...")
		 }

		if !this.loop{
			break
		}

	}

	fmt.Println("你已退出系统...")
	
}



func main() {
	customerView := customerView{
		key : "",
		loop : true,
	}
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()
}