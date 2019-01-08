package main

import "fmt"

func main()  {
	//交换2个变量的值
	num0101 := 1
	num0102 := 2
	fmt.Printf("原始值: num0101 = %d, num0102 = %d\n", num0101, num0102)

	/*var num0103 int
	num0103 = num0101
	num0101 = num0102
	num0102 = num0103
	fmt.Printf("交换后: num0101 = %d, num0102 = %d\n", num0101, num0102)*/

	//方式二:
	num0101 = num0101 + num0102 // num1 <= num1 + num2(1 + 2)  => num1此时等于两数之和，num2不变
	num0102 = num0101 - num0102 // 3 - 2 => num2(1) 相当于原来的num1了， 此时num1没变化，依然是两数之和（3）
	num0101 = num0101 - num0102 // num1是两数之和 减去已经变成之前num1的值的num2 , num1 减去之前的num1 就相当于得到之前的num2
	fmt.Printf("交换后: num0101 = %d, num0102 = %d\n", num0101, num0102)
}
