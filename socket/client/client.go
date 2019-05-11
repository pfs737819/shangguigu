package main
import (
	"net"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.14.5:8080")
	if err != nil {
		// handle error
	}
	fmt.Println(conn, "GET / HTTP/1.0\r\n\r\n")
	//status, err := bufio.NewReader(conn).ReadString('\n')
// ...
}