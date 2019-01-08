package main

import "fmt"

func main() {
	//2. 百元百鸡，一百元钱买100只鸡，公鸡5元一只，母鸡3元一只，小鸡1元3个。

	// 如果只买公鸡 =》 20
	// 如果只买母鸡 =》 33只
	// 如果只买小鸡 =》 300只
	for i := 0; i <= 20; i++ {
		for j := 0; j <= 33; j++ {
			k := 100 - i - j
			if 5 * i + 3 * j + k / 3 == 100 {
				fmt.Println(i, j, k)
			}
		}
	}
}
