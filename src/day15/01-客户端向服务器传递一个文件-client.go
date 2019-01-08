package main

import (
	"net"
	"fmt"
	"os"
	"io"
)

func main() {
	// step1: 提供一个访问的服务器的地址
	serAddr, err := net.ResolveTCPAddr("tcp4", "172.168.30.3:9999")
	checkErr01C(err)

	// step2: 连接服务器
	tcpConn, err := net.DialTCP("tcp", nil, serAddr)
	checkErr01C(err)

	//step3: 数据交互
	file, err := os.Open(".\\day15\\hd.jpeg")
	checkErr01C(err)
	bs := make([]byte, 1024)

	for {
		n, err := file.Read(bs)
		if err != nil {
			if err == io.EOF {
				fmt.Println("已经读取完全部文件...")
				break
			}
			fmt.Println("程序运行中出现了错误:", err.Error())
			os.Exit(1)
		}
		tcpConn.Write(bs[:n])
	}

	//step4: 关闭文件
	file.Close()
	tcpConn.Close()
}

func checkErr01C(err error)  {
	if err != nil {
		fmt.Println("程序运行中出现了错误:", err.Error())
		os.Exit(1)
	}
}
