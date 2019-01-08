package main

import "fmt"

func main() {
	//使用switch完成：模拟计算器操作：输入第一个整数，输入第二个整数，输入运算符(+-*/),打印结果
	num0301 := 0
	num0302 := 0
	oper0301 := ""
	res0301 := 0

	fmt.Println("请输入一个整数：")
	fmt.Scanln(&num0301)
	fmt.Println("请输入第二个整数：")
	fmt.Scanln(&num0302)
	fmt.Println("请输入一个操作符：")
	fmt.Scanln(&oper0301)

	switch oper0301 {
	case "+":
		res0301 = num0301 + num0302
	case "-":
		res0301 = num0301 - num0302
	case "*":
		res0301 = num0301 * num0302
	case "/":
		res0301 = num0301 / num0302
	}
	fmt.Printf("%d %s %d = %d", num0301, oper0301, num0302, res0301)
}
