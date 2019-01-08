package main

import "fmt"

func fna11()  {
	fmt.Println("aaaaaaaaaaa")
}

func fnb11()  {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg)
			fmt.Println("ooooooooooooo")
		}
	}()

	for i := 0; i < 10; i++ {
		if i == 5 {
			panic("程序出现异常...")
		}
		fmt.Println("bbbbbbbbb")
	}
}

func fnc11()  {
	fmt.Println("ccccccccc")
}

func main() {
	// panic() //受到惊吓
	// defer // 延迟
	// recover//恢复 让程序恢复到没有出现错误之前的状态继续往下执行
	fna11()
	fnb11()
	fnc11()
}
