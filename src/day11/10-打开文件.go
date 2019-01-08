package main

import (
	"os"
	"fmt"
)

func main() {
	openFile10("./day11/test.txt")
}

func openFile10(fileName string)  {
	file, err := os.Open(fileName)
	defer closeFile10(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("文件打开成功")
}

func closeFile10(file *os.File)  {
	err := file.Close()
	if err != nil {
		fmt.Println("文件关闭失败...")
	}
	fmt.Println("文件关闭成功")
}