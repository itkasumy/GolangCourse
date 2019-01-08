package main

import (
	"os"
	"fmt"
	"io"
)

func main() {
	/*file, _ := os.Open("./day11/test.txt")
	defer file.Close()

	bs := make([]byte, 1024 * 4)

	count, _ := file.Read(bs)

	file2, _ := os.OpenFile("./day11/demo.txt", os.O_WRONLY|os.O_CREATE, 0666)
	defer file2.Close()
	count2, _ := file2.Write(bs[:count])
	fmt.Println(count2)*/

	copyFile17("./day11/ylj.jpg", "./test.jpg")
}

func copyFile17(srcFileName, dstFileName string) (int, error) {
	srcFile, err := os.Open(srcFileName)
	defer srcFile.Close()
	if err != nil {
		return 0, err
	}

	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	defer dstFile.Close()
	if err != nil {
		return 0, err
	}

	total := 0
	count := 0
	for {
		bs := make([]byte, 1024 * 4)

		count, err = srcFile.Read(bs)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件已经读取完毕！")
				break
			}
			return 0, err
		}

		count, err = dstFile.Write(bs[:count])
		if err != nil {
			return 0, err
		}
		total += count
	}
	fmt.Println("文件写入完毕")
	return total, nil
}
