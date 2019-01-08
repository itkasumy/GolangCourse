package main

import "fmt"

//var num0101 int
// num0102 := 3 // syntax error: non-declaration statement outside function body

func main() {
	/*
		变量:
			定义: 本质上是一块内存，我们通过一个标识去识别它
			语法: var 变量名 变量类型
				变量名 = 值
	*/

	var chinese int
	chinese = 99
	//fmt.Printf("%p\n", &chinese)
	fmt.Println("my chinese score is ", chinese)

	chinese = 100
	fmt.Println("my chinese score after change is", chinese)
	// chinese = "100" // cannot use "100" (type string) as type int in assignment

	// 2. 变量声明方式二：
	var english int = 92
	fmt.Println("my english score is", english)

	// 3. 变量声明方式三：
	var math= 100
	fmt.Printf("my math score is %d, type is %T\n", math, math)
	// math = "hello" // cannot use "hello" (type string) as type int in assignment

	// 4. 变量声明方式四：
	num := 1 // undefined: num
	fmt.Println("num = ", num)

	//多变量声明
	//方式一
	var a, b, c int
	a = 1
	b = 2
	c = 3
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

	//方式二
	var d, e, f= 10, true, "hello"
	fmt.Println(d, e, f)
	fmt.Printf("type of d is %T, type of e is %T, type of f is %T\n", d, e, f)

	//方式三
	var (
		g int = 2
		h    = false
		i    = "golang"
	)
	fmt.Println(g, h, i)
}