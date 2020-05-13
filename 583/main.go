package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(minDistance("sea", "eat"))
}

func minDistance(word1 string, word2 string) int {
	w1 := []rune(word1)
	w2 := []rune(word2)
	cache := make(map[string]int)
	min := minDistanceR(cache, w1, w2, len(w1)-1, len(w2)-1)
	return (len(w1) - min) + (len(w2) - min)
}

func minDistanceR(cache map[string]int, w1, w2 []rune, n1, n2 int) int {
	if n1 < 0 || n2 < 0 {
		return 0
	}
	cacheKey := strconv.FormatInt(int64(n1), 10) + "|" + strconv.FormatInt(int64(n2), 10)
	if _, ok := cache[cacheKey]; ok {
		return cache[cacheKey]
	}
	if w1[n1] == w2[n2] {
		return 1 + minDistanceR(cache, w1, w2, n1-1, n2-1)
	}
	c1 := minDistanceR(cache, w1, w2, n1-1, n2)
	c2 := minDistanceR(cache, w1, w2, n1, n2-1)
	if c1 > c2 {
		cache[cacheKey] = c1
	} else {
		cache[cacheKey] = c2
	}
	return cache[cacheKey]
}
