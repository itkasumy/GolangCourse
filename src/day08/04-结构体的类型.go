package main

import "fmt"

type person0401 struct {
	name string
	age int
}

func main() {
	// 结构体是值类型的
	p0401 := person0401{
		name: "赵云",
		age: 22,
	}
	fmt.Println(p0401)
	fmt.Printf("%T\n", p0401) // main.person0401

	p0402 := p0401
	p0402.age = 19
	fmt.Println(p0401, p0402) // 是值类型

	p0405 := &p0401
	p0405.age = 1000
	fmt.Println("================")
	fmt.Println(p0401, *p0405)
	fmt.Println("================")


	//创建的是指针变量，赋值赋的是地址
	p0403 := new(person0401)
	p0403.name = "ksm"
	p0403.age = 18
	fmt.Println(*p0403)
	fmt.Printf("%p\n", p0403)

	p0404 := p0403
	p0404.age = 19
	fmt.Println(*p0403, *p0404)
	fmt.Printf("%p\n", p0404)
}
