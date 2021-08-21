package main

import "fmt"

// func main() {
// 	var a []int

// 	fmt.Printf("len(a) = %d\n", len(a))
// 	fmt.Printf("cap(a) = %d\n", cap(a))
// }

// func main() {
// 	a := []int{1, 2, 3, 4, 5}

// 	fmt.Printf("len(a) = %d\n", len(a))
// 	fmt.Printf("cap(a) = %d\n", cap(a))
// }

// func main() {
// 	a := make([]int, 0, 8)

// 	fmt.Printf("len(a) = %d\n", len(a))
// 	fmt.Printf("cap(a) = %d\n", cap(a))

// 	a = append(a, 1)

// 	fmt.Printf("len(a) = %d\n", len(a))
// 	fmt.Printf("cap(a) = %d\n", cap(a))
// }

func RemoveBack(a []int) ([]int, int) {
	return a[:len(a)-1], a[len(a)-1]

}

func RemoveFront(a []int) ([]int, int) {
	return a[1:], a[0]
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var back int
	for i := 0; i < 5; i++ {
		a, back = RemoveBack(a)
		fmt.Printf("%d,", back)
	}
	fmt.Println()
	fmt.Println(a)
}
