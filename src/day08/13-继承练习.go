package main

import "fmt"

type animal struct {
	eat string
	leg int
}

type cat struct {
	animal
	catchMouse string
}

type dog struct {
	animal
	watchHome string
}

func main() {
	catee := cat{animal{"fish", 4}, "run for catching"}
	fmt.Println(catee)

	dogg := dog{animal{"boune", 4}, "sitting"}
	fmt.Println(dogg)
}

