package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		for k := 5; k > i+2; k-- {
			fmt.Print(" ")
		}
		for j := 0; j < i*2+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for j := 0; j < 3; j++ {
		for l := 0; l < j+1; l++ {
			fmt.Print(" ")
		}
		for m := 3; m > j*2-2; m-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
