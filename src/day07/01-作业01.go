package main

import (
	"fmt"
	"math"
)

func main() {
	m0101 := map[string]int{"yi": 1, "er": 2, "san": 3}
	m0102 := copyMap(m0101)
	//fmt.Println(m0102)
	m0102["yi"] = 100
	fmt.Println(m0101, m0102)


	arr0101 := [5]int{3, 4, 2, 8, 1}
	arr0102 := arrSort(arr0101)
	fmt.Println(arr0101, arr0102)

	s0101 := []int{1, 3, 8, 19, 2, 3, 1, 3}
	num0101 := compareForBiggest(s0101)
	fmt.Println(num0101)

	//arr0103 := returnArr(arr0102)
	//fmt.Println(arr0103)

	arr0104 := reArrr(arr0102)
	fmt.Println(arr0104)

	fmt.Println(isSu(3))

	fmt.Println(totalDay(2016, 3, 1))

	fmt.Println(getJc(3))

	fmt.Println(getSXH())

	//fmt.Println(fibonaci(44))

	fmt.Println(math.Pow(2, 3))
	fmt.Println(math.Pow(3, 2))
}

//定义一个函数对于map集合进行拷贝
func copyMap(m map[string]int) map[string]int {
	mp := make(map[string]int)

	for key, val := range m {
		mp[key] = val
	}

	return mp
}

//封装一个对数组排序的函数

func arrSort(arr [5]int) [5]int {
	for i := 0; i < len(arr) - 1; i++ {
		for j := 0; j < len(arr) - i - 1; j++ {
			if arr[j] > arr[j + 1] {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
			}
		}
	}
	return arr
}

//求一组数中的最大值//求一组数中的最小值
func compareForBiggest(s []int) int {
	num := s[0]
	for i := 1; i < len(s); i++ {
		if num < s[i] {
			num = s[i]
		}
	}
	return num
}

//翻转数组，返回一个新数组
func returnArr(arr [5]int) [5]int {
	arrNew := [5]int{}
	for i := len(arr) - 1; i >= 0; i-- {
		//fmt.Println(arrNew)
		arrNew[len(arr) - 1 - i] = arr[i]
	}
	//fmt.Println(arrNew)
	return arrNew
}

func reArrr(arr [5]int) [5]int {
	for i := 0; i < len(arr) / 2; i++ {
		arr[i], arr[len(arr) - 1 - i] = arr[len(arr) - 1 - i], arr[i]
	}
	return arr
}

//判断一个数是否是质数
func isSu(num int) bool {
	for i := 2; i < num; i++ {
		if num == 2 {
			return true
		}
		if num % i == 0 {
			return false
		}
	}
	return true
}

//给定一个日期判断是一年中的第几天
func totalDay(year int, month int, day int) int {
	isRun := year % 4 == 0 && year % 100 != 0 || year % 400 == 0
	s := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if isRun {
		s[1] = 29
	}
	for i := 0; i < month - 1; i++ {
		day += s[i]
	}
	return day
}

//求阶乘  （累乘 5的阶乘 从1乘到5）
func getJc(n int) int {
	if n == 1 {
		return 1
	}
	return n * getJc(n - 1)
}

//求1!+2!+3!+....+n!
func getJcSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += getJc(i)
	}
	return sum
}

//求100-1000以内的水仙花数
//123 / 10 = 12 % 10 = 2
func getSXH() []int {
	s := []int{}
	for i := 100; i < 1000; i++ {
		//if (i % 10) * (i % 10) * (i % 10) + (i / 10 % 10) * (i / 10 % 10) * (i / 10 % 10) + (i / 100) * (i / 100) * (i / 100) == i {
		if int(math.Pow(float64(i % 10), 3) + math.Pow(float64(i / 10 % 10), 3) + math.Pow(float64(i / 100), 3)) == i {
			s = append(s, i)
		}
	}
	return s
}

//求兔子第n个月的对数(fibonaci斐波那契数列)1 1 2 3 5 8 13 21...
func fibonaci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fibonaci(n -1) + fibonaci(n - 2)
}