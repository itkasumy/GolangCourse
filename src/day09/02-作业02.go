package main

import (
	"time"
	"math/rand"
	"fmt"
)

func sort02(s []float64) []float64 {
	for i := 0; i < len(s) - 1; i++ {
		for j := 0; j < len(s) - i - 1; j++ {
			if s[j] > s[j + 1] {
				s[j], s[j + 1] = s[j + 1], s[j]
			}
		}
	}
	return s
}

func createRandNum(startNum, endNum, length int) (s []float64) {
	for i := 0; i < length; i++ {
		tm02 := time.Now().UnixNano()
		rand.Seed(tm02)
		num01 := rand.Float64()
		res01 := num01 * float64(endNum - startNum) + float64(startNum)
		/*fmt.Printf("%.2f", res01)
		res01, err := strconv.ParseFloat(fmt.Sprintf("%.2f", res01), 10)
		if err != nil {
			return
		}*/
		s = append(s, res01)
		time.Sleep(1)
	}
	return
}

func main() {
	//创建一个23-67之间的随机浮点数切片，然后排序
	/*s01 := make([]float64, 0)
	//3-5
	//[0, 1) => [0, 10)
	//[0, 2) => [3, 5)
	//0-2 + 3 => r * 2 + 3
	// [0, 1) * 44 => [0, 44) => + 23  => [23, 67)
	for i := 0; i < 12; i++ {
		tm02 := time.Now().UnixNano()
		rand.Seed(tm02)
		num01 := rand.Float64()
		res01 := num01 * 44 + 23
		//fmt.Printf("%.2f", res01)
		s01 = append(s01, res01)
		time.Sleep(1)
	}
	fmt.Println(s01)*/
	//fmt.Println(createRandNum(3, 5, 3))
	s01 := createRandNum(23, 67, 12)
	fmt.Println(s01)
	s02 := sort02(s01)
	fmt.Println(s02)
}
