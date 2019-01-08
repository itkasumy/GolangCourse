package main

import "fmt"

func main() {
	// break 可以给一个循环取个名字，通过break关键字结束某个外层循环
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(j)
			if j == 2 {
				//break outer
			}
		}
	}
}
