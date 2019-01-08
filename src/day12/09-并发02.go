package main

import (
	"fmt"
	"time"
)

func main() {
	go fn09()
	fmt.Println("我是主程序,正在执行中...")

	time.Sleep(time.Second)
	fmt.Println("main over....")
}

func fn09()  {
	fmt.Println("我是子程序...")
}
