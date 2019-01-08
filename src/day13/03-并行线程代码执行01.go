package main

import (
	"fmt"
	"time"
)

func main() {

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
	}()

	time.Sleep(1)

	fmt.Println("main over...")
}
