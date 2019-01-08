package main

import "fmt"

type Employee struct {
	id int
	name string
	age int
	salary float64
}

type car01 struct {
	brand string
	tyre tyre01
}

type tyre01 struct {
	inner_radius float64
	outer_radius float64
}

//创建一个车结构体，提供属性：颜色，速度。。
type carr01 struct {
	color string
	speed float64
}

type byc01 struct {
	carr01
	int
}

type paoCar struct {
	carr01
	int
	oil string
}

func main() {
		//创建一个结构体Employee,设置字段id,name,age,salary，创建几个结构体对象，分别赋值，并打印
	goEg := Employee{
		name: "zhangsan",
		id: 1001,
		salary: 12300,
		age: 21,
	}

	javaEg := Employee{1002, "lisi", 24, 12000.111}

	fmt.Println(goEg)
	fmt.Println(javaEg)

	//定义一个结构体表示车轮：内半径，外半径。再定义一个结构体表示小汽车：车名，4个车轮。
	tyre0101 := tyre01{0.56, 0.72}
	bmw01 := car01{"宝马x3", tyre01{0.56, 0.66}}
	bmw02 := car01{"宝马x5", tyre0101}
	fmt.Println(tyre0101)
	fmt.Println(bmw01)
	fmt.Println(bmw02)

	//创建一个车结构体，提供属性：颜色，速度。。创建两个结构体：自行车，跑车，继承车的结构体。分别新增属性。创建对象，进行测试
	car0101 := carr01{"white", 0}
	byc0101 := byc01{carr01{"yellow", 20}, 2}
	lanB := paoCar{carr01{"purple", 500}, 4, "#93"}
	fmt.Println(car0101)
	fmt.Println(byc0101)
	fmt.Println(lanB)
}
