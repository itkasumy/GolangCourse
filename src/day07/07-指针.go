package main

import "fmt"

func main() {
	/*
	指针是什么？
		指针（存储另外一个变量的地址）是变量
	变量的组成：
		变量名
		变量类型
		变量地址 &0x9999
		变量的值
	指针的类型： *variable_type
	指针存储变量的类型：
	指针的地址：&指针
	指针存储的变量的地址：指针的值
	*/
	a := 1
	fmt.Printf("%p\n", &a)

	var p1 *int
	fmt.Printf("%T\n", p1) // *int
	p1 = &a
	fmt.Println(p1)
	fmt.Println(&p1)

	fmt.Println(a)
	fmt.Println(*p1)
	*p1 = 100
	fmt.Println(a)
	fmt.Println(*p1)

	/*b := a
	fmt.Println(b)
	b = 100
	fmt.Println(a, b)*/
}
