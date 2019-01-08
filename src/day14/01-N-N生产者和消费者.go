package main

import "fmt"

func main() {
	//写出N个生产者对N个消费者的代码
	const N = 5
	const M = 6
	ch01 := make(chan int)
	done := make(chan int)

	for i := 1; i <= N; i++ {
		i := i
		go func() {
			for j := 1; j <= 100; j++  {
				ch01 <- j
				fmt.Println("生产者", i, "生产数据：", j)
			}
			done <- i
		}()
	}

	for i := 1; i <= M; i++ {
		i := i
		go func() {
			for v := range ch01 {
				fmt.Println("消费者", i, "消费数据:", v)
			}
			done <- i
		}()
	}

	for i := 1; i <= N; i++ {
		<- done
	}
	fmt.Println("生产者生产数据over...")
	close(ch01)

	for i := 1; i <= M; i++ {
		<- done
	}
	fmt.Println("消费者消费完成...")
}
