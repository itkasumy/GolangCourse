package main

import "fmt"

func main() {
	//1. 打印数字1-5
	//for i := 1; i <= 5; i++ {
	//	fmt.Println(i)
	//}
	//fmt.Println(i)

	//2. 计算1-10的和
	//1 + 2 + 3 + ... + 9 + 10
	sum0601 := 0
	for i := 1; i <= 10; i++ {
		//sum0601 = sum0601 + i
		sum0601 += i
		//fmt.Println(i)
	}
	fmt.Println(sum0601)

	//求1*2*3...*20
	//求100以内所有奇数的和
	mul0601 := 1
	for i := 1; i <= 20; i++ {
		//mul0601 = mul0601 * i
		mul0601 *= i
	}
	fmt.Println(mul0601)

	//奇数： 1 3 5 7 9...
	sum0602 := 0
	for i := 1; i <= 100; i++ {
		//fmt.Println(i)
		if i % 2 == 1 {
			sum0602 += i
		}
	}
	fmt.Println(sum0602)

	//求100以内,能被3整除，但是不能被5整除的数
	for i := 1; i <= 100; i++ {
		//fmt.Println(i)
		if i % 3 == 0 && i % 5 != 0 {
			fmt.Println(i)
		}
	}

	//求100以内,能被3整除，但是不能被5整除的数,每行打印5个
	counter := 0
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 && i % 5 != 0 {
			fmt.Print(i, "\t")
			counter++
			if counter % 5 == 0 {
				fmt.Println()
			}
		}
	}
}
