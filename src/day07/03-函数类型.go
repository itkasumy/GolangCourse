package main

import "fmt"

func main() {
	/*
		基本类型：数值型(浮点型、整型、复数型)、bool、string
		复合类型：数组、切片、map
		函数的类型：引用类型，去除所有的函数名，参数名，返回值名加上 func
		可以将函数赋值给一个声明为函数类型的变量，该变量指向该函数的内存地址，该变量可以像函数一样调用
	*/
	num1 := 1
	num2 := 2
	num3 := getSum0301(num1, num2)
	fmt.Printf("%T, %T, %T\n", num1, num2, num3)

	fmt.Printf("%T\n", getSum0301) // func(int, int) int
	//函数也是一种数据类型

	var num4 int
	var arr1 [4]int
	//格式是： var variable_name variable_type
	var fn1 func(int, int) int
	num4 = 2
	arr1 = [4]int{1, 3, 4, 6}
	fn1 = getSum0301

	fmt.Println(num4, arr1, fn1)
	fmt.Println(getSum0301) // 引用类型

	res1 := fn1(num1, num2)
	fmt.Println(res1)

	fmt.Printf("%T\n", increment0301)
	fmt.Printf("%T\n", getSum0302) // func (...int) int
	fmt.Printf("%T\n", getRect0301) // func (float64, float64) (float64, float64)
	fmt.Printf("%T\n", fn0301) // func (string)

	var fn2 func (...int) int
	//fn2 = getSum0302(1, 2) // 0 int
	fn2 = getSum0302
	fmt.Printf("%T\n", fn2)
	fmt.Printf("%p, %p\n", getSum0302, fn2)
	fmt.Println(fn2(1, 2, 3, 4))

}

func getSum0301(a, b int) int {
	return a + b
}

func increment0301(a int) (res int) {
	res = a + 1
	return res
}

func getSum0302(nums... int) int {
	return 0
}

func getRect0301(l float64, w float64) (area float64, zc float64) {
	return 1.2, 2.3
}

func fn0301(str string)  {

}
