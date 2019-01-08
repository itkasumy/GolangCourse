package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret) // 20, 0, 2, 2 // 2, 0, 2, 2 // 10, 1, 2, 3 // 1, 1, 3, 4
	return ret
}

func main() {
	a := 1
	//b := 2

	defer func(m int) {
		fmt.Println(m)
	}(a)

	//defer calc("1", a, calc("10", a, b)) // => calc("1", 1, calc("10", 1, 2)) => calc("1", 1, 3)

	a = 0

	//defer calc("2", a, calc("20", a, b)) // => calc("2", 0, calc("20", 0, 2)) => calc("2", 0, 2)

	//b = 1
}
