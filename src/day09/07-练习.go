package main

import (
	"fmt"
	"math"
)

type human07 struct {
	name string
	age int
}

type stu07 struct {
	human07
	shool string
}

func (h human07) eat07() {
	fmt.Println("父亲吃窝窝头...")
}

func (h human07) prtInfo07() {
	fmt.Printf("姓名：%s, 年龄：%d\n", h.name, h.age)
}

func (s stu07) eat07() {
	fmt.Println("学生吃食堂...")
}

type Point07 struct {
	x float64
	y float64
	z float64
}

func (p1 Point07) calcDis07 (p2 Point07) float64 {
	/*x := p2.x - p1.x
	y := p2.y - p1.y
	z := p2.z - p1.z

	sum := math.Pow(x, 2) + math.Pow(y, 2) + math.Pow(z, 2)

	res := math.Sqrt(sum)

	return res*/

	return math.Sqrt(math.Pow(p2.x - p1.x, 2) + math.Pow(p2.y - p1.y, 2) + math.Pow(p2.z - p1.z, 2))
}

func main() {
	/*定义一个人的父类： 属性是：name age
	eat(), prtInfo()
定义一个子类学生：添加属性：shool
重写方法 eat()

2.创建一个点：Point。属性：x坐标，y坐标，z坐标。创建2个点。打印坐标值。
提供一个方法，用于计算两点之间的距离*/

	per01 := human07{"张三", 43}
	fmt.Println(per01)
	per01.eat07()
	per01.prtInfo07()

	stu01 := stu07{human07{"lisi", 17}, "xunlun"}
	fmt.Println(stu01)
	stu01.eat07() // 如果子类没有该方法，就会调用父类的方法，如果子类有自己的方法，就调用自己的
	stu01.prtInfo07() // 自己没有该方法，调用的是父类的

	a := Point07{1, 1, 1}
	b := Point07{2, 2, 2}
	res := a.calcDis07(b)
	fmt.Println(res)
}
