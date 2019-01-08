package main

import "fmt"

type Engine interface {
	start(string)
	speedup(string)
	stop(string)
}

type YAMAHA struct {

}

func (y YAMAHA) start(brand string) {
	fmt.Println(brand, "使用雅马哈引擎启动速度：", 60)
}

func (y YAMAHA) speedup(brand string) {
	fmt.Println(brand, "使用雅马哈引擎加速：", 80)
}

func (y YAMAHA) stop(brand string) {
	fmt.Println(brand, "使用雅马哈引擎停止：", 0)
}

type HONDA struct {

}

func (y HONDA) start(brand string) {
	fmt.Println(brand, "使用宏达引擎启动速度：", 40)
}

func (y HONDA) speedup(brand string) {
	fmt.Println(brand, "使用宏达引擎加速：", 120)
}

func (y HONDA) stop(brand string) {
	fmt.Println(brand, "使用宏达引擎停止：", 0)
}

type Car struct {
	eg Engine
	name string
}

func (c Car) testEngine() {
	c.eg.start(c.name)
	c.eg.speedup(c.name)
	c.eg.stop(c.name)
}

func main() {
	/*
	9. 综合题：（20‘’）

	定义一个Engine接口，3个方法：start(),speedup(),stop()，表示启动，加速，停止
	定义两个结构体实现该接口：YAMAHA和HONDA
	实现方式分别是：

	YAMAHA
	启动：60，加速80，停止0

	HONDA
	启动：40，加速120，停止0

	定义一个Car结构体，将Engine作为字段，再提供一个car的方法：testEngine()，用于测试该小汽车的发动机。也就是在testEngine()中调用start(),speedup(),stop()方法
*/
	bentian := Car{HONDA{}, "本田"}
	bentian.testEngine()

	mob := Car{YAMAHA{}, "摩拜"}
	mob.testEngine()

}
