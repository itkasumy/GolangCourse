package main

import (
	"fmt"
	"strconv"
)

func main() {
	//strconv包实现了基本数据类型和其字符串表示的相互转换。
	str01 := "123"
	//将纯数字字符串转整型数值
	num03, err := strconv.ParseInt(str01, 10, 64)
	//strconv.ParseInt: parsing "1233": value out of range
	//strconv.ParseInt: parsing "12 33": invalid syntax
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num03)

	str02 := "123.456"
	//将string转float
	num04, err := strconv.ParseFloat(str02, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num04)

	str03 := "321"
	num05, err := strconv.ParseUint(str03, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num05)

	str04 := "false"
	//将字符串转bool
	b01, err := strconv.ParseBool(str04)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T, %t\n", b01, b01)

	num06 := 123
	fmt.Printf("%T\n", num06)
	//将int64类型转换为string类型
	str05 := strconv.FormatInt(int64(num06), 10)
	fmt.Printf("%T\n", str05)
	fmt.Println(str05)

	//将floa转string
	num07 := 213.432
	str06 := strconv.FormatFloat(num07, 'f', 4, 64)
	fmt.Println(str06)

	//strconv.FormatBool() // 将bool转string


	/*num01 := 032
	num02 := 0xac
	fmt.Printf("%d, %d\n", num01, num02)*/


	//strconv.Itoa() //Itoa是FormatInt(i, 10) 的简写。将int转string
	num08 := 234
	str07 := strconv.Itoa(num08)
	fmt.Printf("%T\n", str07)
	fmt.Println(str07)

	//strconv.Atoi() Atoi是ParseInt(s, 10, 0)的简写。将string转int
	num09, err := strconv.Atoi(str07)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T\n", num09)
	fmt.Println(num09)
}
