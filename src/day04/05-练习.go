package main

import "fmt"

func main() {
	//创建一个int类型的数组，存储4个数。分别进行赋值，并打印结果
	var arr0501 [4]int
	arr0501[0] = 1
	arr0501[1] = 2
	arr0501[2] = 3
	arr0501[3] = 4
	fmt.Println(arr0501)

	//创建一个float64类型的数组，存储2个小数。(array)
	//var arr0502 [2]float64
	var arr0502 = [2]float64{1.2, 3.4}
	fmt.Println(arr0502)

	//创建一个String类型的数组，存储3个字符串
	arr0503 := [3]string{"a", "b", "c"}
	fmt.Println(arr0503)

	// 给定一个数组，arr1 := [10] int{5,4,3,7,10,2,9,8,6,1}，求数组中所有数据的总和。
	arr0504 := [10] int{5, 4, 3, 7, 10, 2, 9, 8, 6, 1}
	sum0501 := 0
	/*for i := 0; i < len(arr0504); i++ {
		fmt.Println(arr0504[i])
		sum0501 += arr0504[i]
	}*/
	for _, val := range arr0504 {
		fmt.Println(val)
		sum0501 += val
	}
	fmt.Println(sum0501)


}
