package main

import "fmt"

func main() {
	/*
		switch: 选择结构
			语法：
		switch var1 {
			case val1:
				...
			case val2:
				...
			default:
				...
		}
	*/

	season1501 := 3

	switch season1501 {
	case 1:
		fmt.Println("第一季度")
	case 2:
		fmt.Println("第二季度")
	case 3:
		fmt.Println("第三季度")
		fmt.Println("第三季度")
		fmt.Println("第三季度")
	case 4:
		fmt.Println("第四季度")
	}
}
