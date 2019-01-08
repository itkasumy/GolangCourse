package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	//fmt.Printf("%T\n", wg)

	go prtNum16(&wg)
	go prtLetter16(&wg)

	fmt.Println("main执行中")

	wg.Wait()

	fmt.Println("main over...")
}

func prtNum16(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		fmt.Printf("i = %d\n", i)
	}
}

func prtLetter16(wg *sync.WaitGroup)  {
	defer wg.Done()
	for j := 1; j <= 100; j++ {
		fmt.Printf("\tj = %c\n", j)
	}
}
