package main

import "fmt"

var num1102 = 100

func main() {
	/*
		作用域：作用的范围
		一般以{}划分
	*/
	//num1101 := 1
	//fn1101(num1101)

	if i := 0; i < 5 {
		fmt.Println(i)
	} else {
		fmt.Println(i)
	}

	//fmt.Println(i)

	/*for i := 0; i < 5; i++ {
		fmt.Println("xunlun")
	}
	fmt.Println(i)*/

	fmt.Println(num1102)
	fn1102()
	fn1103()
	fn1104()

	//fnc := fn()
}

func fn1101(num int)  {
	fmt.Println(num)
}

func fn1102()  {
	num1102 = 1000
}

func fn1103()  {
	fmt.Println(num1102)
}

func fn1104()  {
	num1103 := 5
	if true {
		fmt.Println(num1103)
		{
			fmt.Println(num1103)
		}
	}
}

func fn() func(a, b int) int {
	c := 2
	return func(a, b int) int {
		c++
		return c
	}
}


