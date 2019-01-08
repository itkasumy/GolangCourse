package main

import "fmt"

func main() {
	ch01 := make(chan int)

	go func() {
		fmt.Println("我是子goroutine...")
		//ch01<- 1
		<- ch01
	}()

	//<-ch01
	ch01<- 10

	fmt.Println("main over...")
}
