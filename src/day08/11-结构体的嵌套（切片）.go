package main

import "fmt"

type book struct {
	name string
	price float64
}

type person11 struct {
	name string
	age int
	books []book
}

func main() {
	book1 := book{"水浒传", 43.2}
	book2 := book{"三国演义", 73.7}
	book3 := book{"红楼梦", 23.99}
	book4 := book{"聊斋志异", 97.2}

	var hqw person11
	hqw.name = "hqw"
	hqw.age = 22
	hqw.books = []book{book1, book2, book3, book4}
	fmt.Println(hqw)
	fmt.Printf("姓名：%s,年龄：%d\n", hqw.name, hqw.age)
	for i, v := range hqw.books {
		fmt.Printf("\t第%d本书看的是：《%s》,书的价钱是：%.2f\n", i + 1, v.name, v.price)
	}

	fmt.Printf("%T\n", hqw.books)
	fmt.Println(hqw.books[3].name)
	fmt.Println(hqw.books[1].price)
}
