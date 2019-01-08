package main

import (
	"errors"
	"fmt"
)

func devi05(a, b int) (div int, err error) {
	if b == 0 {
		//如果b==0，程序异常，需要返回异常信息，后面的代码不再执行
		err = errors.New("除数为0了，除数不能为零")
		return 0, err
	}
	return a / b, err
}

func main() {
	var a, b int
	fmt.Println("请输入一个数字：")
	fmt.Scanln(&a)
	fmt.Println("请再输入一个数字：")
	fmt.Scanln(&b)
	res01, err := devi05(a, b)
	if err != nil {
		fmt.Println(err)
		return // 两个作用，1. 返回函数结果2.终止函数
	}
	fmt.Printf("%d / %d = %d\n", a, b, res01)
}
