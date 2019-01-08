package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	serAddr, err := net.ResolveTCPAddr("tcp4", ":9527")
	CheckError07S(err)

	// 获取监听对象
	listenner, err := net.ListenTCP("tcp", serAddr)
	CheckError07S(err)

	// 监听等待连接
	fmt.Println("服务端准备就绪，等待客户端连入...")
	for {
		conn, err := listenner.Accept()
		CheckError07S(err)

		go func() {
			for {
				//数据交互
				bs := make([]byte, 1024)
				n, err := conn.Read(bs)
				CheckError07S(err)
				content := string(bs[:n])
				if content == "over" {
					fmt.Println("客户端:", conn.RemoteAddr(), "终止了聊天...")
					break
				}
				fmt.Println("客户端", conn.RemoteAddr(), "对服务端说：", content)

			}
			defer conn.Close()
		}()
	}
}

func CheckError07S(err error)  {
	if err != nil {
		os.Exit(1)
	}
}
