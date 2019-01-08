package main

import "fmt"

type person05 struct {
	name string
	age int
}

func main() {
	p0501 := person05{
		name: "李四",
		age: 22,
	}
	fmt.Println(p0501)
	p0502 := p0501
	fmt.Println(p0502)
	p0502.age = 20
	fmt.Println(p0501)
	fmt.Println(p0502)

	var p0503 *person05
	p0503 = &p0501
	p0503.age = 24
	fmt.Println(p0501)
	fmt.Println(*p0503)
}
