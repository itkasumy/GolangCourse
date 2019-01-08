package main

import (
	"time"
	"fmt"
)

func main() {
	t02 := time.Date(2018, 10, 24, 9, 54, 23, 200300004, time.Local)
	fmt.Printf("%T\n", t02)
	fmt.Println(t02)
	//时间格式化
	fmt.Println(t02.Format(time.ANSIC))
	fmt.Println(t02.Format(time.RFC3339Nano))
	fmt.Println(t02.Format("Mon Jan 2 15:04:05 -0700 MST 2006")) // 1-2-3-4-5 6
	fmt.Println(t02.Format("2006-1-2 15:04:05")) // 6-1-2 3:4:5
	fmt.Println(t02.Format("2006年1月2日 15时04分05秒")) // 6-1-2 3:4:5

	str01 := "2018 10月24 10时34分05秒"
	t03, err := time.Parse("2006 1月2 15时04分05秒", str01)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t03)

	t04 := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(t04)
	num01 := t04.UnixNano()
	fmt.Println(num01)
}
