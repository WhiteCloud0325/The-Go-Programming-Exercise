package main

import "fmt"

func main() {
	var slice1 []int
	slice2 := make([]int, 0, 5)
	slice1 = append(slice1, 1)
	fmt.Println(len(slice1))
	fmt.Println(len(slice2))
}
