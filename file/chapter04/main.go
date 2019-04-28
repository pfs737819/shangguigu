package main
import (
	"fmt"
	"os"
	"bufio"
	"io"
)

func main() {
	readFilePath := "E:/goproject/src/source/test2.txt"
	readFile,err := os.OpenFile(readFilePath,os.O_RDONLY,0666)
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()


	reader := bufio.NewReader(readFile) 
	//读取文件的内容


	writeFilePath := "E:/goproject/src/source/test3.txt"
	writeFile,err := os.OpenFile(writeFilePath,os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_TRUNC,0666)
	defer writeFile.Close()
	writer := bufio.NewWriter(writeFile)
	for {
		str,err := reader.ReadString('\n')
		fmt.Println(str)
		writer.WriteString(fmt.Sprintf("%v\r\n",str))
		if err == io.EOF{
			break
		}
	}
	writer.Flush()

	// str := "hello world"

	// 

	// for i := 0;i < 5;i ++ {
	// 	writer.WriteString(fmt.Sprintf("%v%v\r\n",str,i))
	// }
	// writer.Flush()
}