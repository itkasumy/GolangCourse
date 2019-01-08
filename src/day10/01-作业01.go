package main

import "fmt"

type car01 struct {
	color string
	speed float64
}

func (c car01) move01()  {
	fmt.Println("车子在移动")
}

func (c car01) stop01()  {
	fmt.Println("车子停止")
}

type byc01 struct {
	car01
	int
}

type roadSter struct {
	car01
	int
}

func (b byc01) bellRing()  {
	fmt.Println("叮叮叮...")
}

func (b byc01) move01() {
	fmt.Println("摩拜单车越来越好用了...")
}

func (r roadSter) piaoYi()  {
	fmt.Println("正在漂移...")
}

func main() {
	//创建一个车结构体，提供属性：颜色，速度。方法：移动()。停止()。创建两个结构体：自行车，跑车，继承车的结构体。分别新增属性和方法。创建对象，进行测试
	mobi := byc01{car01{"orange", 20}, 2}
	mobi.move01()
	mobi.stop01()
	mobi.bellRing()

	lan := roadSter{car01{"purple", 470}, 4}
	lan.piaoYi()
	lan.move01()
	lan.stop01()
}
