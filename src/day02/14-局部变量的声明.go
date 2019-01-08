package main

import "fmt"

func main() {
	//给定要给数字，判断是正数，负数，零

	if num1301 := 22; num1301 > 0 {
		fmt.Println("正数")
	} else if num1301 < 0 {
		fmt.Println("负数")
	} else {
		fmt.Println("是0")
	}

	//fmt.Println(num1301) // undefined: num1301
}
