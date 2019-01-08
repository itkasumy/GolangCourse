package main

import (
	"os"
	"io"
	"fmt"
)

func main() {
	srcName := "./day12/ubuntu.iso"
	dstName := "ubuntu.iso"
	total, err := copyFile02(srcName, dstName)
	if err != nil {
		if err == io.EOF {
			fmt.Println("文件写入完毕，大小是：", total)
		}
		fmt.Println(err)
	}
	fmt.Println("拷贝总字节大小是：", total)
}

func copyFile02(srcFileName, dstFileName string) (int64, error) {
	srcFile, err := os.Open(srcFileName)
	defer srcFile.Close()
	if err != nil {
		return 0, err
	}

	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0777)
	defer dstFile.Close()
	if err != nil {
		return 0, err
	}

	return io.Copy(dstFile, srcFile)
}
