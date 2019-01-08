package main

import "fmt"

func main() {// 程序的入口，不需要手动调用
	fna(3)
	fmt.Println("main")
}

func fna(a int)  {
	fnb(a - 1)
	fmt.Println(a)
}

func fnb(b int)  {
	fnc(b - 1)
	fmt.Println(b)
}

func fnc(c int)  {
	fmt.Println(c)
}
