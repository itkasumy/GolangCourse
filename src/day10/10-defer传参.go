package main

import "fmt"

func main() {
	a := 1

	defer func() {
		fmt.Println(a, "----------------")
	}()

	defer func(b int) {
		fmt.Println(b, "=========")
	}(a)

	a = 100
	fmt.Println(a)
}
