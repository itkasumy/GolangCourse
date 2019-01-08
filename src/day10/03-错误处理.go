package main

import (
	"fmt"
	"errors"
)

func main() {
	var err error
	fmt.Println(err)

	err = errors.New("错误测试...")
	fmt.Println(err)
	fmt.Printf("%T\n", err)

	age, err := checkAge03(-20)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(age)
}

//创建一个检查年龄的函数
func checkAge03(nl int) (age int, er error) {
	var err error
	age = nl
	if nl < 0 {
		err = errors.New("年龄不能小于0")
		return age, err
	}
	return nl, nil
}
