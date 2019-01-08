package main

import "fmt"

func main() {
	//1. 不写条件
	switch { //相当于是true
	case true:
		fmt.Println("皇家寻论学院")
	case false:
		fmt.Println("厦门思明区")
	}

	//给定一个数字，如果是1：打印熊大，如果是2：熊二，如果是3，光头强，否则不知道
	var num1601 = 12
	switch {
	case num1601 == 1:
		fmt.Println("熊大")
	case num1601 == 2:
		fmt.Println("熊二")
	case num1601 == 3:
		fmt.Println("光头强")
	default:
		fmt.Println("不知道")
	}
	
	// 给定月份，输出该月的天数。
	/*month1601 := 7
	day1601 := 0
	switch month1601 {
	case 1, 3, 5, 7, 8, 10 , 12:
		day1601 = 31
	case 4, 6, 9, 11:
		day1601 = 30
	case 2:
		day1601 = 28
	default:
		fmt.Println("您输入的月份有误")
	}
	fmt.Println(day1601)*/

	// switch 可以初始化一个变量
	// 给定月份，输出该月的天数。

	switch month1601 := 7; month1601 {
	case 1, 3, 5, 7, 8, 10 , 12:
		fmt.Printf("%d月有31天\n", month1601)
	case 4, 6, 9, 11:
		fmt.Printf("%d月有30天\n", month1601)
	case 2:
		fmt.Printf("%d月有28天\n", month1601)
	default:
		fmt.Println("您输入的月份有误")
	}

	//fmt.Println(month1601) // undefined: month1601
}
