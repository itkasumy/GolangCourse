package main

import (
	"fmt"
	"sync"
)

func main() {
	// 如果等待组里添加的数量不能全部结束，会发生死锁错误
	// fatal error: all goroutines are asleep - deadlock!
	var wg sync.WaitGroup
	wg.Add(3)

	go prtNum17(&wg)
	go prtLetter17(&wg)

	fmt.Println("main执行中")

	wg.Wait()

	fmt.Println("main over...")
}

func prtNum17(wg *sync.WaitGroup)  {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		fmt.Printf("i = %d\n", i)
	}
}

func prtLetter17(wg *sync.WaitGroup)  {
	defer wg.Done()
	for j := 1; j <= 100; j++ {
		fmt.Printf("\tj = %c\n", j)
	}
}
