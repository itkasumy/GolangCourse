package main

import (
	"os"
	"fmt"
)

func main() {
	remFile09("./day11/test")

	remFilDir09("./day11/aa/bb")
}

func remFile09(fileName string)  {
	err := os.Remove(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("文件/目录删除成功！")
}

func remFilDir09(fileName string)  {
	err := os.RemoveAll(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("全部删除成功")
}
