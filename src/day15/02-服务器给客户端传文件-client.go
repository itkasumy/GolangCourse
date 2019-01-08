package main

import (
	"net"
	"fmt"
	"os"
	"time"
	"strconv"
)

func main() {
	// step1: 提供一个访问的服务器的地址
	serAddr, err := net.ResolveTCPAddr("tcp4", "172.168.30.3:9999")
	checkErr02C(err)

	// step2: 连接服务器
	tcpConn, err := net.DialTCP("tcp", nil, serAddr)
	checkErr02C(err)

	//step3: 数据交互
	file, err := os.OpenFile(".\\day15\\images\\hd" + strconv.FormatInt(time.Now().UnixNano(), 10) + ".jpeg", os.O_WRONLY|os.O_CREATE, 0777)
	checkErr02C(err)
	bs := make([]byte, 1024)

	for {
		n, err := tcpConn.Read(bs)
		if err != nil || n == 0 {
			fmt.Println("客户端读取数据结束...")
			break
		}
		file.Write(bs[:n])
	}

	//step4: 关闭文件
	file.Close()
	tcpConn.Close()
}

func checkErr02C(err error)  {
	if err != nil {
		fmt.Println("程序运行中出现了错误:", err.Error())
		os.Exit(1)
	}
}
