package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go prtNum15()
	go prtLetter15()

	fmt.Println("main执行中")

	wg.Wait()

	fmt.Println("main over...")
}

func prtNum15()  {
	defer wg.Add(-1)
	for i := 1; i <= 100; i++ {
		fmt.Printf("i = %d\n", i)
	}
}

func prtLetter15()  {
	defer wg.Done()
	for j := 1; j <= 100; j++ {
		fmt.Printf("\tj = %c\n", j)
	}
}
