package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
	z float64
}

func main() {
	//创建一个点：Point。属性：x坐标，y坐标，z坐标。创建2个点。打印坐标值。
	//提供一个方法，用于计算两点之间的距离：
	point01 := Point{1.2, 2.3, 3.4}
	fmt.Println(point01)
	point02 := Point{12, 50.2, 87}
	fmt.Println(point02)

	disp1p2 := getDistance(point01, point02)
	fmt.Println(disp1p2)
	fmt.Printf("%.2f\n", disp1p2)
}

func getDistance(p1, p2 Point) float64 {
	x := math.Pow((p1.x - p2.x), 2)
	y := math.Pow((p1.y - p2.y), 2)
	z := math.Pow((p1.z - p2.z), 2)
	distance := math.Sqrt(x + y + z)
	return distance
}
