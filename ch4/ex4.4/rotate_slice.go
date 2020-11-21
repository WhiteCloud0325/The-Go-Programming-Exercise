package main

import "fmt"

func rotate(nums []int) {
	if len(nums) == 0 {
		return
	}
	first := nums[0]
	copy(nums, nums[1:])
	nums[len(nums)-1] = first
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	rotate(s)
	fmt.Println(s)
}
