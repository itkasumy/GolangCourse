package main

import (
	"fmt"
	"time"
)

/*func main() {
	for i := 1; i <= 5; i++ { // i
		go func(n int) {
			// n = 1
			fmt.Println(n)
		}(i)

		//defer fmt.Println(i)
		//time.Sleep(1)
	}

	time.Sleep(time.Second)

	fmt.Println("main over...")
}*/

func main() {
	for i := 1; i <= 5; i++ { // i
		i := i
		fmt.Println(&i)
		go func() {
			fmt.Println(i)
		}()
	}

	/*{
		n := 1
	}

	{
		n := 1
	}*/

	time.Sleep(time.Second)

	fmt.Println("main over...")
}

