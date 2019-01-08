package main

import "fmt"

func main() {
	str01 := "abcdefgh"

	for _, v := range str01 {
		fmt.Printf("%T\n", v)
		defer fmt.Printf("%q", v)
	}
	fmt.Println()

	char01 := 'a' // int32
	fmt.Printf("%T\n", char01)
}
