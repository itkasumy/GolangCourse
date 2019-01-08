package main

import (
	"fmt"
	"sort"
)

func main()  {
	//格式化输出以下数据
	a := 100
	b := 3.14
	c := "Hello"
	d := `我是一个

小小的

石头`

	g := "我是一个" +
		"小小的石头"
	
	e := true
	f := 'A'


	fmt.Printf("prototype of a is %v, type of a is %T\n", a, a)
	fmt.Printf("prototype of b is %v, type of b is %T\n", b, b)
	fmt.Printf("prototype of c is %v, type of c is %T\n", c, c)
	fmt.Printf("prototype of d is %v, type of d is %T\n", d, d)
	fmt.Printf("prototype of e is %v, type of e is %T\n", e, e)
	fmt.Printf("prototype of f is %v, type of f is %T\n", f, f)

}


