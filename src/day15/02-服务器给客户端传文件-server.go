package main

import (
	"net"
	"os"
	"fmt"
	"io"
)

func main() {
	//step1: 提供一个服务器地址
	serAddr, err := net.ResolveTCPAddr("tcp4", ":9999")
	checkErr02S(err)

	//step2: 创建一个监听对象
	listenner, err := net.ListenTCP("tcp", serAddr)
	checkErr02S(err)

	//step3: 监听对象监听端口，等待客户端连接
	for {
		conn, err := listenner.Accept()
		checkErr02S(err)

		file, err := os.Open(".\\day15\\hd.jpeg")
		checkErr02S(err)

		go aaa(file, conn)
	}
}

func aaa(file *os.File, conn net.Conn) {
	bs := make([]byte, 1024)
	for {
		n, err := file.Read(bs)
		if err != nil {
			if err == io.EOF {
				fmt.Println("服务端读取文件结束...")
				break
			}
			fmt.Println(err)
			break
		}
		conn.Write(bs[:n])
	}
	file.Close()
	conn.Close()
}

func checkErr02S(err error)  {
	if err != nil {
		fmt.Println("程序运行中出现了错误:", err.Error())
		os.Exit(1)
	}
}
