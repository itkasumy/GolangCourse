package main

import "fmt"

func main() {

	var str string
	/*a := 1
	b := a
	fmt.Println("a = ", a, "b = ", b)
	b = 100
	fmt.Println("a = ", a, "b = ", b)

	c := &a
	fmt.Println("a = ", a, "*c = ", *c)
	*c = 100
	fmt.Println("a = ", a, "*c = ", *c)*/

	a := 1
	b := 2
	c := 3

	arr1 :=  [3]int{}
	arr1[0] = a
	arr1[1] = b
	arr1[2] = c
	fmt.Println(arr1)
	p1 := &arr1

	fmt.Printf("%T\n", p1)

	var arr2 [3]*int // 一个长度为3的数组，数组里面存的是zhizhen
	//var arr3 *[3]int // 是一个指针变量，指向一个长度为3的int数组
	arr2[0] = &a
	arr2[1] = &b
	arr2[2] = &c
	fmt.Println(arr2)
	*arr2[0] = 100
	//(*arr2) 数组的指针
	//*arr2[0] 指针数组中的某一个元素
	fmt.Println(a)
}
