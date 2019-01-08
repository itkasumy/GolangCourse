package main

import (
	"os"
	"fmt"
	"io"
)

func main() {
	file, err := os.Open("./day11/test.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	bytes := make([]byte, 4)

	//off: offset 偏移量
	count, err := file.ReadAt(bytes, 3)
	if err != nil {
		if err == io.EOF {
			fmt.Println("文件读取完毕")
		}
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes[:count]))
}
