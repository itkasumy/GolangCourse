package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	fileName := "./day12/test.txt"
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%T\n", bs)
	fmt.Println(bs)
	fmt.Println(string(bs))
}
