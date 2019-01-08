package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	s3 := []int{}
	/*s1 := s[0: 2]
	s2 := s[3: 5]
	fmt.Println(s1, s2)
	s3 = append(s3, s1...)
	s3 = append(s3, s2...)*/
	s3 = append(s3, s[0:2]...)
	s3 = append(s3, s[3:5]...)
	fmt.Println(s3)

}

