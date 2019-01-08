package main

import (
	"fmt"
	"strings"
)

func main() {
	//str1 = "hello hello hello world"
	//1. 统计字符串的长度:
	//2. 统计"llo"出现的次数:
	//3. 统计"wor"出现的位置:
	str0101 := "hello hello hello world"
	fmt.Println(len(str0101))
	fmt.Println(strings.Count(str0101, "llo"))
	fmt.Println(strings.Index(str0101, "wor"))

	//统计该字符串中每个字母出现的次数"ECALIYHWEQAEFSZCVZTWEYXCPIURVCSWTDBCIOYXGTEGDTUMJHUMBJKHFGUKNKN"
	str0102 := "ECALIYHWEQAEFSZCVZTWEYXCPIURVCSWTDBCIOYXGTEGDTUMJHUMBJKHFGUKNKN"
	m0101 := make(map[int32]int)

	//fmt.Println(str0102[0])
	//index, 'E'
	for _, val := range str0102 {
		//if _, ok := m0101[val]; ok == false { // !ok
		if _, ok := m0101[val]; !ok {
			m0101[val] = 1
		} else {
			m0101[val]++
		}
	}
	fmt.Println(m0101)

	for key, val := range m0101 {
		fmt.Printf("字母%c出现了%d次\n", key, val)
	}

	//将字符串"aa:zhangsan@163.com!bb:lisi@sina.com!cc:wangwu@126.com"，按照指定的key-value规则存入map中
	//key：aa,bb,cc
	//value:zhangsan@163.com,lisi@sina.com,wangwu@126.com

	str0103 := "aa:zhangsan@163.com!bb:lisi@sina.com!cc:wangwu@126.com"

	//fmt.Println(strings.Split(str0103, "!"))
	s0101 := strings.Split(str0103, "!")
	//fmt.Println(s0101)

	m0102 := make(map[string]string)
	for _, val := range s0101 {
		//fmt.Println(key, val)
		//fmt.Println(val)
		//fmt.Println(strings.Split(val, ":"))
		s0102 := strings.Split(val, ":")
		//fmt.Println(s0102)
		m0102[s0102[0]] = s0102[1]
		//m0102[strings.Split(val, ":")[0]] = strings.Split(val, ":")[1]
	}

	fmt.Println(m0102)
}
