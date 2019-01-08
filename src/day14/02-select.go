package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*switch expr {
	case :
	case :
	case :
	default:
	}*/
	ch01 := make(chan int)
	ch02 := make(chan int)

	go func() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(60)))
		ch01<- 1
	}()

	go func() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(60)))
		ch02 <- 2
	}()

	select {
	case num1 := <- ch01:
		fmt.Println(num1)
	case num2 := <- ch02:
		fmt.Println(num2)
	/*default:
		fmt.Println("no common")*/
	}
}
