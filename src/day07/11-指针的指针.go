package main

import "fmt"

func main() {
	a := 100 // int
	var p1 *int
	p1 = &a // *int

	var p2 **int
	p2 = &p1

	fmt.Println(p1)
	fmt.Println(*p1)

	fmt.Println(p2)
	fmt.Println(*p2)
	fmt.Println(**p2)

	**p2 = 1000
	fmt.Println(a)
}
