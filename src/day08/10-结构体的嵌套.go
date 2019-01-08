package main

import "fmt"

/*type book struct {
	name string
	price float64
}*/

type person10 struct {
	name string
	age int
	book1 book
}

func main() {
	//聚合关系： has a
	book1001 := book{"Java从入门到放弃", 30.4}
	fmt.Println(book1001)

	zhangsan := person10{"zhangsan", 18, book1001}
	fmt.Println(zhangsan)

	lisi := person10{
		name: "lisi",
		age: 19,
		book1: book1001,
	}
	fmt.Println(lisi)

	fmt.Printf("%T\n", lisi.book1)
	fmt.Println(lisi.book1.name, lisi.book1.price)
}
