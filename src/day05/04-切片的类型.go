package main

import "fmt"

func main() {

	//数组是值类型的，赋值给另外一个数组时，拷贝的时数组里面的值，新数组修改对原数组没有影响
	arr0401 := [3]int{1, 2, 3}
	arr0402 := arr0401

	fmt.Println(arr0401)
	fmt.Println(arr0402)

	arr0402[0] = 100

	fmt.Println(arr0401)
	fmt.Println(arr0402)

	fmt.Println("+++++++++++++++++")

	//切片是引用类型的，它赋值时是将引用的地址赋值给新的遍历，所以他们指向的是同一片内存，当一个地方修改数据，其他所有引用的地方都会跟着改变
	s0401 := []int{1, 2, 3}
	s0402 := s0401
	fmt.Println(s0401)
	fmt.Println(s0402)

	s0402[0] = 100
	fmt.Println(s0401)
	fmt.Println(s0402)
}
