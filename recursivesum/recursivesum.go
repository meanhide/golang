package main

import (
	"fmt"
)

// 재귀호출 이용
// func main() {

// 	s := sum(10, 0)
// 	fmt.Println(s)
// }

// func sum(x, s int) int {
// 	if x == 0 {
// 		return s
// 	}
// 	s += x
// 	return sum(x-1, s)
// }

func main() {
	s := 0
	for i := 1; i <= 10; i++ {
		s += i
	}
	fmt.Println(s)
}
