package main
import (
	"fmt"
	"io/ioutil"
)

func main() {
	file := "E:/goproject/src/source/test.txt"
	content,err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Sprintf("read file err = %v",err)
	}
	fmt.Printf("\n%v",string(content))
}