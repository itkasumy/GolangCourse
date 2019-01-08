package main

import "fmt"

func main() {
	//数组/切片存数据：每个元素是一个单独的值，有序的数据集合，通过索引检索数据
	//map每个元素存的是一组键值对（映射关系），无序的键值集合，通过键检索数据
	//map里面的key是唯一的，如果存相同的key它会覆盖之前的key

	//map创建语法
	//var variable_name map[key_type]val_type
	var m0901 map[int]string
	fmt.Println(m0901)

	//字面量的方式定义map
	//var variable_name = map[key_type]val_type{}
	//var m0902 = map[string]int{}
	m0902 := map[string]int{"xueyuan": 111}
	fmt.Println(m0902)

	//var variable_name = make(map[key_type]val_type)
	m0903 := make(map[int]string)
	fmt.Println(m0903)

	/*m0903[1] = "xunlun"
	m0903[2] = "xueyuan"
	m0903[3] = "huangjia"
	fmt.Println(m0903)*/

	//m0901[100] = "Golang"
	//fmt.Println(m0901)
	//assignment to entry in nil map
	//nil: 空，表示没有
	// 对于map1我们只是声明了一个变量的类型，并没有为其分配内存地址，所以还不可以存数据
	fmt.Println(m0901 == nil, m0902 == nil, m0903 == nil)
	fmt.Printf("%p, %p, %p\n", m0901, m0902, m0903)
	/*var num0901 int
	fmt.Println(num0901)*/
	if m0901 == nil {
		m0901 = make(map[int]string)
	}
	m0901[100] = "Golang"
	fmt.Println(m0901)

	m0902["xunlun"] = 100
	fmt.Println(m0902)
}
