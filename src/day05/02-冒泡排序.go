package main

import "fmt"

func main() {
	arr0201 := [8]int{2, 3, 5, 9, 6, 5, 1, 7}
	fmt.Println(arr0201)

	for i := 0; i < len(arr0201) - 1; i++ {
		for j := 0; j < len(arr0201) - i - 1; j++ {
			if arr0201[j] > arr0201[j + 1] {
				arr0201[j], arr0201[j + 1] = arr0201[j + 1], arr0201[j]
			}
		}
	}

	fmt.Println(arr0201)
}
