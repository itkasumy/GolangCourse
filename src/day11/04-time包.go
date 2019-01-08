package main

import (
	"time"
	"fmt"
)

func main() {
	t01 := time.Date(2018, 10, 24, 10, 0, 0, 0, time.Local)
	fmt.Println(t01)
	//以纳秒为单位增加时间
	t02 := t01.Add(12 * 60 * 1000 * 1000 * 1000)
	fmt.Println(t02)
	//以年月日为单位加时间
	t03 := t01.AddDate(1, 0, 8)
	fmt.Println(t03)

	//计算距离另一个时间的差值，单位时纳秒
	t04 := t01.Sub(t02)
	fmt.Println(t04)

	b01 := t01.After(t02)
	b01 = t02.After(t01)
	fmt.Println(b01)
	b01 = t02.Before(t01)
	fmt.Println(b01)
}
