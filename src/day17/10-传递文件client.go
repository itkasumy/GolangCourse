package main

import (
	"net"
	"fmt"
	"os"
	"io"
)

func main() {
	//10. 客户端向服务器传递一个文件
	sAddr, err := net.ResolveTCPAddr("tcp4", "172.168.30.3:6666")
	checkErrC(err)

	tcpConn, err := net.DialTCP("tcp", nil, sAddr)
	checkErrC(err)
	defer tcpConn.Close()

	file, err := os.Open(".\\day18\\swan.jpg")
	defer file.Close()
	checkErrC(err)

	bs := make([]byte, 1024)
	for {
		n, err := file.Read(bs)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件读取完毕！")
				break
			}
			fmt.Println("文件读取遇到错误：", err.Error())
			break
		}

		_, err = tcpConn.Write(bs[:n])
		if err != nil {
			checkErrC(err)
		}
	}
}

func checkErrC(err error) {
	if err != nil {
		fmt.Println("程序出现异常：", err.Error())
		os.Exit(1)
	}
}
