package main

import "fmt"

func main() {
	//4.  一个四位数，恰好等于去掉它的首位数字之后所剩的三位数的3倍，这个四位数是多少。（5‘’）

	for i := 1000; i < 10000; i++ {
		if i % 1000 * 3 == i {
			fmt.Println(i)
		}
	}
}
