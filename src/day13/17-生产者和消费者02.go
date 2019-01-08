package main

import "fmt"

func main() {
	//2个生产者1个消费者
	ch01 := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 1; i <= 100; i++ {
			ch01<- i
			fmt.Println("\t\t生产者1生产数据：", i)
		}
		done<- true
	}()

	go func() {
		for i := 101; i <= 200; i++ {
			ch01<- i
			fmt.Println("\t生产者2生产数据：", i)
		}
		done<- true
	}()

	go func() {
		for v := range ch01 {
			fmt.Println("消费者消费数据：", v)
		}
		done<- true
	}()

	<-done
	<-done

	close(ch01)
	<-done
	fmt.Println("main over...")
}
