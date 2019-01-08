package main

import "fmt"

func main() {
	arr0501 := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr0501)

	s0501 := arr0501[2:]
	fmt.Println(s0501)
	s0502 := s0501
	fmt.Println(s0502)

	s0502[0] = 100
	fmt.Println(s0501)
	fmt.Println(s0502)
	fmt.Println(arr0501)
}
