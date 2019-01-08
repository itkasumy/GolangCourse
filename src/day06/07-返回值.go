package main

import "fmt"

func main() {
	//返回值：
	//retrun: 将return后面的结果返回给函数调用处
	/*func 函数名(参数列表) (返回值列表) {
		函数体
	}*/
	//(返回值名字1 返回值类型1, 返回值名字2 返回值类型2...)
	//return :用于函数体内，返回我们需要的返回值，结果返回到函数调用的地方
/*
	返回值注意事项：
		1. 如果返回值列表要求要有结果返回，我们必须返回结果，而且返回的结果要和返回值列表严格匹配（类型、数量、顺序）
		2. 如果返回值列表没有取名字，也可以，返回的结果严格匹配返回值列表的数量，类型和顺序就可以了
		3. 如果每个返回值列表都有名字，我们可以只写一个return就可以了
*/
	//计算1-10的和，把结果*2打印输出
	//fmt.Println(getSum0701())
	//res0701 := getSum0701()
	//fmt.Println(res0701 * 2)

	//fmt.Println(getSum0702())
	//fmt.Println(getSum0703())
	//fmt.Println(fn0701())
	//
	//fmt.Println(fn0702())

	//求矩形的周长和面积
	length := 4.0
	width := 3.0
	zc, area := getRectangle(length, width)
	fmt.Printf("长为%f, 宽为%f的矩形的周长是%f,面积是%f\n", length, width, zc, area)
	_, mj := getRectangle(5, 6.6)
	fmt.Println(mj)
}

func getSum0701() (sum int) {
	for i := 1; i <= 10; i++ {
		sum += i
	}
	return sum // return 返回
	//return ""
	//return 1, 3
}

func getSum0702() int {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	return sum
}

func getSum0703() (sum int) {
	for i := 1; i <= 10; i++ {
		sum += i
	}
	return
}

func fn0701() (a, b int) {
	a = 1
	b = 2
	//return a, b
	return
}

func fn0702() (int, string) {
	a := 1
	str := "hello"
	return a, str
}

func getRectangle(length float64, width float64) (zc float64, area float64) {
	zc = 2 * (length + width)
	area = length * width
	//return zc, area
	return
}
