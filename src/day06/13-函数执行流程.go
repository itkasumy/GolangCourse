package main

import "fmt"

func main() {
	fna()
}

func fna()  {
	fnb()
	fmt.Println("aaaaaaaaaaaa")
}

func fnb()  {
	fnc()
	fmt.Println("bbbbbbbbbbb")
}

func fnc()  {
	fmt.Println("ccccccccccccc")
}
