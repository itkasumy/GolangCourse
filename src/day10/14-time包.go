package main

import (
	"time"
	"fmt"
)

func main() {
	//Now返回当前本地时间。
	t01 := time.Now()
	fmt.Printf("%T\n", t01) // Time
	fmt.Println(t01)

	//获取指定日期的年份
	fmt.Println(t01.Year())
	//获取指定日期的月份
	fmt.Println(t01.Month())
	//获取指定日期的日期
	fmt.Println(t01.Day())
	//获取指定日期的小时
	fmt.Println(t01.Hour())
	//获取指定日期的分钟
	fmt.Println(t01.Minute())
	//获取指定日期的秒钟
	fmt.Println(t01.Second())
	//获取指定日期的纳秒时间
	fmt.Println(t01.Nanosecond())
	//获取指定日期的年份和周数
	fmt.Println(t01.ISOWeek())
	//获取指定日期是那一年的第几天
	fmt.Println(t01.YearDay())
	//获取指定日期是周几
	fmt.Println(t01.Weekday())
	//返回指定日期的时分秒
	fmt.Println(t01.Clock())
	h, m, s := t01.Clock()
	fmt.Printf("今天是：%d:%d:%d\n", h, m, s)

	y, M, d := t01.Date()
	fmt.Printf("%T\n", M)
	fmt.Printf("今天是：%d-%v-%d\n", y, M, d)

}
