package main

import "fmt"

func main() {
	//2. 使用switch完成：模拟计算器操作：输入第一个整数，输入第二个整数，输入运算符(+-*/),打印结果（5‘’）

	var a int
	var b int
	var res int
	var opt string

	fmt.Println("请输入第一个数（a）：")
	fmt.Scanln(&a)
	fmt.Println("请输入第二个数（b）：")
	fmt.Scanln(&b)
	fmt.Println("请输入一个操作符（+、-、*、/）：")
	fmt.Scanln(&opt)

	switch opt {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	}

	fmt.Printf("%d %s %d = %d", a, opt, b, res)
}
