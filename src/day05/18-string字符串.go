package main

import "fmt"

func main() {
	//len()计算的是字节长度，而不是字符长度
	var str1801 string = "xunlun"
	fmt.Println(str1801)
	fmt.Println(len(str1801))

	str1802 := "xunlun学院"
	fmt.Println(str1802)
	fmt.Println(len(str1802))

	/*var char1801 byte = 'A'
	fmt.Println(char1801)*/

	/*for i := 0; i < len(str1801); i++ {
		fmt.Println(str1801[i])
	}*/

	/*for _, val := range str1801 {
		fmt.Println(val)
	}*/

	for _, val := range str1801 {
		fmt.Printf("%c\t", val)
	}
	fmt.Println()
}
