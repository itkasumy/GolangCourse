package main

import "fmt"

func addressStill(v *int)  {
	x := 456
	v = &x
}

func main() {
	x := 1000
	fmt.Println(x)
	addressStill(&x)
	fmt.Println("after x", x)
}
