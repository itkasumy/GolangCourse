package main

import "fmt"

func main() {

	//map是引用类型的
	m1601 := make(map[int]string)

	m1601[1] = "Golang"
	m1601[2] = "Java"
	m1601[3] = "Python"

	fmt.Println(m1601)
	fmt.Printf("%T\n", m1601)

	m1602 := m1601
	fmt.Println(m1602)
	fmt.Printf("%T\n", m1602)

	m1602[2] = "JavaScript"
	fmt.Println(m1602)
	fmt.Println(m1602)
}
