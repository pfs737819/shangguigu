package main
import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	filePath := "E:/goproject/src/source/test2.txt"
	file,err := os.OpenFile(filePath,os.O_TRUNC|os.O_CREATE|os.O_WRONLY,0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	str := "hello world"

	writer := bufio.NewWriter(file)

	for i := 0;i < 5;i ++ {
		writer.WriteString(fmt.Sprintf("%v%v\r\n",str,i))
	}
	writer.Flush()
}