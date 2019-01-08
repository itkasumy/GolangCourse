package main

import (
	"fmt"
	"sync"
)

func main() {
	// 有四个窗口卖票，总共是100张票，四个窗口并行卖票
	var wg sync.WaitGroup
	wg.Add(4)

	go saleTickets18("窗口01", &wg)
	go saleTickets18("窗口02", &wg)
	go saleTickets18("窗口03", &wg)
	go saleTickets18("窗口04", &wg)

	wg.Wait()

	fmt.Println("main over...")
}

func saleTickets18(name string, wg *sync.WaitGroup)  {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		fmt.Printf("%s卖票 => %d\n", name, i)
	}
}
