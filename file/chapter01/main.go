package main
import (
	"fmt"
	"os"
	"bufio"
	"io"
)

func main() {
	file,err := os.Open("E:/goproject/src/source/test.txt")
	if err != nil {
		fmt.Println(err)
	}



	defer file.Close()

	reader := bufio.NewReader(file)

	for{
		str,err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}

	
}