package main

import (
	"net"
	"os"
	"fmt"
)

func main() {
	//获取服务器地址
	serAddr, err := net.ResolveTCPAddr("tcp4", "172.168.30.3:9527")
	CheckError05C(err)

	//连接服务器
	tcpConn, err := net.DialTCP("tcp", nil, serAddr)
	CheckError05C(err)
	defer tcpConn.Close()

	for {
		//数据交互
		info := ""
		fmt.Scanln(&info)
		n, err := tcpConn.Write([]byte(info))
		CheckError05C(err)
		if info == "over" {
			fmt.Println("我终止了聊天...")
			break
		}
		//fmt.Println("传输到服务器的数据量是：", n)

		bs := make([]byte, 1024)
		n, err = tcpConn.Read(bs)
		CheckError05C(err)
		fmt.Println("服务端对客户端说：", string(bs[:n]))
	}

}

func CheckError05C(err error)  {
	if err != nil {
		os.Exit(1)
	}
}
