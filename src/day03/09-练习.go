package main

import "fmt"

func main() {
	//求100以内,能被3整除，但是不能被5整除的数,只打印前5个
	//打印100以内前10个奇数

	couter0901 := 0
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 && i % 5 != 0 {
			fmt.Println(i)
			couter0901++
			if couter0901 == 5 {
				break
			}
		}
	}
	fmt.Println(couter0901)

	couter0902 := 0
	for i := 1; i < 100; i++ {
		if i % 2 != 0 {
			fmt.Println(i)
			couter0902++
			if couter0902 == 10 {
				break
			}
		}
	}
}
