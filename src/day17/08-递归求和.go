package main

import "fmt"

func main() {
	//8. 写出函数求1!+2!+3!+....+n!（10‘’）
	//fmt.Println(getJ(6))
	sum := getSum(5)
	fmt.Println(sum)
}

func getJ(n int) int {
	if n == 1 {
		return 1
	}

	return n * getJ(n - 1)
}

func getSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += getJ(i)
	}

	return sum
}
