package main
import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		// handle error
	}
	for {
		fmt.Println("等待连接.......")
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}else{
			fmt.Println("连接成功.......",conn.RemoteAddr().Network(),conn.RemoteAddr().String())
		}
		//go handleConnection(conn)
	}
}