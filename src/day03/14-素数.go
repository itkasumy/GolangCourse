package main

import "fmt"

func main() {
	//3. 打印2-100内的素数。素数（只能被1和其本身整除的数）
	//6 7
	for i := 2; i < 100; i++ {
		flag := true
		for j := 2; j < i / 2; j++ {
			if i % j == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(i)
		}
	}
}
