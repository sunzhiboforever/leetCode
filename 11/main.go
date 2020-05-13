package main

import "fmt"

func main() {
	a := maxArea([]int{2, 3, 10, 5, 7, 8, 9})
	fmt.Println(a)
}

func maxArea(height []int) int {
	start := 0
	end := len(height) - 1
	max := 0

	for start < end {
		h := 0
		if height[start] < height[end] {
			h = height[start]
		} else {
			h = height[end]
		}
		if h*(end-start) > max {
			max = h * (end - start)
		}
		if height[start] < height[end] {
			start++
		} else {
			end--
		}
	}
	return max
}
