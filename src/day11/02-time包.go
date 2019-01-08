package main

import (
	"time"
	"fmt"
)

func main() {
	/*年
	天 365
	日 24
	时 60
	分 60
	秒 60
	毫秒 1000
	微妙 1000
	纳秒 1000*/
	t01 := time.Now()
	fmt.Printf("%T\n", t01)

	t02 := time.Date(2018, 10, 24, 9, 54, 23, 200300004, time.Local)
	fmt.Printf("%T\n", t02)
	fmt.Println(t02)
}
