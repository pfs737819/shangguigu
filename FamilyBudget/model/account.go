package model
import (
	"fmt"
)

type FamilyBudget struct {

	key uint //用户选项

	loop bool  //系统状态

	detail string  //收入明细

	balance float64 //账户余额

	money float64 //收支金额

	note string //收支记录

	flag bool //收支状态
}


func NewFamilyBudget() *FamilyBudget{
	return &FamilyBudget{
		key : 0,
		loop : true,
		detail : "家庭收支\t账户金额\t收支金额\t说明",
		balance : 0.00,
		money : 0.00,
		note : "",
		flag : false,
	}
}

func (f *FamilyBudget) Menu(){

	for{
		fmt.Println("\n\n---------------------------家庭收支系统----------------------------")
		fmt.Println("                           1、收支明细                    ")
		fmt.Println("                           2、登记收入                    ")
		fmt.Println("                           3、登记支出                    ")
		fmt.Println("                           4、退出系统                    ")
		fmt.Println("------------------------------------------------------------------")

		fmt.Println("请选择1-4")
		fmt.Scanln(&(f.key))

		switch f.key {
			case 1:
				f.Query()
			case 2:	
				f.Income()
			case 3:
				f.Pay()
			case 4:	
				f.Exit()
			default :
				fmt.Println("请输入正确的选项...")
		}

		if !f.loop{
			break
		}

	}
	fmt.Println("你已退出系统...")
}

func (f *FamilyBudget) Query(){
	fmt.Println("\n--------------------------- 收支明细 ---------------------------\n")
	if f.flag {
		fmt.Println(f.detail)
	}else{
		fmt.Println("当前没有收支记录！来一笔吧~")
	}
	fmt.Println("\n--------------------------- 收支明细 ---------------------------\n")
}

func (f *FamilyBudget) Income(){
	fmt.Println("\n--------------------------- 登记收入 ---------------------------\n")
	fmt.Print("请输入收入金额：")
	fmt.Scanln(&f.money)
	f.flag = true
	f.balance += f.money
	fmt.Print("请输入收入说明：")
	fmt.Scanln(&f.note)
	f.detail += fmt.Sprintf("\n收入\t\t%v\t\t%v\t\t%v",f.balance,f.money,f.note)

	fmt.Println("\n--------------------------- 登记收入 ---------------------------\n")
}

func (f *FamilyBudget) Pay(){
	fmt.Println("\n--------------------------- 登记支出 ---------------------------\n")
	fmt.Print("请输入支出金额：")
	fmt.Scanln(&f.money)
	if f.balance - f.money >= 0 {
		f.flag = true
		f.balance -= f.money
		fmt.Print("请输入支出说明：")
		fmt.Scanln(&f.note)
		f.detail += fmt.Sprintf("\n支出\t\t%v\t\t-%v\t\t%v",f.balance,f.money,f.note)
	}else{
		fmt.Print("\n\n账户余额不足..")
		//break;
	}
	
	fmt.Println("\n--------------------------- 登记支出 ---------------------------\n")
}

func (f *FamilyBudget) Exit(){
	fmt.Println("你确认要退出吗?输入y/n...")
	choice := " "
	for{
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("输入有误! 请输入y/n...")
	}
	if choice == "y"{
		f.loop = false
	}

}

