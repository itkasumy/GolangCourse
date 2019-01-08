package main

import "fmt"

type Run10 interface {
	run()
}

type Runner struct {
	name string
}

func (r Runner) run()  {
	fmt.Printf("跑步爱好者 %s 围着湖边每晚一圈...\n", r.name)
}

type dog10 struct {
	name string
	int
}

func (d dog10) run() {
	fmt.Printf("%d条腿的小狗%s围着湖边跑一圈..\n", d.int, d.name)
}

func testInterface10(r Run10)  {
	r.run()
}

func main() {
	/*多态：通过接口来体现*/
	r01 := Runner{"xiaoming"}
	//r01.run()

	d01 := dog10{"xiaohei", 4}
	//d01.run()

	testInterface10(r01)
	testInterface10(d01)
}
