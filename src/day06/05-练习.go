package main

import "fmt"

func main() {
	//1.定义一个方法：求n的阶乘。调用的时候，由键盘输入。
	jieCheng0501()

	//2.定义一个方法：求2个数的和。2个数由参数传入。。
	num0501 := 1
	num0502 := 2
	getSum0501(num0501, num0502)
}

//1.定义一个方法：求n的阶乘。调用的时候，由键盘输入。
func jieCheng0501()  {
	var n int
	fmt.Println("请输入一个数字，用于求该数字的阶乘:")
	fmt.Scanln(&n)

	res := 1

	for i := 1; i <= n; i++ {
		res *= i
	}

	fmt.Printf("%d的阶乘是：%d\n", n, res)
}

//2.定义一个方法：求2个数的和。2个数由参数传入。。
func getSum0501(num1 int, num2 int)  {
	res := num1 + num2
	fmt.Printf("%d + %d = %d\n", num1, num2, res)
}
