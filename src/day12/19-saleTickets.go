package main

import (
	"fmt"
	"sync"
)

func main() {
	// 有四个窗口卖票，总共是100张票，四个窗口并行卖票
	var wg sync.WaitGroup
	wg.Add(4)
	tickets := 80
	var mutex sync.Mutex

	go saleTickets19("窗口01", &wg, &tickets, &mutex)
	go saleTickets19("窗口02", &wg, &tickets, &mutex)
	go saleTickets19("窗口03", &wg, &tickets, &mutex)
	go saleTickets19("窗口04", &wg, &tickets, &mutex)

	wg.Wait()

	fmt.Println("main over...")
}

func saleTickets19(name string, wg *sync.WaitGroup, tickets *int, mutex *sync.Mutex)  {
	defer wg.Done()
	for {
		mutex.Lock()
		if *tickets > 0 {
			fmt.Printf("%s卖票 => %d\n", name, *tickets)
			*tickets-- //
		} else {
			fmt.Println("票已售罄...")
			mutex.Unlock()
			break
		}
		mutex.Unlock()
		//fatal error: all goroutines are asleep - deadlock!
	}
}
