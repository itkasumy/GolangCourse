package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	fileName := "./day12/test.txt"
	/*bs := []byte{'X', 'u', 'n', 'L', 'u', 'n'}
	err := ioutil.WriteFile(fileName, bs, 0777)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("文件写入成功")*/

	str01 := "welcome to xun lun" //int()// string()// []byte()
	err := ioutil.WriteFile(fileName, []byte(str01), 0777)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("文件写入成功")

}
