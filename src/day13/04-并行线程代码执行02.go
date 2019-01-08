package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		fmt.Println("我是子线程...")
		fmt.Println("我是子线程...")
		fmt.Println("我是子线程...")
		fmt.Println("我是子线程...")
		fmt.Println("我是子线程...")
		fmt.Println("我是子线程...")
		fmt.Println("我是子线程...")
		fmt.Println("我是子线程...")
		fmt.Println("我是子线程...")
		fmt.Println("我是子线程...")
		fmt.Println("我是子线程...")
		fmt.Println("我是子线程...")
		fmt.Println("我是子线程...")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("main over...")
}
