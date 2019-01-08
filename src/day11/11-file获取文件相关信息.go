package main

import (
	"os"
	"fmt"
)

func main() {
	openFile11("./day11/test.txt")

}

func openFile11(fileName string) {
	file, err := os.Open(fileName)
	defer closeFile11(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("文件打开成功")

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Mode())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.ModTime())
	fmt.Println(fileInfo.IsDir())
}

func closeFile11(file *os.File)  {
	err := file.Close()
	if err != nil {
		fmt.Println("文件关闭失败...")
	}
	fmt.Println("文件关闭成功")
}
