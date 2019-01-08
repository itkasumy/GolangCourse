package main

import "fmt"

func main()  {
	/*
		流程结构：
			顺序结构： 程序至上而下一行一行的执行
			选择结构： 当满足某个条件的时候才执行，而且执行一次
							if 语句，switch 语句
							if 语句语法：
								if condition(布尔值) {
									如果condition为true执行的代码
								}
			循环结构： 当条件满足时，循环执行一段代码，直到条件不满足才不会执行
							for
	*/
	var num0801  = 10
	if num0801 < 10 {
		fmt.Println("num0801小于10")
	}

	fmt.Println("程序执行结束...")
}
