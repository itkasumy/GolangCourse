package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	//step1: 提供一个服务器地址
	serAddr, err := net.ResolveTCPAddr("tcp4", ":9999")
	checkErr01S(err)

	//step2: 创建一个监听对象
	listenner, err := net.ListenTCP("tcp", serAddr)
	checkErr01S(err)

	//step3: 监听对象监听端口，等待客户端连接
	conn, err := listenner.Accept()
	checkErr01S(err)

	//step4: 数据交互
	file, err := os.OpenFile(".\\day15\\images\\hd.jpeg", os.O_CREATE|os.O_WRONLY, 0777)
	checkErr01S(err)
	bs := make([]byte, 1024)

	for {
		n, err := conn.Read(bs)
		if n == 0 || err != nil {
			fmt.Println("文件读取完毕...")
			break
		}
		file.Write(bs[:n])
	}

	//step5: 关闭文件
	file.Close()
	conn.Close()
}

func checkErr01S(err error)  {
	if err != nil {
		fmt.Println("程序运行中出现了错误:", err.Error())
		os.Exit(1)
	}
}
