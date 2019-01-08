package main

import (
	"path/filepath"
	"fmt"
)

func main() {
	path, err := filepath.Abs("./day11/test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(path)
	fmt.Printf("%T\n", path)

	b01 := filepath.IsAbs("C:\\Users\\Administrator\\Desktop\\Golang\\src\\day11\\test.txt")
	fmt.Println(b01)

	b02 := filepath.IsAbs(path)
	fmt.Println(b02)
}
