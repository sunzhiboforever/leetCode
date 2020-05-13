package main

import (
	"fmt"
	"math"
)

func main() {
	r := combinationSum4([]int{1, 2, 3}, 4)
	fmt.Println(r)
}

func combinationSum4(nums []int, target int) int {
	cache := make(map[int]int)
	return combinationSum4R(cache, nums, target)
}

func combinationSum4R(cache map[int]int, nums []int, target int) int {
	if _, ok := cache[target]; ok {
		return cache[target]
	}
	if target < 0 {
		return math.MinInt64
	}
	if target == 0 {
		return 1
	}
	total := 0
	for _, num := range nums {
		count := combinationSum4R(cache, nums, target-num)
		if count > 0 {
			total += count
		}
	}
	cache[target] = total
	return total
}
