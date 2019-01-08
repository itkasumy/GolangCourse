package main

import "fmt"

type ani06 struct {
	name string
	age int
}

func (ani ani06) eat06() {
	fmt.Printf("动物:%s在吃东西\n", ani.name)
}

func (ani ani06) prtInfo06()  {
	fmt.Printf("动物:%s,年龄是:%d\n", ani.name, ani.age)
}

type cat06 struct {
	ani06
	skill string
}

type dog06 struct {
	ani06
	skill string
}

func (c cat06) showSkill06()  {
	fmt.Println("猫的技能是", c.skill)
}

func (c cat06) eat06()  {
	fmt.Println("猫吃鱼猫吃鱼猫吃鱼猫吃鱼猫吃鱼猫吃鱼猫吃鱼猫吃鱼猫吃鱼猫吃鱼猫吃鱼")
}

func main() {
	/*继承的好处：
		避免了重复的代码
		子类可以定义自己独有的放
		子类可以复写父类的方法，如果复写，子类对象会调用复写的方法
	*/
	rabit := ani06{"玉兔", 1200}
	rabit.eat06()
	rabit.prtInfo06()

	tom := cat06{ani06{"tom", 3}, "catch mouch"}
	tom.eat06()
	tom.prtInfo06()
	tom.showSkill06()

	dogg := dog06{ani06{"xiaohei", 2}, "watch home"}
	dogg.eat06()
	dogg.prtInfo06()
}

/*定义一个人的父类： 属性是：name age
	eat(), prtInfo()
定义一个子类学生：添加属性：shool
重写方法 eat()

2.创建一个点：Point。属性：x坐标，y坐标，z坐标。创建2个点。打印坐标值。
提供一个方法，用于计算两点之间的距离*/
