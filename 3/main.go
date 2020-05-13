package main

import "fmt"

func main() {
	a := lengthOfLongestSubstring("aab")
	fmt.Println(a)
}

func lengthOfLongestSubstring(s string) int {
	r := []rune(s)
	fmt.Println(r)
	max := 0

	for k := 0; k < len(r); k++ {
		m := make(map[rune]int)
		for i := k; i < len(r); i++ {
			if _, ok := m[r[i]]; ok {
				break
			}
			m[r[i]] = 0
		}
		if max < len(m) {
			max = len(m)
		}
	}
	return max
}
