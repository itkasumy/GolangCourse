package main

import "fmt"

func main() {
	//赋值的数据个数不能超过数组的长度

	//方式1
	var arr0601 [5]int
	fmt.Println(arr0601)

	//方式2
	var arr0602 = [5]int{}
	fmt.Println(arr0602)

	//方式3
	var arr0603 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr0603)

	//方式4
	var arr0604 = [5]int{1, 3, 5}
	fmt.Println(arr0604)

	//方式5
	arr0605 := [5]int{}
	fmt.Println(arr0605)

	//方式6 给数组的指定位置赋值
	var arr0606 = [5]int{2: 3, 3: 4}
	fmt.Println(arr0606)
	arr0607 := [5]string{3: "Golang", 4: "xunlun"}
	fmt.Println(arr0607)
	fmt.Println(arr0607[0])
	fmt.Println(arr0607[3])

	//方式7
	var arr0608 = [...]int{1, 3, 5}
	fmt.Println(arr0608)

	//方式8
	var arr0609 = [...]int{0: 1, 2: 100, 9: 1000}
	fmt.Println(arr0609)
	fmt.Println(len(arr0609))
}
