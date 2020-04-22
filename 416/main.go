package main

import "fmt"

func main() {
	fmt.Println(canPartition([]int{1, 2, 3, 5}))
}

func canPartition(nums []int) bool {
	// 求出总合的一半
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	cache := make(map[int]bool)
	return pack(cache, nums, sum/2, len(nums)-1)
}

func pack(cache map[int]bool, nums []int, size, i int) bool {
	if _, ok := cache[size]; ok {
		return cache[size]
	}
	if size == 0 {
		return true
	}
	if i < 0 {
		return false
	}

	cache[size] = pack(cache, nums, size-nums[i], i-1) || pack(cache, nums, size, i-1)
	return cache[size]
}
