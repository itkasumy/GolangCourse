package main

import (
	"fmt"
	"time"
)

func main() {
	//func After(d Duration) <-chan Time
	/*ch01 := time.After(time.Second  * 3)
	tm01 := time.Now()
	fmt.Println(tm01)
	fmt.Println(<- ch01)*/

	tm01 := time.Now()
	fmt.Println(tm01)
	timer := time.NewTimer(3 * time.Second)

	fmt.Println(<- timer.C)


	fmt.Println("main over...")
}
