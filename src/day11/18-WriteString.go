package main

import (
	"os"
	"fmt"
)

func main() {
	file, _ := os.OpenFile("./day11/demo.txt", os.O_WRONLY|os.O_CREATE, 0777)
	defer file.Close()
	n, _ := file.WriteString("you are welcome")
	fmt.Println(n)
}
