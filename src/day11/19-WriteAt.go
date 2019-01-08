package main

import (
	"os"
	"fmt"
)

func main() {
	file, _ := os.OpenFile("./day11/demo.txt", os.O_WRONLY|os.O_CREATE, 0777)
	defer file.Close()

	//bs := []byte{'X', 'U', 'N', 'L'} // xun论
	str := "wel"
	n, _ := file.WriteAt([]byte(str), 2)
	fmt.Println(n)
}
