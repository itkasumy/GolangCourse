package main

import (
	"fmt"
	"time"
)

func main() {
	/*子线程1，每睡眠250毫秒打印数字，一共打印5个。1-5
	子线程2，每睡眠400毫秒打印字母，一共打印5个。A-E
	主线程中，启动两个子线程，然后进入睡眠3000毫秒。
	观察程序运行结果*/
	go fn1201()
	go fn1202()

	time.Sleep(time.Millisecond * 3000)

	fmt.Println("main over...")
}

func fn1201()  {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Millisecond * 250)
		fmt.Println("子线程1打印数字i =", i)
	}
}

func fn1202()  {
	for i := 65; i <= 69; i++ {
		time.Sleep(time.Millisecond * 400)
		fmt.Printf("子线程2打印字母i = %c\n", i)
	}
}
