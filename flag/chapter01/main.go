package main
import (
	"fmt"
	"flag"
)

func main() {

	var name string
	var pwd string
	var host string
	var port string

	flag.StringVar(&name,"u","","用户名默认为空")
	flag.StringVar(&pwd,"pwd","","密码默认为空")
	flag.StringVar(&host,"host","localhost","主机默认为localhost")
	flag.StringVar(&port,"port","3306","端口默认为空3306")

	flag.Parse()

	fmt.Printf("name %v,pwd %v,host %v,port %v",name,pwd,host,port)
}