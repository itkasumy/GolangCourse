package main

import "fmt"

func main() {
	fn0702()
	fn0703()
}

func fn0701()  {
	fmt.Println("这是程序01")
}

func fn0702()  {
	for i := 0; i < 10; i++ {
		fmt.Println("程序2运行中...")
	}
	fn0701()
}

func fn0703()  {
	fmt.Println("程序3执行中...")
	fn0701()
}
