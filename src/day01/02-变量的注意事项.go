package main

import "fmt"

func main()  {
	/*
		变量的注意事项：
			1. 必须先声明再使用，变量只能声明一次
			2. go是强类型的语言，变量的类型和赋值必须一致
			3. :=赋值的时候需要注意，左边的变量名至少有一个是新的（之前未定义过的）
			4. 零值（默认值），每一种数据类型都有自己的初始化值，不同数据类型的初始化值是不一样的
			5. := 我们可以用_省略接收值
	*/
	//num0201 = 3
	var num0202 = 3
	fmt.Println(num0202)
	//var num0202 int


	//var num0203 int
	//num0203 = 8
	//num0203 = "abc" // cannot use "abc" (type string) as type int in assignment

	var num0204 int
	var num0205 int
	//num0204, num0205 := 2, 3 // no new variables on left side of :=
	num0204, num0205, num0206 := 1, 2, 3
	fmt.Println(num0204, num0205, num0206)

	var num0207 int
	fmt.Println(num0207)
	var b0201 bool
	var str0201 string
	fmt.Println(b0201, str0201)
	fmt.Printf("%t,'%s'\n", b0201, str0201)
	str0201 = "hello"
	str0201 = ""

	a, _ := num0205, num0206
	fmt.Println(a)
}
