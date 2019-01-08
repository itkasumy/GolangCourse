package main

import (
	"net"
	"fmt"
)

func main() {
	//第一步：获取tcp地址
	serverAddr := "172.168.30.3:6666"

	TCPAddr, err := net.ResolveTCPAddr("tcp4", serverAddr)
	fmt.Println(TCPAddr, err)
	fmt.Printf("%T\n", TCPAddr)

	//第二步，监听该端口
	listenner, err := net.ListenTCP("tcp", TCPAddr)
	fmt.Println(listenner, err)
	fmt.Printf("%T\n", listenner)

	fmt.Println("服务器准备就绪,等待客户端连接...")
	//第三步，等待客户端呼叫,阻塞
	conn, err := listenner.Accept()
	fmt.Println(conn, err)
	fmt.Printf("%T\n", conn)
	defer conn.Close()

	//第四步，读取客户端发送过来的数据
	bs := make([]byte, 1024)
	n, err := conn.Read(bs)
	fmt.Println(n, err)
	fmt.Println(string(bs[:n]))

	//向客户端发送数据
	conn.Write([]byte("撒哇你卡..."))
}
