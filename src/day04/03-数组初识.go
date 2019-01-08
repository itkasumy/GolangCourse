package main

import "fmt"

func main() {
	//数组是什么：
		//数组是一组具有相同类型的有序的数据结构的集合
		//Go数组的长度是固定的
	//数组的声明：
	//var variable_name [SIZE] variable_type
	//var variable_name = [SIZE] variable_type {}

	var num0302 int
	fmt.Println(num0302)

	var arr0301 [5] int
	fmt.Println(arr0301)
	//相当于找了一个容器，里面存放数据
	//len() 计算数组的长度
	//cap() 计算数值的容量
	fmt.Println(len(arr0301))

	var arr0302 [2] string
	fmt.Println(arr0302, len(arr0302))

	//var variable_name = [SIZE] variable_type {}

	var num0301 int = 3
	//变量的组成：
	//变量名 类型 值 内存地址
	fmt.Printf("%d ,%p\n", num0301, &num0301)
	var arr0303 = [5] int {1, 2, 3, 4, 5}
	fmt.Print(arr0303, " ")
	fmt.Printf("%p\n", &arr0303)
	//数组的访问
	fmt.Println("数组的第1个数据是", arr0303[0])
	fmt.Println("数组的第2个数据是", arr0303[1])
	fmt.Println("数组的第3个数据是", arr0303[2])

	fmt.Printf("%p\n", &arr0303)
	fmt.Printf("%p\n", &arr0303[0])
	arr0303[0] = 100
	fmt.Println(arr0303)

	var arr0304 = [5]int {1, 2, 3}
	fmt.Println(arr0304)
	fmt.Println("len(arr4) =", len(arr0304))
	//fmt.Println(arr0304[5])
	//invalid array index 5 (out of bounds for 5-element array)

	/*
	array index 5 out of bounds [0:5]
	var arr0305 = [5]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr0305)*/

	arr0305 := [4] int {}
	fmt.Println(arr0305)
	arr0306 := [4]int {1, 2, 3, 4}
	fmt.Println(arr0306)

	var num0303 int = 3
	fmt.Println(num0303)
	num0303 = 4
	fmt.Println(num0303)

	arr0306[0] = 100
	fmt.Println(arr0306)
}
