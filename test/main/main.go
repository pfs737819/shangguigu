package main
import (
	"fmt"
	"go_code/test/service"
)

func main() {
	
	rs := service.NewStudent()
	fmt.Println(rs.QueryInfo())

}