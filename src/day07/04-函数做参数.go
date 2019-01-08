package main

import "fmt"

func main() {
	num1 := 1
	num2 := 2
	res1 := getSum0401(num1, num2)

	fmt.Println(res1)

	fn0401(num1, num2, getSum0401)

	res2 := fn0401(num1, num2, getSum0401)
	fmt.Println(res2)
	fmt.Println(getSum0401)

	res3 := fn0401(num1, num2, getSum0401)
	fmt.Println(res3)
}

func getSum0401(a, b int) int {
	return a + b
}

func fn0401(a int, b int, fn func (int, int) int) int {
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", fn)
	fmt.Println(fn)
	//res := fn(a, b)
	//return res
	return fn(a, b)
}
