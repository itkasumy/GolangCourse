package main

import "fmt"

type person01 struct {
	name string
	age int
}

type car04 struct {
	color string
	int
}

//定义一个方法
func (p person01) prtInfo()  {
	fmt.Printf("姓名是： %s, 年龄是： %d\n", p.name, p.age)
}

func (p person01) study() {
	fmt.Println("我在寻论学习")
}

func (c car04) prtInfo() {
	fmt.Printf("颜色是： %s, 轮胎数是： %d\n", c.color, c.int)
}

func main() {
	/*定义：
		func (var_str_name str_type) method_name (argms...) (returns...) {
			方法体
		}
	如果要调用一个方法，必须实力化一个对象
	所有给定类型的方法属于该类型的方法集
	*/
	zhangsan := person01{"zhangsan", 27}
	zhangsan.prtInfo()

	bmw := car04{"white", 4}
	bmw.prtInfo()
}
