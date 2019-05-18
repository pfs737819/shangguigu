package main
import (
	"fmt"
	"os"
)

func main() {
	var loop bool = true //用户输入控制
	var key int //用户选项

	var userName string // 用户名称
	var userPwd string // 用户密码

	//菜单显示
	for loop {
		fmt.Println("----------欢迎登陆多人聊天系统----------")
		fmt.Println("\t\t1.登陆聊天室")
		fmt.Println("\t\t2.用户注册")
		fmt.Println("\t\t3.退出系统")
		fmt.Println("\t\t请选择1-3")

		fmt.Scanln(&key)

		switch key{
			case 1 :
				fmt.Println("登陆聊天室")
				loop = false
			case 2 :
				fmt.Println("用户注册")
				loop = false
			case 3 :
				fmt.Println("退出系统")
				// loop = false
				os.Exit(0)
			default :
				fmt.Println("你的输入有误请重新输入!")	
		}
	}

	if key == 1 {
		fmt.Println("请输入用户名")
		fmt.Scanln(&userName)
		fmt.Println("请输入用户密码")
		fmt.Scanln(&userPwd)
		err := login(userName,userPwd)
		if err != nil {
			fmt.Println("登陆失败")
		} else {
			fmt.Println("登陆成功")
		}
	}



}