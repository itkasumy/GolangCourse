package main

import "fmt"

func main() {
	m1201 := make(map[int]string)
	fmt.Println(m1201)

	m1201[1] = "Golang"
	m1201[2] = "Java"
	m1201[3] = "C/C++/C#"
	m1201[4] = "Python"
	m1201[5] = "JavaScript"
	fmt.Println(m1201)

	/*fmt.Println(m1201[1])
	fmt.Println(m1201[2])
	fmt.Println(m1201[3])
	fmt.Println(m1201[4])
	fmt.Println(m1201[5])

	for i := 0; i < len(m1201); i++ {}*/

	for key, val := range m1201{
		fmt.Println(key, "=>", val)
	}
}
