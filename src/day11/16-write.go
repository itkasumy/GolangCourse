package main

import (
	"os"
	"fmt"
)

func main() {
	/*file, err := os.Open("./day11/test.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}*/

	file, err := os.OpenFile("./day11/test.txt", os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	bs := []byte{'A', 'B', 'C', 'D'}

	count, err := file.Write(bs)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("写入数据成功", count)
}
