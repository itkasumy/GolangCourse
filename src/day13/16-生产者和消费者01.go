package main

import "fmt"

func main() {
	//1个生产者两个消费者
	ch01 := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 1; i <= 100; i++ {
			ch01<- i
			fmt.Println("生产者生产数据：", i)
		}
		close(ch01)
	}()

	go func() {
		for v := range ch01 {
			fmt.Println("\t\t消费者1消费数据：", v)
		}
		done<- true
	}()

	go func() {
		for v := range ch01 {
			fmt.Println("\t\t消费者2消费数据：", v)
		}
		done<- true
	}()

	<-done
	<-done
	fmt.Println("main over...")
}
