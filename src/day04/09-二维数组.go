package main

import "fmt"

func main() {
	var arr0901 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr0901)

	var arr0902 = [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(arr0902)
	fmt.Println(arr0902[0])
	fmt.Println(arr0902[0][1])

	arr0902[1][2] = 100
	fmt.Println(arr0902)

	fmt.Println(len(arr0902))
	fmt.Println(len(arr0902[0]))

	for i := 0; i < len(arr0902); i++ {
		for j := 0; j < len(arr0902[i]); j++ {
			fmt.Println(arr0902[i][j])
		}
	}

}
