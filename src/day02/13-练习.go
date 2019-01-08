package main

import "fmt"

func main() {
	/*
		练习1：给定要给数字，判断是正数，负数，零
		练习2：给定一个数字，如果是1：打印熊大，如果是2：熊二，如果是3，光头强，否则不知道
		练习3：使用if语句，
					成绩：
					[90,100],优秀
					[80,89],良好
					[70,79],中等
					[60,69],及格
					[0,59],不及格
		练习4：给定月份，输出该月的天数。
	*/

	//1. 给定要给数字，判断是正数，负数，零
	num1301 := 0
	if num1301 > 0 {
		fmt.Println("正数")
	} else if num1301 < 0 {
		fmt.Println("负数")
	} else {
		fmt.Println("是0")
	}

	//2. 给定一个数字，如果是1：打印熊大，如果是2：熊二，如果是3，光头强，否则不知道
	num1302 := 20
	if num1302 == 1 {
		 fmt.Println("熊大")
	} else if num1302 == 2 {
		fmt.Println("熊二")
	} else if num1302 == 3 {
		fmt.Println("光头强")
	} else {
		fmt.Println("不知道")
	}

	/*3. 使用if语句，成绩：
			[90,100],优秀
			[80,89],良好
			[70,79],中等
			[60,69],及格
			[0,59],不及格*/
	score1301 := 67
	if score1301 >= 90 {
		fmt.Println("优秀")
	} else if score1301 >= 80 {
		fmt.Println("良好")
	} else if score1301 >= 70 {
		fmt.Println("中等")
	} else if score1301 >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}

	//4. 给定月份，输出该月的天数。
	month1301 := 11
	day1301 := 0
	if month1301 == 1 || month1301 == 3 || month1301 == 5 || month1301 == 7 || month1301 == 8 || month1301 == 10 || month1301 == 12 {
		day1301 = 31
	} else if month1301 == 2 {
		day1301 = 28
	} else {
		day1301 = 30
	}
	fmt.Println(day1301)
}
