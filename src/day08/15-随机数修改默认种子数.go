package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	//[0, 1) * [0, 10) [3, 8)
	rand.Seed(time.Now().UnixNano())
	num1 := rand.Intn(100)
	num2 := rand.Float64()
	num3 := num2 * 10
	num4 := num2 * 5 //[0, 5)
	num5 := num4 + 3 // [3, 8)
	fmt.Println(num1, num2, num3, num4, num5)

	fmt.Println(int(rand.Float64() * float64(200 - 27) + float64(27)))
}
