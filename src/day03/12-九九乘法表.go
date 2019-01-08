package main

import "fmt"

func main() {
	//1. 打印99乘法表

	/*
	1 * 1 = 1
	1 * 2 = 2	2 * 2 = 4
	1 * 2 = 3	2 * 3 = 6	3 * 3 = 9
	1 * 4 = 4	2 * 4 = 8	3 * 4 = 12	4 * 4 = 16
	*/
	/*
	*
	**
	***
	****
	*/
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, j * i)
		}
		fmt.Println()
	}
}


