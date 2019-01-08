package main

import "fmt"

func main() {
	/*
		深拷贝：拷贝的是值
			基本类型（数值型、bool、string），数组
		浅拷贝：拷贝的是地址
			切片、map、函数

		可以使用指针将值类型的数据做浅拷贝（地址传递）
	*/
	a := 10
	b := a
	fmt.Printf("改变前：a = %d, b = %d\n", a, b)
	b = 100
	fmt.Printf("改变后：a = %d, b = %d\n", a, b)

	arr1 := [3]int{1, 2, 3}
	fmt.Println("函数执行前：arr1 = ", arr1)
	changeArr0901(arr1)
	fmt.Println("函数执行后：arr1 = ", arr1)

	fmt.Printf("%T\n", arr1)
	var p1 *[3]int = &arr1
	fmt.Printf("%T\n", p1) //  //*[3]int
	fmt.Println(*p1)

	fmt.Println("参数为指针的函数执行前：arr1 = ", arr1)
	chageArr0902(&arr1)
	fmt.Println("参数为指针的函数执行后：arr1 = ", arr1)
}

func changeArr0901(arr [3]int)  {
	arr[0] = 100
	fmt.Println("函数执行中arr = ", arr)
}

func chageArr0902(arr *[3]int)  {
	(*arr)[0] = 100
	fmt.Println("函数执行中arr = ", *arr)
}
