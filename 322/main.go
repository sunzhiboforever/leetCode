package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.MinInt64)
	coins := []int{2}
	amount := 3
	r := coinChange(coins, amount)
	fmt.Println(r)
}

func coinChange(coins []int, amount int) int {
	cache := make(map[int]int)
	count := coinChangeR(cache, coins, amount)
	if count > 99999 {
		count = -1
	}
	return count
}

func coinChangeR(cache map[int]int, coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	if _, ok := cache[amount]; ok {
		return cache[amount]
	}
	var min, count int
	min = math.MaxInt64
	for _, coin := range coins {
		count = coinChangeR(cache, coins, amount-coin)
		if count >= 0 && count < min {
			min = count + 1
		}
	}

	cache[amount] = min
	return min
}
