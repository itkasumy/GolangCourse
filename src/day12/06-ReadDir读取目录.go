package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	/*fileName := "H:"
	infos, err := ioutil.ReadDir(fileName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T\n", infos)
	fmt.Println(infos)
	for _, info := range infos {
		fmt.Println(info.Name())
		fmt.Println(info.Size())
		fmt.Println(info.IsDir())
	}*/
	readDir06("H:\\02-Golang")
}

func readDir06(fileName string) {
	dirInfos, err := ioutil.ReadDir(fileName)
	if err != nil {
		fmt.Println(err)
	}

	for _, fileInfo := range dirInfos {
		if fileInfo.IsDir() { // true == true => true
			fmt.Println("这是一个目录，目录的名字是：", fileInfo.Name())
			//"H:\02-Golang"
			//"H:"
			fileName := fileName + "\\" + fileInfo.Name()
			readDir06(fileName)
		} else {
			fmt.Println("\t这是一个文件，文件的名字是：", fileInfo.Name())
		}
	}
}
