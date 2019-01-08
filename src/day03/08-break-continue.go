package main

import "fmt"

func main() {
	/*for 表达式1; 表达式2; 表达式3 {
		循环体
	}*/
	// 正常情况下只有表达式2可以结束循环，循环结束的条件是：表达式2的结果是一个 false 值的时候
	// for 循环控制语句： break，continue
	// break： switch里面是结束某个case后面的代码，结束switch，在for循环里面是用作结束循环
	// continue: 作用，用于结束本次循环，进入下一轮循环循环条件改变表达式（表达式3）

	/*for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 5 {
			break
		}
		fmt.Println("皇家寻论学院")
		//break
	}*/

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
