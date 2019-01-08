package main

import "fmt"

func main() {
	/*排序：
	从小到大：升序
	从大到小：降序*/
	arr0701 := [5]int{6, 3, 7, 4, 1}
	fmt.Println(arr0701)

	for i := 0; i < len(arr0701) - 1; i++ {
		for j := 0; j < len(arr0701) - i - 1; j++ {
			if arr0701[j] > arr0701[j + 1] {
				arr0701[j], arr0701[j + 1] = arr0701[j + 1], arr0701[j]
			}
		}
	}
	fmt.Println(arr0701)

	/*for i := 1; i < len(arr0701); i++ {
		for j := 1; j < len(arr0701) - i + 1 ; j++ {
			if arr0701[j - 1] > arr0701[j] {
				arr0701[j - 1], arr0701[j] = arr0701[j], arr0701[j - 1]
			}
		}
	}
	fmt.Println(arr0701)*/
}
