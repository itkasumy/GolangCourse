package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	num1 := rand.Intn(100)
	fmt.Println(num1)

	//[0, 1) 0-1但不包含1
	num2 := rand.Float64()
	fmt.Println(num2)

	t1 := time.Now()
	ts := t1.UnixNano()
	fmt.Println(ts)
	/*fmt.Println(t1)
	ts := t1.Unix()
	fmt.Println(ts)
	ts1 := t1.Unix()
	fmt.Println(ts1)
	ts2 := t1.UnixNano()
	fmt.Println(ts2)
	ts3 := t1.UnixNano()
	fmt.Println(ts3)*/

	zzdx := rand.NewSource(1) // 1. 根据指定的int类型作为要给种子数，返回一个资源对象
	rdx := rand.New(zzdx) // 2. 根据资源对象获取一个随机对象
	num3 := rdx.Intn(100) // 3. 用随机对象获取指定范围内的随机数
	fmt.Println(num3)

	zzdx1 := rand.NewSource(time.Now().UnixNano())
	rdx1 := rand.New(zzdx1)
	num4 := rdx1.Intn(100)
	fmt.Println(num4)
}
