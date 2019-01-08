package main

import "fmt"

func fna()  {
	fmt.Println("aaaaaaaaaa")
}

func fnb()  {
	fmt.Println("bbbbbbbbb")
}

func fnc()  {
	fmt.Println("ccccccccccc")
}

func fnd()  {
	fmt.Println("ddddddddd")
}

func main() {
	/*defer: 延迟
		延迟函数/方法执行
		后进先出（Last In First Out）LIFO
	*/
	defer fna()
	fnb()
	defer fnc()
	fnd()
	defer fmt.Println("eeeeeeeeee")
}
