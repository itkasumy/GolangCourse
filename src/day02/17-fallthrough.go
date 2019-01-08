package main

import "fmt"

func main() {
	//fallthrough
	// 给定月份，输出该月的天数。
	month1701 := 9
	day1701 := 0
	switch month1701 {
	case 1:
		fallthrough
	case 3:
		fallthrough
	case 5:
		fallthrough
	case 7:
		fallthrough
	case 8:
		fallthrough
	case 10:
		fallthrough
	case 12:
		day1701 = 31
	case 4:
		fallthrough
	case 6:
		fallthrough
	case 9:
		fallthrough
	case 11:
		day1701 = 30
	case 2:
		day1701 = 28
	}
	fmt.Println(day1701)

	var num1701 = 1
	switch num1701 {
	case 1:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	}
}
