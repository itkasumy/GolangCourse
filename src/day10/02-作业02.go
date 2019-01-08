package main

import "fmt"

type Engine interface {
	start() int // 启动
	speedup() int // 加速
	stop() int // 停止
}

type YAMAHA struct {
	speed int
}

func (y YAMAHA) start() int {
	y.speed = 60
	return y.speed
}

func (y YAMAHA)  speedup() int {
	y.speed = 80
	return y.speed
}

func (y YAMAHA)  stop() int {
	y.speed = 0
	return y.speed
}

type HONDA struct {
	speed int
}

func (h HONDA) start() int {
	h.speed = 40
	return h.speed
}

func (h HONDA) speedup() int {
	h.speed = 120
	return h.speed
}

func (h HONDA) stop() int {
	h.speed = 0
	return h.speed
}

type car02 struct {
	e Engine
	speed int
}

func (c car02) testEngine() {
	c.speed = c.e.start()
	fmt.Printf("车子启动，启动速度是%d\n", c.speed)
	c.speed = c.e.speedup()
	fmt.Printf("车子行驶，行驶速度是%d\n", c.speed)
	c.speed = c.e.stop()
	fmt.Println("车子停止")
}

func main() {
	/*定义一个Engine接口，3个方法：start(),speedup(),stop()，表示启动，加速，停止
	定义两个结构体实现该接口：YAMAHA和HONDA
	实现方式分别是：

	YAMAHA
	启动：60，加速80，停止0

	HONDA
	启动：40，加速120，停止0

	定义一个Car结构体，将Engine作为字段，再提供一个car的方法：testEngine()，用于测试该小汽车的发动机。也就是在testEngine()中调用start(),speedup(),stop()方法。*/

	bentian := car02{YAMAHA{0}, 0}
	bentian.testEngine()

	fengtian := car02{HONDA{0}, 0}
	fengtian.testEngine()
}
