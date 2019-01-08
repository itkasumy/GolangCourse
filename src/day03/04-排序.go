package main

import "fmt"

func main() {
	//给定3个整数，从大到小输出
	num0401 := 0
	num0402 := 0
	num0403 := 0
	fmt.Println("请输入三个整数：")
	fmt.Scanln(&num0401, &num0402, &num0403)

	if num0401 > num0402 {
		//num1 > num2
		if num0401 > num0403 {
			// num1是最大的
			if num0402 > num0403 {
				fmt.Printf("%d > %d > %d", num0401, num0402, num0403)
			} else {
				fmt.Printf("%d > %d > %d", num0401, num0403, num0402)
			}
		} else {
			// num3是最大的
			fmt.Printf("%d > %d > %d", num0403, num0401, num0402)
		}
	} else {
		//num2 > num1
		if num0402 > num0403 {
			// num2是最大的
			if num0401 > num0403 {
				fmt.Printf("%d > %d > %d", num0402, num0401, num0403)
			} else {
				fmt.Printf("%d > %d > %d", num0402, num0403, num0401)
			}
		} else {
			//num3是最大的
			fmt.Printf("%d > %d > %d", num0403, num0402, num0401)
		}
	}
}
