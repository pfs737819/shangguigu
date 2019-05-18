package main
import (
	"fmt"
)

func login(userName string, userPwd string) (err error) {
	fmt.Printf("用户名为%s 用户名密码为%s \n",userName,userPwd)
	return nil
}
