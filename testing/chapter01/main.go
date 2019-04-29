package main
import (
	"fmt"
	// "encoding/json"
)
func AddUpper(n int) int {

	res := 0

	for i := 0;i <= n;i ++ {
		res += i
	}

	return res

}



func main() {
	rs :=AddUpper(10)
	fmt.Println(rs)
}