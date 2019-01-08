package main

import "fmt"

func main() {
	//创建并启动两个goroutine，一个打印100个数字，一个打印100个字符，用chan实现
	ch01 := make(chan bool)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
		<-ch01
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Printf("%c\n", i)
		}
		<-ch01
	}()

	ch01 <- true
	ch01 <- true
	fmt.Println("main over...")
}
