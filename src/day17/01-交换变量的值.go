package main

import "fmt"

func main() {
	//1. 交换2个变量的值。（5‘’）
	a := 2
	b := 3
	fmt.Printf("交换之前a = %d, b = %d\n", a, b)
	a, b = b, a
	fmt.Printf("交换之后a = %d, b = %d\n", a, b)
}