package main

import "fmt"

func main() {
	/*1. 定义一个函数，用于求5的阶乘。然后进行调用
	2. 定义一个函数，打印三角形*，然后进行调用。*/
	jieCheng0301()
	prtAngle0301()
}

//1. 定义一个函数，用于求5的阶乘。然后进行调用
func jieCheng0301()  {
	res := 1
	for i := 1; i <= 5; i++ {
		res *= i
	}
	fmt.Println(res)
}

//2. 定义一个函数，打印三角形*，然后进行调用。
func prtAngle0301()  {
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
