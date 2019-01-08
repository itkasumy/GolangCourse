package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
		4.创建一个map，存储以下数据，并按照key的大小顺序打印输出。
			1:"王者荣耀"
			2:"绝地求生"
			3:"贪玩蓝月"
			4:"传奇"
			5:"红色警戒"
	*/

	m1701 := map[int]string{
		1: "王者荣耀",
		2: "绝地求生",
		3: "贪玩蓝月",
		4: "传奇",
		5: "红色警戒"}

	fmt.Println(m1701)

	s1701 := []int{}

	for key := range m1701 {
		s1701 = append(s1701, key)
	}
	fmt.Println(s1701)

	sort.Ints(s1701)

	fmt.Println(s1701)

	for _, val := range s1701 {
		fmt.Println(m1701[val])
	}
}
