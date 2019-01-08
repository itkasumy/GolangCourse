package main

import (
	"fmt"
)

func main() {
	arr0401 := [5]int {1, 2, 3, 4, 5}
	fmt.Println(arr0401)

	/*fmt.Println(arr0401[0])
	fmt.Println(arr0401[1])
	fmt.Println(arr0401[2])
	fmt.Println(arr0401[3])
	fmt.Println(arr0401[4])*/

	for i := 0; i < len(arr0401); i++ {
		fmt.Println(arr0401[i])
	}

	/*arr0402 := [7]int {1, 2, 3, 4, 5, 6, 7}
	for i := 0; i < len(arr0402); i++ {
		fmt.Println(arr0402[i])
	}*/

	// for ranage 用于遍历数组
	//把数组的索引和对应的存储的数值作为一对返回给我们的变量
	//不需要之前的判断条件，内部会遍历数组的索引，当遍历数组最后一项就自动结束了
	for index, val := range arr0401 {
		fmt.Println(index, "=>", val)
	}

	fmt.Println("-------------")

	sum := 0
	for _, val := range arr0401 {
		sum += val
		fmt.Println(val)
	}

	fmt.Println("-------------")

	for i, _ := range arr0401{
		fmt.Println(i)
	}

	fmt.Println("+++++++++")

	for i := range arr0401 {
		fmt.Println(i, arr0401[i])
	}
}
