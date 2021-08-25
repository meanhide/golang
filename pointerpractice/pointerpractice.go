package main

import "fmt"

func main() {
	var a int
	var p *int

	p = &a
	a = 3

	fmt.Println(p)
	fmt.Println(a)
	fmt.Println(*p)
}
