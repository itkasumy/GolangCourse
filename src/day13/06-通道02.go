package main

import (
	"fmt"
	"time"
)

func main() {
	//go不推荐通过内存共享数据，而是通过通道传递消息（数据）
	//var 通道名 chan 通道里传递的数据类型
	// 通道传递数据：
		//传递数据到通道： ch <- 要传递的数据,会阻塞，直到有另一个goroutine把数据取出才会解除阻塞
		//从通道获取数据：接收数据的容器（变量） <- ch，也会发生阻塞，直到有另一个goroutine往通道里面存入数据才会解除阻塞

	ch01 := make(chan bool)

	go func() {

		time.Sleep(1)

		fmt.Println("我是子goroutine...")
		<- ch01
	}()

	ch01 <- true
	fmt.Println("main over...")
}
