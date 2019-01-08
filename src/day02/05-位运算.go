package main

import "fmt"

func main()  {
	/*
		位运算：&， | ， ^，<< ， >>
	*/
	num0501 := 12
	num0502 := 60

	res0501 := num0501 & num0502
	res0502 := num0501 | num0502
	res0503 := num0501 ^ num0502
	fmt.Println(res0501, res0502, res0503)

	num0503 := 5
	res0504 := num0503 << 2
	res0505 := num0503 >> 2
	fmt.Println(res0504, res0505)
}
