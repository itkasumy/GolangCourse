package main

import "fmt"

func main()  {
	/*
		练习1：给定一个年龄，如果大于22岁，输出可以娶媳妇啦
		练习2：给定一个数值，输出他的绝对值
		练习3：给定一个成绩，如果大于等于60，输出及格
	*/

	//1. 给定一个年龄，如果大于22岁，输出可以娶媳妇啦
	age0901 := 24
	if age0901 > 22 {
		fmt.Println("可以娶媳妇啦...")
	}

	//2. 给定一个数值，输出他的绝对值
	/*var num0901 int
	fmt.Println("请输入一个数字：")
	fmt.Scanln(&num0901)*/
	/*if num0901 >= 0 {
		fmt.Printf("%d的绝对值是%d\n", num0901, num0901)
	}
	if num0901 < 0 {
		num0902 := -num0901
		fmt.Printf("%d的绝对值是%d\n", num0901, num0902)
	}*/

	/*if num0901 < 0 {
		num0901 = -num0901
	}
	fmt.Println("输入数的绝对值是", num0901)*/

	//3. 给定一个成绩，如果大于等于60，输出及格
	score0901 := 59
	if score0901 >= 60 {
		fmt.Println("及格...")
	}
}
