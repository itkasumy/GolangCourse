package main

import (
"math"
"fmt"
)

type Shape11 interface {
	getArea() float64
	getPeri() float64
}

type triangle11 struct {
	a float64
	b float64
	c float64
}

type rectangle11 struct {
	length float64
	width float64
}

type circle11 struct {
	radius float64
}

func (tri triangle11) getArea() float64 {
	peri := (tri.a + tri.b + tri.c) / 2
	return math.Sqrt(peri * (peri - tri.a) * (peri - tri.b) * (peri - tri.c))
}

func (tri triangle11) getPeri() float64 {
	return tri.a + tri.b + tri.c
}

func (rect rectangle11) getArea() float64 {
	return rect.width * rect.length
}

func (rect rectangle11) getPeri() float64 {
	return 2 * (rect.width + rect.length)
}

func (c circle11) getArea() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle11) getPeri() float64 {
	return 2 * math.Pi * c.radius
}

func testInt11(s Shape11)  {
	fmt.Printf("面积是： %.2f, 周长是：%.2f\n", s.getArea(), s.getPeri())
}

func main() {
	tri01 := triangle11{3, 4, 5}
	rect01 := rectangle11{2, 3}
	c1 := circle11{2}

	testInt11(tri01)
	testInt11(rect01)
	testInt11(c1)
}
