package main

import "fmt"

func main() {
	//函数声明不会执行，只有调用才会被执行，当函数调用是会到函数声明的地方取执行函数，执行完成之后再次回到函数调用的地方继续往下执行，如果有返回值，返回的结果也会被返回到函数调用的地方
	num1 := 1
	num2 := 2
	res1 := getSum0601(num1, num2)
	fmt.Println(res1)

	res2 := fn0601(num1, num2)
	fmt.Printf("%T\n", res2)
	fmt.Println(res2)
	r1 := res2(1, 3)
	fmt.Printf("%T\n", r1)
	fmt.Println(r1)

	fmt.Println("++++++++++++++++")

	fn1 := fn0602()
	fmt.Printf("%T\n", fn1)
	res3 := fn1()
	fmt.Println(res3)
	res4 := fn1()// 是一个闭包，程序不能被销毁，无法释放内存
	fmt.Println(res4)
}

func getSum0601(a, b int) int { // func (int, int) int
	return a + b
}

func fn0601(a, b int) func(m, n int) int { // func (int, int) int
	return getSum0601 // int
}

func fn0602() func() int {
	m := 1
	fn := func() int {
		m++
		return m
	}
	return fn
	//return fn0603
}

func fn0603 () int  {
	return 1
}
