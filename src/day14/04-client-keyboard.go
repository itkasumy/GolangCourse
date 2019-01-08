package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	//获取服务器地址
	serAddr, err := net.ResolveTCPAddr("tcp4", "172.168.30.3:9527")
	CheckError02(err)

	//连接服务器
	tcpConn, err := net.DialTCP("tcp", nil, serAddr)
	CheckError02(err)
	defer tcpConn.Close()

	//数据交互
	info := ""
	fmt.Scanln(&info)
	n, err := tcpConn.Write([]byte(info))
	CheckError02(err)
	fmt.Println("传输到服务器的数据量是：", n)

	bs := make([]byte, 1024)
	n, err = tcpConn.Read(bs)
	CheckError02(err)
	fmt.Println("服务端对客户端说：", string(bs[:n]))

}

func CheckError02(err error)  {
	if err != nil {
		os.Exit(1)
	}
}
