package main

import "fmt"

func main() {
	//求给定的年份，是否是闰年
	year0201 := 2010
	if year0201 % 4 == 0 && year0201 % 100 != 0 || year0201 % 400 == 0 {
		fmt.Println("是闰年")
	} else {
		fmt.Println("是平年")
	}
}
