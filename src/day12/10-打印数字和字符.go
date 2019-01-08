package main

import (
	"fmt"
	"time"
)

func main() {
	go prtNum10()
	go prtLetter10()

	time.Sleep(time.Second)
}

func prtNum10()  {
	for i := 0; i < 100; i++ {
		fmt.Println("i =", i)
	}
}

func prtLetter10()  {
	for i := 0; i < 100; i++ {
		fmt.Printf("\t%c\n", i)
	}
}
