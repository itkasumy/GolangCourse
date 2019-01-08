package main

import "fmt"

type person0301 struct {
	name string
	age int
	gender string
}

func main() {
	//var num0301 int
	var p0301 person0301
	p0301.name = "张三丰"
	p0301.age = 108
	p0301.gender = "男"
	fmt.Println(p0301)


	// 顺序必须和定义时一致
	p0302 := person0301{"张思枫", 18, "女"}
	fmt.Println(p0302)

	p0303 := person0301{name: "刘备",gender: "男", age: 25}
	fmt.Println(p0303)

	p0304 := person0301{
		age: 24,
		name: "关羽",
		gender: "男",
	}
	fmt.Println(p0304)

	p0305 := new(person0301) // 是一个指针
	fmt.Printf("%T\n", p0305)
	(*p0305).name = "张飞"
	p0305.age = 23
	p0305.gender = "男"
	//fmt.Println(p0305)
	fmt.Println(*p0305)
}
