package main

import "fmt"

func main() {
	f1(10)
}

func f1(x int) {
	if x == 0 {
		return
	}
	fmt.Println(x)
	f1(x - 1)
}
