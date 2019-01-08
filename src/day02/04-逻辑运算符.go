package main

import "fmt"

func main()  {
	/*
		逻辑运算符： && 、|| 、！
		1. &&: 只要有一个是假（false），结果就是假（false），两个全为真(true)，结果才为真（true）(一假全假，全真则真)
		2. ||: 只要有一个为真（true）, 结果就是真(true), 两个全为假（false），结果才为假（false）（一真则真，全假才假）
		3. !: 原来为假则变真，原来为真则变假
	*/
	b0401 := true
	b0402 := false
	b0403 := true
	b0404 := false

	res0401 := b0401 && b0402 // true && false => false
	fmt.Println(res0401)
	res0402 := b0401 && b0403
	fmt.Println(res0402)

	res0403 := b0401 || b0402
	fmt.Println(res0403)
	res0404 := b0401 || b0403
	fmt.Println(res0404)
	res0405 := b0402 || b0404
	fmt.Println(res0405)

	res0406 := !b0401
	fmt.Println(res0406)
	res0407 := !b0402
	fmt.Println(res0407)
}
