package main

import "fmt"

func main() {
	/*
	return的作用：
		1. 返回函数的返回值:如果函数有返回值就必须要返回结果
		2. 强制结束函数运行
	*/
	//fn0901()
	//fn0902()
	//fmt.Println("hello")
	//fn0903()
	fn0904(1)
}

func fn0901() int {
	fmt.Println("xunlun")
	return 3
}

func fn0902()  {
	fmt.Println("xunlun")
	return// 在这里强行结束了函数，后面的代码都不会被执行
	fmt.Println("xueyuan")
}

func fn0903()  {
	for i := 0; i < 10; i++ {
		if i == 5 {
			//break
			return
		}
		fmt.Println(i)
	}
}

func fn0904(a int) int {
	if a == 1 {
		return 1
	}
	/*else {
		return 2
	}*/
	return 2
}
