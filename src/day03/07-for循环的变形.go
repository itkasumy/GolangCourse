package main

import "fmt"

func main() {
	//1. 变形1，省略初始化变量
	/*i := 0
	for ; i < 5; i++ {
		fmt.Println("皇家寻论学院")
	}
	fmt.Println(i)*/

	//2. 变形2，省略判断条件=>进入死循环
	//for i := 0; ; i++ {
	//	fmt.Println(i)
	//}

	//3. 变形3，省略条件改变表达式
	/*for i := 0; i < 5; {
		fmt.Println("皇家寻论学院")
		i++
	}*/
	// 把改变条件的表达式写在里面，它是循环体的部分，执行顺序也不一样

	//4. 变形4，全部省略
	/*for{
		fmt.Println("皇家寻论学院")
	}*/

	//5. 变形5，省略1和3
	i := 0
	for i < 5 {
		fmt.Println("皇家寻论")
		i++
	}
}
