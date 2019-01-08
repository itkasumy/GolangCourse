package main

import "fmt"

//空接口
type nilInt interface {

}

type person12 struct {
	name string
}

/*func (p person12) fn12()  {
	fmt.Println("-------------")
}*/

func main() {
	//任意类型都是空接口的实现
	arr01 := [4]nilInt{"hello", 123, true, []int{1, 2, 3}}
	fmt.Println(arr01)
}
