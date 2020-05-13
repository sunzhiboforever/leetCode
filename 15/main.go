package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{1, -1, -1, 0}))
}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {

			sum := nums[right] + nums[left]

			if sum == 0-nums[i] {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
			if sum > 0-nums[i] {
				right--
				continue
			}
			if sum < 0-nums[i] {
				left++
				continue
			}
		}
	}
	return res
}
