package main

import "fmt"

func main() {
	ch01 := make(chan int)

	go sendData09(ch01)

	for {
		data, ok := <- ch01
		if ok == false  {
			break
		}
		fmt.Printf("从子goroutine中取得数据%d, ok = %t\n", data, ok)
	}
}

func sendData09(ch chan int)  {
	for i := 1; i <= 10; i++ {
		ch <- i

		fmt.Printf("\t子goroutine传来第%d条数据\n", i)
	}
	close(ch)
}
