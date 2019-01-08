package main

import "fmt"

func main()  {
	/*
		声明变量： var 变量名 数据类型 = 值
		变量的组成： 变量名 数据类型 数据值 内存地址值（用一个符号获取： &）
	*/
	/*键盘上输入以下内容并打印输出：

	用户名，username

	密码，password*/

	fmt.Println("请输入您的姓名和年龄:")
	//var name0701 string
	var age0701 int
	/*fmt.Scanln(&name0701, &age0701) // 阻塞式的
	fmt.Println(name0701, age0701)*/

	var age0702 int
	fmt.Scanf("%d-%d", &age0701, &age0702)
	fmt.Println(age0701, age0702)

}
