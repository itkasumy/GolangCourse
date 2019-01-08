package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	//5. 创建一个23-67之间的随机浮点数切片，然后冒泡排序（5‘’）
	s := createRandFloatNum(5)
	fmt.Println(s)
	babelSort(s)
	fmt.Println(s)
}

func createRandFloatNum(n int) []float64 {
	s := make([]float64, 0)

	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(1)
		s = append(s, rand.Float64() * (67 - 23) + 23)
	}

	return s
}

func babelSort(s []float64) []float64 {
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j + 1] {
				s[j], s[j + 1] = s[j + 1], s[j]
			}
		}
	}
	return s
}
