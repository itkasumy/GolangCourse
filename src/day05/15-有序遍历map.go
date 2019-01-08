package main

import (
	"fmt"
	"sort"
)

func main() {
	m1501 := map[int]string {
		1: "Golang",
		2: "Java",
		3: "Python",
		4: "JavaScript",
	}
	fmt.Println(m1501)

	s1501 := []int{}

	for key := range m1501{
		s1501 = append(s1501, key)
	}
	fmt.Println(s1501)

	sort.Ints(s1501)
	fmt.Println(s1501)

	for _, val := range s1501 {
		fmt.Println(m1501[val])
	}
}
