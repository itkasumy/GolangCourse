package main

import "fmt"

type person08 struct {
	string
	int
	gender string
}

func main() {
	p0801 := person08{"曹操", 26, "男"}
	fmt.Println(p0801)
	fmt.Println(p0801.gender)
	fmt.Println(p0801.string)

	p0802 := person08{gender: "男", int: 24, string: "孙权"}
	fmt.Println(p0802)

	p0803 := struct {
		string
		int
		gender string
	}{"刘备", 22, "男"}
	fmt.Println(p0803)
}
