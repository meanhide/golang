package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	temp := [5]int{}
	fmt.Println("arr:", arr)

	for i := 0; i < len(arr); i++ {
		temp[i] = arr[len(arr)-1-i]
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = temp[i]
	}
	fmt.Println("temp:", temp)
	fmt.Println("arr:", arr)
}
