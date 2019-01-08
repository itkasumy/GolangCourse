package main

import "fmt"

func main() {

	//练习1：定义一个函数，接收两个参数，用于比较两个数的大小，并将大的数返回。
	fmt.Println(compare0801(1, 2))
	//练习2：定义一个函数，用于接收两个参数，求和，并将结果返回。
	fmt.Println(getSum0802(1, 2))
	//练习3：定义一个函数，用于求圆的面积和周长，并返回结果。
	//面积： PI* r * r
	//周长：2 * PI* r
	mianji, zhouchang := getCircle0801(2)
	fmt.Println(mianji, zhouchang)
}

func compare0801(a int, b int) (bigger int) {
	if a > b {
		bigger = a
	} else {
		bigger = b
	}
	return bigger
}

func compare0802(a int, b int) int {
	if a > b {
		return a
	} else {
		return  b
	}
}

func compare0803(a, b int) int {
	if a > b {
		return a
	} else {
		return  b
	}
}

func getSum0801(a, b int) (sum int) {
	sum = a + b
	return sum
}

func getSum0802(a, b int) int {
	return a + b
}

func getCircle0801(r float64) (area float64, zc float64) {
	const PI float64 = 3.14
	area = PI * r * r
	zc = 2 * PI * r
	return
}
