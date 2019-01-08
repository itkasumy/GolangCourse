package main

import (
	"os"
	"fmt"
)

func main() {
	mkDir07("./day11/test")

	mkDirAll07("./day11/aa/bb/cc")
}

func mkDir07(dirName string)  {
	err := os.Mkdir(dirName, 0666)
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	fmt.Println("文件创建成功")
}

func mkDirAll07(dirNameList string)  {
	err := os.MkdirAll(dirNameList, 0666)
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	fmt.Println("文件创建成功")
}
