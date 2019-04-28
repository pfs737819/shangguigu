package main
import (
	"fmt"
	"io/ioutil"
)

func main() {
	readFilePath := "E:/goproject/src/source/test2.txt"
	writeFilePath := "E:/goproject/src/source/test4.txt"
	content,err := ioutil.ReadFile(readFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	err2 := ioutil.WriteFile(writeFilePath,content,0666)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	
}