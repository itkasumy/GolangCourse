package main

import "fmt"

func main() {
	ch01 := make(chan int, 3)

	go sendData12(ch01)

	for {
		v, ok := <- ch01
		if !ok {
			fmt.Println("数据读取完毕！！！")
			break
		}

		fmt.Println("\t\t从通道中读取到的数据是:", v)
	}

	fmt.Println("main over...")

}

func sendData12(ch chan int)  {
	for i := 1; i <= 1000; i++ {
		fmt.Println("第", i, "条数据存到chan中...")
		ch <- i
	}
	close(ch)
}
