package main

import "fmt"

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}

func lengthOfLIS(nums []int) int {
	// 以 key 为结尾的最大上升子序列的值为 value
	totalMax := 0
	maxMap := make(map[int]int)
	for i, num := range nums {
		max := 1
		for index := i - 1; index >= 0; index-- {
			if nums[index] < num && max < maxMap[index]+1 {
				max = maxMap[index] + 1
			}
		}
		maxMap[i] = max
		if totalMax < max {
			totalMax = max
		}
	}
	return totalMax
}

func lengthOfLISR(nums []int, index int) int {
	max := 1
	last := nums[index]
	for i := index - 1; i >= 0; i-- {
		if nums[i] < last {
			max++
			last = nums[i]
		}
	}
	fmt.Println(max)
	return max
}
