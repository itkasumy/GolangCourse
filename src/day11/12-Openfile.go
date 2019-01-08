package main

import (
	"os"
	"fmt"
)

func main() {
	file, err := os.OpenFile("./day11/test.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("文件打开成功！")
	file.Close()
}
