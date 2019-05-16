package main
import (
	"fmt"
	"net"
)

var (
	tcp = make([]net.Conn,0)
)

func handleConnection(conn net.Conn,conn2 net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte,1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err",err)
			return 
		}
		fmt.Printf("\n %v : %v",conn.RemoteAddr().String(),string(buf[:n]))
		conn2.Write([]byte(conn.RemoteAddr().String() + ": hello"))
	}
	
}

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
			tcp = append(tcp,conn)
			fmt.Println(tcp)
		}

		var tcpc net.Conn

		if len(tcp) > 1 {
			tcpc = tcp[1]
		}else{
			tcpc = tcp[0]
		}
		go handleConnection(conn,tcpc)
	}
}