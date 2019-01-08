package main

import "fmt"

func main()  {
	num0201 := 7656
	geWei := num0201 % 10 // = 765 ... 6
	shiWei := num0201 / 10 % 10 // => 765 % 10 = 76 ... 5
	baiWei := num0201 / 100 % 10 // => 76 % 10 = 7 ... 6
	qianwei := num0201 / 1000 % 10 // => 7 % 10 = 0 ... 7
	fmt.Println(qianwei, baiWei, shiWei, geWei)
}
