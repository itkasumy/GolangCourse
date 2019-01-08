package main

import (
	"os"
	"fmt"
)

func main() {
	/*seek:
	func (*File) Seek
	func (f *File) Seek(offset int64, whence int) (ret int64, err error)
	Seek设置下一次读/写的位置。offset为相对偏移量，而whence决定相对位置：0为相对文件开头，1为相对当前位置，2为相对文件结尾。它返回新的偏移量（相对开头）和可能的错误。*/

	file, err := os.Open("./day11/test.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	bs := make([]byte, 4)

	count, err := file.Read(bs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs[:count]))

	ret, err := file.Seek(2, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	count, err = file.Read(bs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs[:count]))

	ret, err = file.Seek(2, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ret)
}
