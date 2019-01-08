package main

import "fmt"

func main() {
	//map的键是唯一的，不可以重复定义,map通过键检索值
	m1001 := make(map[int]string)

	m1001[1] = "Golang"
	m1001[2] = "Java"
	m1001[3] = "Python"
	m1001[2] = "JavaScript"
	fmt.Println(m1001)

	fmt.Println(m1001[1])
	fmt.Println(m1001[2])
	fmt.Println(m1001[3])
}
