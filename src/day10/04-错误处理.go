package main

import (
	"fmt"
)

func main() {
	age, err := checkAge04(-20)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(age)
}

//创建一个检查年龄的函数
func checkAge04(nl int) (age int, er error) {
	var err error
	//fmt.Println(err) // nil
	if nl < 0 {
		err = fmt.Errorf("年龄是%d，年龄不能小于0", nl)
		return age, err
	}
	return nl, nil
}
