package main

import (
	"net"
	"fmt"
	"os"
	"time"
	"strconv"
)

func main() {
	//10. 客户端向服务器传递一个文件
	sAddr, err := net.ResolveTCPAddr("tcp4", "localhost:77677")
	checkErrS(err)

	listenner, err := net.ListenTCP("tcp", sAddr)
	checkErrS(err)

	for {
		conn, err := listenner.Accept()
		checkErrS(err)

		go func() {
			file, err := os.OpenFile(".\\day18\\images\\swan" + strconv.FormatInt(time.Now().UnixNano(), 10) + ".jpg", os.O_WRONLY|os.O_CREATE, 0777)
			checkErrS(err)
			defer file.Close()

			bs := make([]byte, 1024)

			for {
				n, err := conn.Read(bs)
				if n == 0 || err != nil {
					fmt.Println("文件读取完毕...")
					break
				}

				file.Write(bs[:n])
			}
			conn.Close()
		}()
	}
}

func checkErrS(err error) {
	if err != nil {
		fmt.Println("程序出现异常：", err.Error())
		os.Exit(1)
	}
}

