package main

import "fmt"

func main() {
	/*
		数组：长度是固定的，长度不可以改变
		切片：长度不是固定的，可以改变，容量也可以改变
	*/

	// 第一种创建方式：
	// 创建语法：var slice_name []Type
	// var slice_name = []Type{}
	var arr0301 = [5]int{}
	fmt.Println(arr0301)
	fmt.Println(len(arr0301), cap(arr0301))

	var slice0101 = []int{}
	fmt.Println(len(slice0101), cap(slice0101))

	var arr0302 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(len(arr0302))

	var slice0302 = []int{1, 2, 3, 4, 5}
	fmt.Println(len(slice0302), cap(slice0302))

	var arr0303 [5]int
	var slice0303 []int
	fmt.Println(arr0303, slice0303)

	//方式二：
	//创建语法： make([]T, length, capacity)
	fmt.Println("----------------------------------------------")
	slice0304 := make([]int, 3, 5)
	fmt.Println(slice0304)
	fmt.Println(len(slice0304), cap(slice0304))

	fmt.Println("----------------------------------------------")
	//方式三：
	//在原有数组的基础上创建切片
	arr0304 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr0304)

	slice0305 := arr0304[0: 5] // 相当于一个做闭右开的区间
	fmt.Println(slice0305)
	//从索引为2开始取，取到索引为5（不包括5）
	slice0306 := arr0304[2: 5]
	fmt.Println(slice0306)
	slice0307 := arr0304[0: 10]
	fmt.Println(slice0307)
	slice0308 := arr0304[3:]
	fmt.Println(slice0308)
	slice0309 := arr0304[0: 6]
	fmt.Println(slice0309)
	slice03010 := arr0304[:6]
	fmt.Println(slice03010)
	slice03011 := arr0304[:]
	fmt.Println(slice03011)
	// 容量和长度
	slice03012 := arr0304[3: 7]
	fmt.Println(len(slice03012), cap(slice03012))
	//注意：长度为所取数组被取到的元素个数，容量为从开始取的位置到数组末尾的元素个数
}
