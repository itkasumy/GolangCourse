package main

import "fmt"

func main() {
	/*参数：
		函数在执行过程中需要外部提供的数据
		形参：函数定义时所写的参数，只是一个占位的作用
		实参：函数调用时所写的参数，它是函数真正参与运算的参数
		实参是函数调用的时间传入的参数，必须与定义时的参数严格匹配（个数，类型，顺序）
	*/

	//1. 计算1-10的和
	getSum0401(10)
	//2. 计算1-50的和
	getSum0401(50)
	//3. 计算1-100的和
	getSum0401(100)

	getAdd0401(1, 2)
	getAdd0401(2, 1)
}

func getSum0401(n int)  {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Printf("1-%d的和是：%d\n", n, sum)
}

func getAdd0401(a int, b int)  {
	res := a + b
	fmt.Printf("%d + %d = %d\n", a, b, res)
}
