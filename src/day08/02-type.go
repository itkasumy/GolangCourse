package main

import "fmt"

type my_int int
type my_fn func(int, int) int

func main() {
	//type: 定义类型，用于取别名 // alias
	var num1 int
	num1 = 1
	fmt.Println(num1)
	fmt.Printf("%T\n", num1)

	var num2 my_int
	num2 = 2
	fmt.Println(num2)
	//var num3 my_int = "xunlun"
	fmt.Printf("%T\n", num2)

	//num2 = num1 // cannot use num1 (type int) as type my_int in assignment
	//num1 = num2
}

func fn0201(a, b int, fn func(int, int) int) int {
	return 0
}

func fn0202(a, b int, fn my_fn) int {
	return 0
}
