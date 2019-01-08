package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	serAddr, err := net.ResolveTCPAddr("tcp4", ":9527")
	CheckError06S(err)

	// 获取监听对象
	listenner, err := net.ListenTCP("tcp", serAddr)
	CheckError06S(err)

	// 监听等待连接
	fmt.Println("服务端准备就绪，等待客户端连入...")
	conn, err := listenner.Accept()
	CheckError06S(err)
	defer conn.Close()

	for {
		//数据交互
		bs := make([]byte, 1024)
		n, err := conn.Read(bs)
		CheckError06S(err)
		content := string(bs[:n])
		if content == "over" {
			fmt.Println("客户端:", conn.RemoteAddr(), "终止了聊天...")
			break
		}
		fmt.Println("客户端", conn.RemoteAddr(), "对服务端说：", content)

		n, err = conn.Write([]byte("ECHO:" + content))
	}
}

func CheckError06S(err error)  {
	if err != nil {
		os.Exit(1)
	}
}
