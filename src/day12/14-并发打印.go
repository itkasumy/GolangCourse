package main

import (
	"fmt"
	"time"
)

func main() {
	go prtNum14()
	go prtLetter14()

	time.Sleep(100)
	fmt.Println("main over...")
}

func prtNum14()  {
	for i := 1; i <= 100; i++ {
		fmt.Printf("i = %d\n", i)
	}
}

func prtLetter14()  {
	for j := 1; j <= 100; j++ {
		fmt.Printf("\tj = %c\n", j)
	}
}
