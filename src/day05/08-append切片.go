package main

import "fmt"

func main() {
	s0801 := []int{1, 2, 3, 4, 5}
	s0802 := []int{6, 7, 8, 9}
	fmt.Println(s0801, s0802)

	/*for _, val := range s0802{
		s0801 = append(s0801, val)
	}
	fmt.Println(s0801, s0802)*/

	//解构
	s0801 = append(s0801, s0802...)
	fmt.Println(s0801, s0802)
}
