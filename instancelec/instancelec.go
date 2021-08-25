package main

import "fmt"

type Student struct {
	name  string
	age   int
	grade int
}

// func main() {
// 	a := Student{"aaa", 21, 10}
// 	b := a
// 	b.age = 30

// 	fmt.Println(a)
// 	fmt.Println(b)
// }

func main() {
	a := Student{"aaa", 21, 10}
	var b *Student
	b = &a
	b.age = 30

	fmt.Println(a)
	fmt.Println(b)
}
