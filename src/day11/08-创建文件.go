package main

import (
	"os"
	"fmt"
)

func main() {
	creFile08("d:/demo.txt")
}

func creFile08(fileName string)  {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("文件创建成功", file)
}
