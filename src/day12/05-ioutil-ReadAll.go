package main

import (
	"strings"
	"io/ioutil"
	"fmt"
)

func main() {
	str01 := "welcome to xunlun"

	r := strings.NewReader(str01) // *Reader
	bs, err := ioutil.ReadAll(r) // io.Reader=>Read() (n int, err error)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)
	fmt.Println(string(bs))
}
