package main

import (
	"os"
	"fmt"
	"io"
)

func main() {
	//1. 打开文件
	/*file, err := os.Open("./day11/test.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	//2. 创建一个byte切片
	bs := make([]byte, 4)*/

	//3. 读取文件
	//第一次读取
	/*count, err := file.Read(bs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(count)
	fmt.Println(string(bs))

	// 第二次读取
	count, err = file.Read(bs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(count)
	fmt.Println(string(bs))

	//第三次读取
	count, err = file.Read(bs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(count)
	fmt.Println(string(bs))

	//第四次读取
	count, err = file.Read(bs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(count)
	fmt.Println(string(bs))*/

	/*for {
		count, err := file.Read(bs)
		if err != nil {
			fmt.Println()
			if err == io.EOF {
				fmt.Println("文件读取完毕")
			}
			fmt.Println(err)
			return
		}
		fmt.Print(string(bs[:count]))
	}*/

	readFile("./day11/test.txt")
}

func readFile(fileName string)  {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	bs := make([]byte, 4)
	for {
		count, err := file.Read(bs)
		if err != nil {
			fmt.Println()
			if err == io.EOF {
				fmt.Println("文件读取完毕")
			}
			fmt.Println(err)
			return
		}
		fmt.Print(string(bs[:count]))
	}
}
