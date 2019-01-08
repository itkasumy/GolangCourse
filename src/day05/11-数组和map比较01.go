package main

import "fmt"

func main() {
	arr1101 := [5]int{1, 2, 3, 5, 4}
	m1101 := map[string]int{"yi": 1, "er": 2, "san": 3}
	fmt.Println(arr1101)
	fmt.Println(m1101)
	fmt.Println(m1101["san"])
	fmt.Println(arr1101[2])
}
