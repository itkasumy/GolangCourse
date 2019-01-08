package main

import "fmt"


func main() {
	//匿名结构体：没有名字的结构体
	zs := struct {
		name string
		age int
	}{"张三", 18}
	fmt.Println(zs)

	lisi := struct {
		name string
		age int
	}{
		age: 22,
		name: "lisi",
	}
	fmt.Println(lisi)
}
