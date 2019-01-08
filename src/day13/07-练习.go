package main

import "fmt"

func main() {
	//创捷一个子goroutine，打印100个数字，保证打印完成，用chan实现

	ch01 := make(chan string)

	go func() {
		for i := 1; i <= 100; i++ {
			fmt.Println(i)
		}

		<- ch01
	}()

	ch01 <- "xunlun"
	fmt.Println("main over...")
}
