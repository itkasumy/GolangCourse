package main

import "fmt"

func main() {
	// 不要通过共享内存操作数据，而要通过通道传递消息
	go fn08()
	fmt.Println("我是主程序,正在执行中...")
	fmt.Println("main over....")
}

func fn08()  {
	fmt.Println("我是子程序...")
}
