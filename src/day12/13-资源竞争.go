package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		go不推荐通过内存共享数据
		而是通过通道来传递消息
	*/

	num := 1

	go func() {
		num = 2
		fmt.Println("我是子线程1里打印的num =", num)
	}()

	go func() {
		num = 4
		fmt.Println("我是子线程2里打印的num =", num)
	}()

	num = 3
	time.Sleep(100)
	fmt.Println("我是main里打印的num =", num)

	fmt.Println("main over...")
}
