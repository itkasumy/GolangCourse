package main

import "fmt"

func main() {

	//带缓存的通道，当缓存区满的时间才会发生阻塞(存),如果缓存区没有数据也会发生阻塞（取）
	ch01 := make(chan int)
	fmt.Printf("len(ch01) = %d, cap(ch01) = %d\n", len(ch01), cap(ch01))

	ch02 := make(chan int, 5)
	fmt.Printf("len(ch02) = %d, cap(ch02) = %d\n", len(ch02), cap(ch02))

	ch02 <- 1
	fmt.Printf("len(ch02) = %d, cap(ch02) = %d\n", len(ch02), cap(ch02))
	ch02 <- 2
	fmt.Printf("len(ch02) = %d, cap(ch02) = %d\n", len(ch02), cap(ch02))
	ch02 <- 3
	ch02 <- 4
	ch02 <- 5
	ch02 <- 6

	fmt.Println("main over...")
}
