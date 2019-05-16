package main
import (
	"net"
	"fmt"
	"bufio"
	"os"
)

func readContent(conn net.Conn) {
	defer conn.Close()
	for {
		content := make([]byte,1024)
		n, err := conn.Read(content)
		if err != nil {
			fmt.Println("conn.Read err",err)
			return 
		}
		fmt.Printf("\n %v",string(content[:n]))
	}
}

func main() {
	conn, err := net.Dial("tcp", "192.168.14.5:8080")
	if err != nil {
		fmt.Println("与服务器链接失败 err",err)
	}
	content := os.Stdin
	buf := bufio.NewReader(content)
	for {
		ln, readErr := buf.ReadString('\n')
		if  readErr != nil {
			fmt.Println("buf.ReadString err ",readErr)
		}
		fmt.Println("ln",ln)
		n, err := conn.Write([]byte(ln))
		if err != nil {
			fmt.Println("conn.Write err ",err)
		}
		fmt.Printf("写入字节 %d\n",n)
	}
	go readContent(conn)
	//fmt.Println(conn, "GET / HTTP/1.0\r\n\r\n")
	//status, err := bufio.NewReader(conn).ReadString('\n')
// ...
}