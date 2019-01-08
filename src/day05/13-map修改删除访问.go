package main

import "fmt"

func main() {
	m1301 := map[int]string{1: "Golang", 2: "Java", 3: "Python"}
	fmt.Println(m1301)

	//访问
	//fmt.Println(m1301[1])
	//fmt.Println(m1301[4])

	/*val1301, ok := m1301[4]
	fmt.Println(val1301, ok)

	val1302, y1301 := m1301[2]
	fmt.Println(val1302, y1301)*/

	if val, ok := m1301[3]; ok {
		fmt.Println(val)
	}

	if val, y1302 := m1301[100]; y1302 {
		fmt.Println(val)
	}

	//修改
	if _, y1303 := m1301[2]; y1303 {
		m1301[2] = "JavaScript"
	}

	//删除
	fmt.Println(m1301)
	delete(m1301, 3)
	fmt.Println(m1301)
	delete(m1301, 2)
	fmt.Println(m1301)
}
