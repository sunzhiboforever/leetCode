package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{42, 1, 42, 35, 33, 37, 26, 3, 23, 29, 22, 50, 34, 31, 11, 28, 20, 31, 32, 28}

	S := 2
	fmt.Println(findTargetSumWays(nums, S))
}

func findTargetSumWays(nums []int, S int) int {
	cache := make(map[string]int)
	return findTargetSumWaysR(cache, nums, S, len(nums)-1)
}

func findTargetSumWaysR(cache map[string]int, nums []int, S int, i int) int {
	if S == 0 && i < 0 {
		return 1
	}
	if S != 0 && i < 0 {
		return 0
	}
	if _, ok := cache[strconv.FormatInt(int64(S), 10)+"|"+strconv.FormatInt(int64(i), 10)]; ok {
		return cache[strconv.FormatInt(int64(S), 10)+"|"+strconv.FormatInt(int64(i), 10)]
	}
	// 加法
	yes := findTargetSumWaysR(cache, nums, S+nums[i], i-1)
	// 减法
	no := findTargetSumWaysR(cache, nums, S-nums[i], i-1)
	cache[strconv.FormatInt(int64(S), 10)+"|"+strconv.FormatInt(int64(i), 10)] = yes + no
	return yes + no
}
