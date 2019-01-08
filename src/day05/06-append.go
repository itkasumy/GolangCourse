package main

import "fmt"

func main() {
	//append():从后面添加一条数据
	//append(往哪个切片里面添加数据, 添加的数据) 返回一个切片
	s0601 := make([]int, 3, 5)
	fmt.Println(s0601)
	fmt.Println(len(s0601), cap(s0601))
	s0601 = append(s0601, 1, 2)
	fmt.Println(s0601)
	fmt.Println(len(s0601), cap(s0601))

	s0601 = append(s0601, 3)
	fmt.Println(s0601)
	fmt.Println(len(s0601), cap(s0601))
	// 容量是怎么变化的？1. 每次append一个如果容量不够，就会扩容，以2倍的容量进行扩容

	s0601 = append(s0601, 4, 5, 6)
	fmt.Println(s0601)
	fmt.Println(len(s0601), cap(s0601))

	s0601 = append(s0601, 8, 9, 10)
	fmt.Println(s0601)
	fmt.Println(len(s0601), cap(s0601))

	fmt.Println("-----------------------------------")

	s0602 := []int{}
	fmt.Println(s0602)
	fmt.Println(len(s0602), cap(s0602))
	s0602 = append(s0602, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	// 容量是怎么变化的？2. 每次append一个如果容量不够，就会扩容，以+2的容量进行扩容
	fmt.Println(s0602)
	fmt.Println(len(s0602), cap(s0602))
}
