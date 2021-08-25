package main

import "fmt"

// func main() {
// 	s := "Hello 월드"

// 	s2 := []rune(s)
// 	fmt.Println("len(s2)=", len(s2))
// 	for i := 0; i < len(s2); i++ {
// 		fmt.Print(s2[i], ", ")
// 	}
// }

func main() {
	s := "Hello 월드"

	s2 := []rune(s)
	fmt.Println("len(s2)=", len(s2))
	for i := 0; i < len(s2); i++ {
		fmt.Print(string(s2[i]), ", ")
	}
}
