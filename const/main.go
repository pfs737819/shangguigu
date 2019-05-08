package main
import (
	"fmt"
)

func main() {
	const (
		num = iota
		a = iota
		b,t = iota,iota
		c = iota
		d = 9/3
	)
	fmt.Println(num,a,b,t,c,d)
}