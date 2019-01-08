package main

import "fmt"

func main() {
	/*
		if else 的嵌套结构
		if 条件表达式 {
			if为true执行的代码
		} else {
			if 条件表达式2 {
				条件表达式2成立执行的代码
			} else {
				条件1为false且条件2也为false执行的代码
			}
		}

		if 条件表达式 {
			if 条件表达式2 {
				条件表达式1为true且条件表达式2为true
			} else {
				条件表达式1为true条件表达式2为false
			}
		} else {
			条件表达式1为false执行的代码
		}
	*/

	//练习2：给定一个性别：如果是男，输出去男厕所，否则去女厕
	gender1201 := "太监"

	/*if gender1201 == "男" {
		fmt.Println("去男厕所...")
	} else {
		if gender1201 == "女" {
			fmt.Println("去女厕所...")
		} else {
			fmt.Println("随便去哪...")
		}
	}*/

	/*if gender1201 != "女" {
		if gender1201 == "男" {
			fmt.Println("去男厕所...")
		} else {
			fmt.Println("随便去哪...")
		}
	} else {
		fmt.Println("去女厕所...")
	}*/

	if gender1201 == "男" {
		fmt.Println("去男厕所...")
	} else if gender1201 == "女" {
		fmt.Println("去女厕所...")
	} else {
		fmt.Println("随便去哪...")
	}
}
