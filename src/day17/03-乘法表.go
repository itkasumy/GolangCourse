package main

import "fmt"

func main() {
	//3. 打印乘法表（5‘’）
	mul()
}

func mul()  {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t\t", j, i, j * i)
		}
		fmt.Println()
	}
}
