package main

import (
	"fmt"
	"time"
)

func main() {
	//1.主线程中打印100个数字，子线程中打印100个数字
	go prtNum11()

	for i := 0; i < 100; i++ {
		fmt.Printf("我是主线程中的i = %d\n", i)
	}
	time.Sleep(time.Second)

	fmt.Println("main over...")
}

func prtNum11()  {
	for j := 0; j < 100; j++ {
		fmt.Printf("\t我是子线程中的j = %d\n", j)
	}
}
