package main

import (
	"fmt"
	"math"
)

func main() {
	//求绝对值
	fmt.Println(math.Abs(2))

	//求两个数中较大的那个
	fmt.Println(math.Max(2.1, 1.99999))
	//求两个数中较大的那个
	fmt.Println(math.Min(2.1, 1.99999))

	//表示向上取整
	fmt.Println(math.Ceil(3.1))
	//表示向下取整
	fmt.Println(math.Floor(9.9))
	//四舍五入
	fmt.Println(math.Round(4.4))

	//分别取整数和小数部分
	fmt.Println(math.Modf(3.14))

	//取余数
	fmt.Println(math.Mod(7, 4))
	//求平方根
	fmt.Println(math.Sqrt(100))
	//开立方根
	fmt.Println(math.Cbrt(64))
	//求冥
	fmt.Println(math.Pow(2, 3))
	//10为底数的冥
	fmt.Println(math.Pow10(5))

	//pi的值
	fmt.Println(math.Pi)
}
