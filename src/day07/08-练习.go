package main

import "fmt"

func main()  {
	a := 10
	fmt.Println(a) // 10
	fmt.Println(&a)

	var p1 *int
	fmt.Println(p1)// nil

	p1 = &a
	fmt.Println(p1)
	fmt.Println(*p1)

	b := 20
	p1 = &b
	fmt.Println(p1)
	fmt.Println(*p1)

	*p1 = 100
	fmt.Println(p1)
	fmt.Println(*p1)

	fmt.Println(a)
	fmt.Println(b)
}
