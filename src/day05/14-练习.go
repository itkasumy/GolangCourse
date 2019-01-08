package main

import "fmt"

func main() {
	/*
	1.创建一个map用于存储一个人的信息：
		name：张三丰
		age：108
		sex：男性
		address："湖北省十堰市武当山"
	2.打印遍历该map
	3.修改sex为女性
	*/
	m1401 := make(map[string]string)
	m1401["name"] = "张三丰"
	m1401["age"] = "108"
	m1401["sex"] = "男"
	m1401["address"] = "湖北省十堰市武当山"

	for key, val := range m1401{
		fmt.Println(key, "=>", val)
	}

	if _, ok := m1401["sex"]; ok {
		m1401["sex"] = "女"
	}
	fmt.Println(m1401)

}
