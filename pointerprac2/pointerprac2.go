package main

import "fmt"

func main() {
	var a int = 1
	Increase(&a)

	fmt.Println(a)

}

func Increase(x *int) {
	*x = *x + 1
}
