package main

import "fmt"

func main() {
	/*
		数据类型：简单类型、复合类型
		简单类型： 数值型、布尔型、字符串型
	*/

	// 布尔类型: true、false
	var b0601 bool = true
	fmt.Printf("value of b0601 is %t, type of b0601 is %T\n", b0601, b0601)
	var b0602 = false
	fmt.Printf("value of b0602 is %t, type of b0602 is %T\n", b0602, b0602)

	//数值型
	// 整形和浮点型
	// 浮点型（float32、float64）
	var height float64 = 1.75
	fmt.Printf("type of height is %T, value of height is %f\n", height, height)
	var weight float32 = 120
	fmt.Printf("type of weight is %T, value of weight is %f\n", weight, weight)

	// 整型(有符号型、无符号型)
	//int8 [-128 - 127] => uint8 [0-255](byte)
	//int16 [-32768 - 32767] => uint16 [0 - 65535]
	// int32(rune) => uint32
	// int64 => uint64
	// uint int(根据个人电脑系统的位数而定)
	var num0601 int8 = 127
	fmt.Println(num0601)
	fmt.Printf("type of num0601 is %T, value of num0601 is %d\n", num0601, num0601)

	var num0602 int32 = 123
	fmt.Printf("type of num0602 is %T, value of num0602 is %d\n", num0602, num0602)

	var num0603 = 1
	fmt.Printf("type of num0603 is %T, value of num0603 is %d\n", num0603, num0603)

	var str0601 string = "我是一个小小的石头"
	fmt.Printf("type of str0601 is %T, value of str0601 is %s\n", str0601, str0601)
	var str0602 = `疯狂的
							石头`
	fmt.Printf("type of str0602 is %T, value of str0602 is %s\n", str0602, str0602)
}
