package main

import "fmt"

func main() {
	//启动一个子goroutine，向缓存通道中写入数据，再启动另外一个子goroutine,从中读取数据
	ch01 := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for i := 1; i <= 100; i++ {
			ch01 <- i
			fmt.Println("往通道中写入数据：", i)
		}
		close(ch01)
		done<- true
	}()

	go func() {
		for v := range ch01 {
			fmt.Println("\t\t从通道中读取数据：", v)
		}
		done<- true
	}()

	<- done
	<- done
	fmt.Println("main over...")
}
