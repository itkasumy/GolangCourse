package main

import (
	"fmt"
	"strings"
)

func main() {
	str1901 := "xunlun"
	fmt.Println(str1901)

	//Contains
	fmt.Println(strings.Contains(str1901, "un"))
	fmt.Println(strings.Contains(str1901, "nu"))

	//ContainsAny子串里面只要有一个字符在str1里面就返回true
	fmt.Println(strings.ContainsAny(str1901, "xyz"))

	//Count统计出现了多少次给定的子串
	fmt.Println(strings.Count(str1901, "un"))

	//HasPrefix判断是否以指定子串开头
	fmt.Println(strings.HasPrefix(str1901, "xun"))

	//HasSuffix判断是否以指定子串结尾
	fmt.Println(strings.HasSuffix(str1901, "lun"))

	//Index统计从字符串开头第一次出现子串的位置索引
	fmt.Println(strings.Index(str1901, "un"))

	//LastIndex统计从字符串结尾第一次出现子串的位置索引
	fmt.Println(strings.LastIndex(str1901, "un"))

	//Join 将给定的字符串切片拼接成一个字符串（以给定的字符串作为分割）
	s1902 := []string{"xunlun", "xunyuan", "huangjia"}
	fmt.Println(strings.Join(s1902, "-"))

	//Repeat 重复给定字符串多少次
	fmt.Println(strings.Repeat("ha", 6))

	// Split以指定的字符分隔成一个切片数组
	s1903 := strings.Join(s1902, "-")
	fmt.Println(s1903)
	fmt.Println(strings.Split(s1903, "-"))
	slice1901 := strings.Split(s1903, "-")
	fmt.Println(slice1901)

	//ToUpper 将小写字符转大写
	str1902 := "abcd"
	fmt.Println(strings.ToUpper(str1902))

	//ToLower将大写字母转小写
	str1903 := "XYZ"
	fmt.Println(strings.ToLower(str1903))

	//Trim 去除指定首尾字符
	str1904 := "   xun lun   xye    "
	fmt.Println("***", str1904, "***")
	fmt.Println("***", strings.Trim(str1904, " "), "***")

	//Replace 将指定字符，以指定的字符替换
	fmt.Println(strings.Replace(str1904, "un", "xxx", -1))

}
