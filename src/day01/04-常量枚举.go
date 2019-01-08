package main

import "fmt"

func main() {
	const (
		SUN = 0
		MON = 1
		TUE = 2
		WED = 3
		THR = 4
		FRI = 5
		SAT = 6
	)

	const (
		cNUM0401 = 1
		cNUM0402 = 2
		cNUM0403
		cNUM0404
		cNUM0405 = 5
		cNUM0406
	)
	fmt.Println(cNUM0401, cNUM0402, cNUM0403, cNUM0404, cNUM0405, cNUM0406)

	const (
		cA0401 = iota
		cA0402 = iota
		cA0403 = iota
		cA0404 = iota
	)
	fmt.Println(cA0401, cA0402, cA0403, cA0404)

	const (
		cB0401 = iota
		cB0402
		cB0403
		cB0404
	)
	fmt.Println(cB0401, cB0402, cB0403, cB0404)

	// 在每一个const关键字出现时，被重置为0，然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1
	// 常量组中如不指定类型和初始化值，则与上一行非空常量右值相同
	const (
		cC0401 = iota  // 0
		cC0402 // 1
		cC0403 = "hello" // hello
		cC0404 // hello
		cC0405 = iota // 4
		cC0406 // 5
	)
	fmt.Println(cC0401, cC0402, cC0403, cC0404, cC0405, cC0406)

}
