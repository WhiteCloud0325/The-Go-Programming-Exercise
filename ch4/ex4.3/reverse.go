package main

import "fmt"

func reverse(nums *[5]int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func main() {
	nums := [5]int{1, 2, 3, 4, 5}
	reverse(&nums)
	fmt.Println(nums)
}
