package main

import "fmt"

func main()  {
	/*
		关系运算符：>、 < 、 >=、<=、==、!=
		运算的结果是布尔值: true、false
	*/
	num0301 := 1
	num0302 := 2
	b0301 := num0301 > num0302
	fmt.Println(b0301)
	b0302 := num0302 > num0301
	fmt.Println(b0302)

	b0303 := num0301 < num0302
	fmt.Println(b0303)

	num0303 := 2
	b0304 := num0301 == num0302
	fmt.Println(b0304)
	b0305 := num0302 == num0303
	fmt.Println(b0305)
}
