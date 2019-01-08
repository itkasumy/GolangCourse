package main

import "fmt"

type USB08 interface {
	deviceName() string
	plugIn()
}

type mouse08 struct {
	brand string
	price float64
}

func (m mouse08) deviceName() string {
	fmt.Printf("鼠标的品牌是：%s, 鼠标的价格是：%.2f\n", m.brand, m.price)
	return m.brand
}

func (m mouse08) plugIn() {
	fmt.Println("鼠标插入到了usb中，可以正常工作了...")
}

type UDisk struct {
	brand string
	pric float64
}

func (u UDisk) deviceName() string {
	fmt.Println("我是：", u.brand, "U盘")
	return u.brand
}

func (u UDisk) plugIn() {

}

func main() {
	/*接口： 约定了功能的集合，只有方法的定义，而没有方法的实现
		谁实现了接口中的所有方法，谁就是该接口的实现类
		定义：
		type interface_var_name interface {
			method_name() [return_type]
			...
		}
		方法的实现：就是方法体里面的代码段
	*/

	m01 := mouse08{"罗技", 540.3}
	m01.deviceName()
	m01.plugIn()

	u01 := UDisk{"三星", 200}
	u01.deviceName()
}
