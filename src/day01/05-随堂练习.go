package main

import "fmt"

func main() {
	// 在每一个const关键字出现时，被重置为0，然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1
	// 常量组中如不指定类型和初始化值，则与上一行非空常量右值相同
	const (
		cA0501 = iota   // 0
		cA0502          // 1
		cA0503          // 2
		cA0504 = "iota" // iota
		cA0505          // iota
		cA0506 = 100    // 100
		cA0507          // 100
		cA0508 = iota   // 7
		cA0509          // 8
	)

	const (
		cA0510 = iota // 0
	)

	fmt.Println(cA0501, cA0502, cA0503, cA0504, cA0505, cA0506, cA0507, cA0508, cA0509, cA0510)

	const (
		cB0501 = 'A'  // 65 a => 97
		cB0502        // 65
		cB0503 = iota // 2
		cB0504        // 3
	)

	fmt.Println(cB0501, cB0502, cB0503, cB0504)

	var ling = '0' // 0 "0"
	fmt.Printf("%v\n", ling)
	var guo = '国'
	fmt.Printf("%v\n", guo)
}
