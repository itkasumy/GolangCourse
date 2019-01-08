package main

import "fmt"

func main() {
	// 函数调用时传入实参必须与函数声明时形参的个数严格匹配（类型、个数、顺序）(...interface{})
	var num0601, num0602, num0603 int
	fmt.Println(num0601, num0602, num0603)

	getSum0601(1, 2, "")
}

func getSum0601(a, b int, str string)  {
	sum := a + b
	fmt.Println(sum)
	fmt.Println(str)
}
