package main

import "fmt"

func main() {
	/*
		if else 语法
		if 条件表达式 {
			如果条件表达式成立执行的代码
		} else {
			如果条件表达式不成了执行的代码
		}
		二者必需执行其中一个
	*/
	//给定一个数值，输出他的绝对值
	num1001 := -6
	if num1001 >= 0 {
		fmt.Println("绝对值是", num1001)
	} else {
		fmt.Println("绝对值是", -num1001)
	}
}
