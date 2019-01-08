package main

import (
	"os"
	"fmt"
)

func main() {
	//返回一个包含文件信息的接口对象
	fileINfo, err := os.Stat("./day11/test.txt")
	if err != nil {
		fmt.Println(err)
	}
	/*fmt.Printf("%T\n", fileINfo)
	fmt.Println(fileINfo)*/

	// 文件的名字
	name := fileINfo.Name()
	fmt.Println(name)
	//文件的大小
	size := fileINfo.Size()
	fmt.Println(size)
	//文件模式
	mod := fileINfo.Mode()
	fmt.Println(mod)
	//文件的修改时间
	modTime := fileINfo.ModTime()
	fmt.Println(modTime)
	//判断是否是一个目录
	b01 := fileINfo.IsDir()
	fmt.Println(b01)
}
