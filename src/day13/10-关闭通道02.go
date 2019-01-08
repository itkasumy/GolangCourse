package main

import "fmt"

func main() {
	ch01 := make(chan int)

	go sendData10(ch01)

	//for range 可以取到通道中的值赋值给v,直到通道关闭为止
	for v := range ch01 { // v <- ch01
		fmt.Println(v)
	}

	fmt.Println("main over...")
}

func sendData10(ch chan int)  {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("\t子goroutine传来第%d条数据\n", i)
	}
	close(ch)
}
