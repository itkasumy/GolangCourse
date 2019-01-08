package main

import "fmt"

type person12 struct {
	name string
	age int
}

type student struct {
	person12
	study string
}

func main() {
	//继承：一个类作为另一个类的子类可以继承其属性和方法 is a
	//避免重复的代码，提高代码的复用性
	//提高程序的扩展型
	// 用法是将复结构体作为子结构体的一个匿名字段，此字段会提升
	//p1201 := person12{"荀彧", 50}
	p1202 := student{person12{"xiaozhao", 19}, "keku"}
	fmt.Println(p1202)
	fmt.Println(p1202.name)

	p1203 := student{
		person12: person12{"xiaowang", 22},
		study: "qingfu",
	}
	fmt.Println(p1203.name)
	fmt.Println(p1203.person12.name)
}
