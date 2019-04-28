package main
import (
	"fmt"
	"io"
	"os"
	"bufio"
)

func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {

	srcFile,err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("打开文件出错",err)
	}
	defer srcFile.Close()

	reader := bufio.NewReader(srcFile)

	dstFile,err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("OpenFile执行错误",err)
	}
	
	writer := bufio.NewWriter(dstFile)

	return io.Copy(writer,reader)

}

func main() {
	
	dstFile := "E:/goproject/src/source/gntmnew.mp4"
	srcFile := "E:/goproject/src/source/gntm.mp4"
	written,err := CopyFile(dstFile,srcFile)
	if err == nil {
		fmt.Println("拷贝成功",written)
	}else{
		fmt.Println("拷贝失败",err)
	}

}