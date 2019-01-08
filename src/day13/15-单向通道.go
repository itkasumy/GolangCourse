package main

func main() {
	ch00 := make(chan int)
	//ch01 := make(chan <- int)
	//ch02 := make(<- chan int)

	//ch00 <- 10
	//<- ch00

	//<- ch01
	//invalid operation: <-ch01 (receive from send-only type chan<- int)
	//ch01<- 10

	//ch02 <- 10
	//invalid operation: ch02 <- 10 (send to receive-only type <-chan int)
	//<- ch02

	go func(ch chan <- int) {

	}(ch00)
}
