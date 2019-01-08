package main

import (
	"fmt"
	"time"
)

func main() {
	//7. 给定一个日期判断是一年中的第几天（5‘’）

	//方式1
	totalDay := whichDay(2018, 2, 7)
	fmt.Println(totalDay)

	//方式二
	allDays := statisticsDays(2018, 2, 7)
	fmt.Println(allDays)
}

func whichDay(year, month, day int) int {
	monthes := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if isRun(year) {
		monthes[1] = 29
	}

	for i := 0; i < month - 1; i++ {
		day += monthes[i]
	}
	return day
}

func isRun(year int) bool {
	if year % 4 == 0 && year % 100 != 0 || year % 400 == 0 {
		return true
	}
	return false
}

func statisticsDays(year int, mon time.Month, day int) int {
	t := time.Date(year, mon, day, 0, 0, 0, 0, time.Local)
	return t.YearDay()
}
