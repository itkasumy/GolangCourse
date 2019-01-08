package main

import "fmt"

func main()  {
	/*
		常量：
			道可道，非常道。常：不变的，固定的。
			定义： 存储的值是固定的，不可以被改变。
			语法： const 常量名 [数据类型] = 常量的值（可以是一个固定的值，也可以是一个常量表达式）
	*/
	const PI float64 = 3.1415926
	fmt.Println("PI = ", PI)
	//PI = 3.14 // cannot assign to PI

	const _LENGTH = 4
	const cWIDTH = 3
	const cAREA = _LENGTH * cWIDTH
	fmt.Println("cAREA = ", cAREA)

	//var num0301 = 3
	//var num0302 = 4
	//const cAREA02 = num0301 * num0302 // const initializer num0301 * num0302 is not a constant

	const cSTR = "helloworld"
	fmt.Println(len(cSTR))
	const cLEN0301 = len(cSTR)
	fmt.Println("cLEN0301 = ", cLEN0301)

	//var str0301 = "hello"
	//const cLEN0302 = len(str0301) // const initializer len(str0301) is not a constant

	// 多常量赋值
	const cNUM0301, cNUM0302 = 1, "我是一个小小的石头"
	fmt.Println(cNUM0301, cNUM0302)
	fmt.Printf("type of cNUM0301 is %T, type of cNUM0302 is %T\n", cNUM0301, cNUM0302)

	const (
		cNUM0303 = "好好学习，天天向上"
		cNUM0304 = 10
	)
	fmt.Println(cNUM0303, cNUM0304)
}
