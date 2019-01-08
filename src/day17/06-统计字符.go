package main

import "fmt"

func main() {
	//6. 统计该字符串中每个字母出现的次数"ECALIYHWEQAEFSZCVZTWEYXCPIURVCSWTDBCIOYXGTEGDTUMJHUMBJKHFGUKNKN"（5‘’）

	str := "ECALIYHWEQAEFSZCVZTWEYXCPIURVCSWTDBCIOYXGTEGDTUMJHUMBJKHFGUKNKN"

	mp := make(map[int32]int)

	for _, char := range str {
		//fmt.Printf("%T\n", char)
		if _, ok := mp[char]; ok {
			mp[char]++
		} else {
			mp[char] = 1
		}
	}

	for key, val := range mp {
		fmt.Printf("%c: %d\n", key, val)
	}
}
