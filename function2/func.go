package main

import "fmt"

func main() {
	a, b := fun1(2, 3)

	fmt.Println(a, b)
}

func fun1(x, y int) (int, int) {
	return y, x
}
