package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4}
	arr2 := arr1
	arr2[0] = 100
	fmt.Println(arr1)
	fmt.Printf("%p, %p\n", &arr1, &arr2)

	var num1 int = 1
	num2 := num1
	num2 = 100
	fmt.Println(num1)
	fmt.Printf("%p, %p\n", &num1, &num2)
}
