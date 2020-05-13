package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		m[v] = k
	}
	fmt.Println(m)
	for k, num := range nums {
		if v, ok := m[target-num]; ok && k != v {
			return []int{k, v}
		}
	}
	return []int{}
}
