package main

import "fmt"

func main()  {
	/*
		强类型语言运算的时候类型必须一致
		=> 类型转换
				语法： T(变量名)
	*/
	var num0801 = 3
	var num0802 = 2.5 // 2
	var num0803 = int(num0802)
	var res0801 = num0801 + num0803 // invalid operation: num0801 + num0802 (mismatched types int and float64)
	fmt.Println(res0801)

	var num0804 = float64(num0801)
	var res0802 = num0804 + num0802
	fmt.Println(res0802)
}
