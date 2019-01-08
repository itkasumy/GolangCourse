package main

type Student struct {
	name string
}

func main() {
	// map的值是结构体类型，不能通过访问键给结构体赋值
	m := map[string]Student{"people": {"zhoujielun"}}
	m["people"] = Student{"wuyanzu"}
	//m["people"].name = "wuyanzu"
	// cannot assign to struct field m["people"].name in map
}
