package main

import "fmt"

type car05 struct {
	name string
	speed float64
}

func (c car05) run()  {
	fmt.Printf("%s 的最高速度是： %.2f\n", c.name, c.speed)
}

func main() {
	//1定义一个car的结构体，字段有名字，速度。定义一个方法，run(),打印该车的速度
	carByc := car05{"摩拜单车", 30}
	carByc.run()

	bmw := car05{"宝马", 340.12}
	bmw.run()
}
