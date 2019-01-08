package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	serAddr, err := net.ResolveTCPAddr("tcp4", ":9527")
	CheckError05S(err)

	// 获取监听对象
	listenner, err := net.ListenTCP("tcp", serAddr)
	CheckError05S(err)

	// 监听等待连接
	fmt.Println("服务端准备就绪，等待客户端连入...")
	conn, err := listenner.Accept()
	CheckError05S(err)
	defer conn.Close()

	for {
		//数据交互
		bs := make([]byte, 1024)
		n, err := conn.Read(bs)
		CheckError05S(err)
		content := string(bs[:n])
		if content == "over" {
			fmt.Println("客户端:", conn.RemoteAddr(), "终止了聊天...")
			break
		}
		fmt.Println("客户端对服务端说：", content)

		info := ""
		fmt.Scanln(&info)
		n, err = conn.Write([]byte(info))
		//fmt.Println("服务端向客户端发送数据，数据量是：", n)
	}
}

func CheckError05S(err error)  {
	if err != nil {
		os.Exit(1)
	}
}
