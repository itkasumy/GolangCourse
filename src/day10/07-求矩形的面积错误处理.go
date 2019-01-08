package main

import "fmt"

type rect07 struct {
	msg string
	length float64
	width float64
}

func (r *rect07) Error() string {
	return r.msg
}

func getRectArea(r rect07) (area float64, err error) {
	errMsg := ""
	if r.length <= 0 {
		errMsg = fmt.Sprintf("%f矩形的长度不能小于等于0", r.length)
	}
	if r.width <= 0 {
		if errMsg == "" {
			errMsg = "矩形的宽度不能小于等于0"
		} else {
			errMsg += ",矩形的宽度也不能小于等于0"
		}
	}

	if errMsg != "" {
		return 0, &rect07{errMsg, r.length, r.width}
	}

	area = r.width * r.length
	return area, nil
}

func main() {
	rect01 := rect07{"", -2, 0}
	res01, err := getRectArea(rect01)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res01)
}
