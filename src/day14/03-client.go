package main

import (
	"net"
	"fmt"
)

func main() {
	//1. 获取要连接的服务器的IP+port
	serverAddr := "172.168.30.3:6666"
	rAddr, err := net.ResolveTCPAddr("tcp4", serverAddr)
	fmt.Println(rAddr, err)
	fmt.Printf("%T\n", rAddr)

	//2. 连接服务器
	tcpConn, err := net.DialTCP("tcp", nil, rAddr)
	fmt.Println(tcpConn, err)
	fmt.Printf("%T\n", tcpConn)
	fmt.Println("客户端已连接上服务器...")
	defer tcpConn.Close()

	//3. 向服务端发送数据
	n, err := tcpConn.Write([]byte("空你齐瓦..."))
	fmt.Println(n, err)

	//接收服务端的数据
	bs := make([]byte, 1024)
	n, err = tcpConn.Read(bs)
	fmt.Println(n, err)
	fmt.Println(string(bs[:n]))
}
