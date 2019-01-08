package main

import (
"math"
"fmt"
)

type Shape09 interface {
	getArea() float64
	getPeri() float64
}

type triangle09 struct {
	a float64
	b float64
	c float64
}

type rectangle09 struct {
	length float64
	width float64
}

type circle09 struct {
	radius float64
}

func (tri triangle09) getArea() float64 {
	peri := (tri.a + tri.b + tri.c) / 2
	return math.Sqrt(peri * (peri - tri.a) * (peri - tri.b) * (peri - tri.c))
}

func (tri triangle09) getPeri() float64 {
	return tri.a + tri.b + tri.c
}

func (rect rectangle09) getArea() float64 {
	return rect.width * rect.length
}

func (rect rectangle09) getPeri() float64 {
	return 2 * (rect.width + rect.length)
}

func (c circle09) getArea() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle09) getPeri() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	/*定义一个接口：形状
	定义两个方法：
	周长()
	面积()
	定义三个结构体分别实现该接口： 三角形、	矩形、圆形
	在主函数中：分别创建三种形状的对象，并调用对应的周长和面积的方法*/
	tri01 := triangle09{3, 4, 5}
	fmt.Println(tri01.getArea())
	fmt.Println(tri01.getPeri())

	rect01 := rectangle09{2, 3}
	fmt.Println(rect01.getArea())
	fmt.Println(rect01.getPeri())

	c1 := circle09{2}
	fmt.Println(c1.getArea())
	fmt.Println(c1.getPeri())
}
