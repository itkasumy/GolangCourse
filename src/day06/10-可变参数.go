package main

import "fmt"

func main() {
	//可变参数
	//函数执行时传递的参数个数时可以变的
	/*
	注意点
		1. 可变的只是参数的**个数**
		2. 一个函数只可以有一个可变参数
		3. 如果有可变参数，可变参数必须放在参数列表的最后
	*/

	sum := getSum1001(1, 3, 4, 3, 6, 6 )
	fmt.Println(sum)
	fmt.Println(getSum1002(1, 3, 4, 3, 6, 6 ))

	//append(切片, 1, 2, 3, 4, 5)
	//append(切片, 切片...)

	sum1001 := getSum1002(1, 2)
	fmt.Println(sum1001)
}

func getSum1001(a, b, c, d, e, f int) int {
	return a + b + c + d + e + f
}

func getSum1002(elems... int) int {
	fmt.Printf("%T\n", elems) // []int
	sum := 0
	for _, val := range elems {
		sum += val
	}
	return sum
}
