package main

import "fmt"

func main() {
	//操场上有一群人，人数在100到200之间。三人一组多1人，四人一组多2人，五人一组多3人。问操场上共有多少人
	for i := 100; i <= 200; i++ {
		if i % 3 == 1 && i % 4 == 2 && i % 5 == 3 {
			fmt.Println(i)
		}
	}

	//两个自然数x，y相除，商3余10，被除数，除数，商，余数的和是163，求被除数，除数
	for i := 1; i < 150; i++ {
		if i / (150 - i) == 3 && i % (150 - i) == 10 {
			fmt.Println(i)
		}
	}

	//一个四位数，恰好等于去掉它的首位数字之后所剩的三位数的3倍，这个四位数是多少。
	for i := 1000; i < 10000; i++ {
		if i == 3 * (i % 1000) {
			fmt.Println(i)
		}
	}

	//打印出1-500所有的自然数中不包含4的自然数
	/*24 / 1 = 24  % 10 = 4
	54 / 1 =54 % 10 = 4
	104 / 1 = 104 % 10 = 4

	45 / 10 = 4 % 10 = 4
	46 / 10 = 4 % 10 = 4
	147 / 10 = 14 % 10 = 4
	345 / 10 = 34 % 10 = 4

	429 / 100 = 4 % 10 = 4
	409 / 100 = 4 % 10 = 4
	423 / 100 = 4 % 10 = 4*/
	for i := 1; i <= 500; i++ {
		//个位不为4而且十位不为4而且百位不为4
		// 147
		//i % 10 != 4 && i / 10 % 10 != 4 && i / 100 != 4
		//个位不为4或者十位不为4或者百位不为4
		// 147
		//i % 10 != 4 || i / 10 % 10 != 4 || i / 100 != 4
		/*if i % 10 != 4 || i / 10 % 10 != 4 || i / 100 != 4 {
			fmt.Print(i, "\t")
		}*/
		if i % 10 != 4 && i / 10 % 10 != 4 && i / 100 != 4 {
			fmt.Print(i, "\t")
		}
	}
}
