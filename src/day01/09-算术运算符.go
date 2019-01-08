package main

import "fmt"

func main()  {
	//算术运算符
	// +、-、*、/、%、++、--
	var num0901 = 4
	var num0902 = 6

	var sum0901 = num0901 + num0902
	var sub0901 = num0901 - num0902
	var mul0901 = num0901 * num0902
	var div0901 = num0901 / num0902
	var mod0901 = num0901 % num0902

	fmt.Println("sum0901 = ", sum0901)
	fmt.Println("sub0901 = ", sub0901)
	fmt.Println("mul0901 = ", mul0901)
	fmt.Println("div0901 = ", div0901)
	fmt.Println("mod0901 = ", mod0901)

	var num0903 = 10
	fmt.Println(num0903)
	num0903++ //num0903 = num0903 + 1 /*+=*/
	fmt.Println(num0903)
	num0903++
	fmt.Println(num0903)

	var num0904 = 10
	num0904--
	fmt.Println(num0904)

}
