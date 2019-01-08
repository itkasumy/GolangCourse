package main

import "fmt"

type person struct {
	name string
}

func (p person) prtInfo09() {
	fmt.Printf("%s\n", p.name)
}

func main() {
	fmt.Println("xunlun")
	p01 := person{"zhangsan"}
	defer p01.prtInfo09()
	fmt.Println("xueyuan")
}
