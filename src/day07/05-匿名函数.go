package main

import "fmt"

func main() {
	/*匿名函数：没有名字的函数*/
	//函数不能嵌套
	fmt.Println(getSum0501(1, 2))

	//立即执行函数（IIFE）
	func () {
		fmt.Println("xunlunxuey")
	}()

	//这个没有立即执行，只是赋值给了一个 变量
	fn1 := func(a int, b int) int {
		return a + b
	}

	var fn2 func(int, int) int
	fn2 = func(a int, b int) int {
		return a + b
	}

	res1 := fn2(100, 200)
	res2 := fn1(1000, 2000)
	fmt.Println(res1, res2)

	fmt.Printf("%T, %T\n", fn1, fn2)

	for i := 0; i < 5; i++ {
		func() {
			fmt.Println(i)
		}()
	}

	fmt.Printf("%T\n", fn0501)

	res3 := fn0501(1, 2, func(m int, n int) int {
		return m + n
	})
	fmt.Println(res3)
}

func getSum0501(a, b int) int {
	return a + b
}

//func () {}

func fn0501(a int, b int, fn func (int, int) int) int {
	return fn(a, b)
}
