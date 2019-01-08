package main

import "fmt"

func main() {
	/*打印1：

	*****
	*****
	*****
	*****
	*****
	*/

	/*for i := 0; i < 5; i++ {
		fmt.Println("*****")
	}*/

	/*for i := 0; i < 5; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	for i := 0; i < 5; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	for i := 0; i < 5; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	for i := 0; i < 5; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	for i := 0; i < 5; i++ {
		fmt.Print("*")
	}
	fmt.Println()*/

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	/*
	*
	**
	***
	****
	*****
	*/
	//第一次执行1遍
	//第二次执行2遍
	//。。。
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
