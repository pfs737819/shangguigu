package main
import (
	// "fmt"
	"go_code/FamilyBudget/model"
)


func main() {

	rs := model.NewFamilyBudget()
	rs.Menu()

	// var key uint //用户选项

	// var loop bool = true //系统状态

	// var detail string = "家庭收支\t账户金额\t收支金额\t说明" //收入明细

	// balance := 0.00 //账户余额

	// money := 0.00

	// note := "" //收支记录

	// flag := false


	// for{
	// 	fmt.Println("\n\n---------------------------家庭收支系统----------------------------")
	// 	fmt.Println("                           1、收支明细                    ")
	// 	fmt.Println("                           2、登记收入                    ")
	// 	fmt.Println("                           3、登记支出                    ")
	// 	fmt.Println("                           4、退出系统                    ")
	// 	fmt.Println("------------------------------------------------------------------")

	// 	fmt.Println("请选择1-4")
	// 	fmt.Scanln(&key)

	// 	switch key {
	// 		case 1:
	// 			fmt.Println("\n--------------------------- 收支明细 ---------------------------\n")
	// 			if flag {
	// 				fmt.Println(detail)
	// 			}else{
	// 				fmt.Println("当前没有收支记录！来一笔吧~")
	// 			}
	// 			fmt.Println("\n--------------------------- 收支明细 ---------------------------\n")
	// 		case 2:	
	// 			fmt.Println("\n--------------------------- 登记收入 ---------------------------\n")
	// 			fmt.Print("请输入收入金额：")
	// 			fmt.Scanln(&money)
	// 			flag = true
	// 			balance += money
	// 			fmt.Print("请输入收入说明：")
	// 			fmt.Scanln(&note)
	// 			detail += fmt.Sprintf("\n收入\t\t%v\t\t%v\t\t%v",balance,money,note)

	// 			fmt.Println("\n--------------------------- 登记收入 ---------------------------\n")
	// 		case 3:
	// 			fmt.Println("\n--------------------------- 登记支出 ---------------------------\n")
	// 			fmt.Print("请输入支出金额：")
	// 			fmt.Scanln(&money)
	// 			if balance - money >= 0 {
	// 				flag = true
	// 				balance -= money
	// 				fmt.Print("请输入支出说明：")
	// 				fmt.Scanln(&note)
	// 				detail += fmt.Sprintf("\n支出\t\t%v\t\t-%v\t\t%v",balance,money,note)
	// 			}else{
	// 				fmt.Print("\n\n账户余额不足..")
	// 				break;
	// 			}
				
	// 			fmt.Println("\n--------------------------- 登记支出 ---------------------------\n")
	// 		case 4:	
	// 			fmt.Println("你确认要退出吗?输入y/n...")
	// 			choice := " "
	// 			for{
	// 				fmt.Scanln(&choice)
	// 				if choice == "y" || choice == "n" {
	// 					break
	// 				}
	// 				fmt.Println("输入有误! 请输入y/n...")
	// 			}
	// 			if choice == "y"{
	// 				loop = false
	// 			}
	// 		default :
	// 			fmt.Println("请输入正确的选项...")
	// 	}

	// 	if !loop{
	// 		break
	// 	}

	// }
	// fmt.Println("你已退出系统...")

}
