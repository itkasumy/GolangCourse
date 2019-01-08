package main

import "fmt"

type person06 struct {
	name string
	age int
}

func main() {
	p0601 := person06{"王五", 18}
	fmt.Println(p0601)

	//changeName0601(p0601)
	//fmt.Println(p0601)
	changeName0602(&p0601)
	fmt.Println(p0601)
}

func changeName0601(p person06)  {
	p.age = 100
	fmt.Printf("p的姓名是: %s, 年龄是：%d\n", p.name, p.age)
}

func changeName0602(p *person06)  {
	p.age = 100
	//fmt.Printf("p的姓名是: %s, 年龄是：%d\n", p.name, p.age)
	fmt.Printf("p的姓名是: %s, 年龄是：%d\n", (*p).name, (*p).age)
}
