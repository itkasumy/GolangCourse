package main

import (
	"os"
	"io"
	"fmt"
)

func main() {
	srcName := "./day12/ubuntu.iso"
	dstName := "ubuntu.iso"
	total, err := copyFile01(srcName, dstName)
	if err != nil {
		if err == io.EOF {
			fmt.Println("文件写入完毕，大小是：", total)
		}
		fmt.Println(err)
	}
}

func copyFile01(srcFileName, dstFileName string) (total int64, err error) {
	srcFile, err := os.Open(srcFileName)
	defer srcFile.Close()
	if err != nil {
		return 0, err
	}

	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0764)
	defer dstFile.Close()
	if err != nil {
		return 0, err
	}

	const BS_SIZE = 1024 * 4

	for {
		bs := make([]byte, BS_SIZE)
		count, err := srcFile.Read(bs)
		if err != nil {
			if err == io.EOF {
				return total, err
			}
			return 0, err
		}

		count, err = dstFile.Write(bs)
		if err != nil {
			return 0, err
		}
		total += int64(count)
	}
}