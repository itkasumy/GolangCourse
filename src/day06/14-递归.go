package main

import "fmt"

func main() {
	/*递归：就是一个函数调用自身*/

	//计算1-5的和
	/*5 + getSum1401() 1-4
		4 + getSum1401() 1-3
			3 + getSum1401() 1-2
				2 + getSum1401() 1
					return 1*/
	sum1401 := getSum1401(5)
	fmt.Println(sum1401)
}

func getSum1401(i int) int {
	if i == 1 {
		return 1
	}

	return i + getSum1401(i - 1)
}