package main

import "fmt"

func main() {
	/*函数的参数：
		如果是值类型的传递的是值，做的是拷贝
		如果是引用类型的，传递是地址*/
	num1201 := 12
	fmt.Println("函数执行前num =", num1201)
	fn1201(num1201)
	fmt.Println("函数执行后num =", num1201)

	fmt.Println("-----------------------------------")

	arr1201 := [4]int{1, 2, 3, 4}
	fmt.Println("函数执行前", arr1201)
	fn1202(arr1201)
	fmt.Println("函数执行后", arr1201)

	fmt.Println("-----------------------------------")

	s1201 := []int{1, 2, 3, 4}
	fmt.Println("函数执行前", s1201)
	fn1203(s1201)
	fmt.Println("函数执行后", s1201)

	//map[] 引用类型
}

func fn1201(num int)  {
	fmt.Println("=======", num)
	num = 100
	fmt.Println("函数执行中num=", num)
}

func fn1202(arr [4]int)  {
	fmt.Println("=============", arr)
	arr[0] = 100
	fmt.Println("函数执行中arr = ", arr)
}

func fn1203(s []int)  {
	s[0] = 100
	fmt.Println("函数执行中", s)
}
