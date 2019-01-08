package main

import "fmt"

func main()  {
	var num0701 int = 3
	fmt.Printf("%T\n", num0701)
	var str0701 string = "xunlun"
	fmt.Printf("%T\n", str0701)
	//num0701 = str0701 // cannot use str0701 (type string) as type int in assignment
	var num0702 int = 5

	fmt.Println("赋值之前", num0701, num0702)
	num0701 = num0702
	fmt.Println("赋值之后", num0701, num0702)


	var arr0701 = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", arr0701)
	var arr0702 = [3]int{1, 2, 3}
	fmt.Printf("%T\n", arr0702)

	var arr0703 = [3]string{"xunlun", "Golang", "Java"}
	fmt.Printf("%T\n", arr0703)

	//arr0701 = arr0702 // cannot use arr0702 (type [3]int) as type [5]int in assignment

	arr0704 := arr0702
	fmt.Println(arr0704)
	fmt.Printf("%T\n", arr0704)
	arr0704[0] = 100
	fmt.Println(arr0702, arr0704)
	fmt.Printf("%p, %p\n", &arr0702, &arr0704)

}
