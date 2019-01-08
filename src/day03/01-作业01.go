package main

import "fmt"

func main() {
	//模拟登录，键盘上输入用户名和密码，如果用户名是admin密码是123，或者用户名是zhangsan密码是zhangsan123，都表示可以登录。否则打印登录失败
	name0101 := ""
	code0101 := ""
	fmt.Println("请输入您的用户名和密码：")
	fmt.Scanln(&name0101, &code0101)
	/*if name0101 == "admin" && code0101 == "123" || name0101 == "zhangsan" && code0101 == "zhangsan123" {
		fmt.Println("恭喜你，登陆成功")
	} else {
		fmt.Println("登陆失败")
	}*/

	if name0101 == "admin" && code0101 == "123" {
		fmt.Println("恭喜你，登陆成功")
	} else if name0101 == "zhangsan" && code0101 == "zhangsan123" {
		fmt.Println("恭喜你，登陆成功")
	} else {
		fmt.Println("登陆失败")
	}
}
