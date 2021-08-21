package main

import "fmt"

func main() {
	x := 33

	switch x {
	case 31:
		fmt.Println("x = 31")
	case 32:
		fmt.Println("x = 32")
	case 33:
		fmt.Println("x = 33")
	}
}
