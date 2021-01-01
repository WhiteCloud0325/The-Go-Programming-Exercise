package main

import "fmt"

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func min(vals ...int) int {
	minVal := 0
	for _, val := range vals {
		if minVal > val {
			minVal = val
		}
	}

	return minVal
}

func max(vals ...int) int {
	maxVal := 0
	for _, val := range vals {
		if maxVal < val {
			maxVal = val
		}
	}

	return maxVal
}

func main() {
	fmt.Println(sum())
	fmt.Println(min())
	fmt.Println(max())
}
