package main

import "fmt"

func main() {
	a := findMedianSortedArrays([]int{1, 3}, []int{2})
	fmt.Println(a)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 判断奇偶
	if (len(nums1)+len(nums2))%2 == 0 {
		return float64(findTopK(nums1, nums2, (len(nums1)+len(nums2))/2)+findTopK(nums1, nums2, (len(nums1)+len(nums2))/2-1)) / 2
	} else {
		return float64(findTopK(nums1, nums2, (len(nums1)+len(nums2))/2))
	}

}

func findTopK(nums1 []int, nums2 []int, k int) int {
	// 取其中一个的中位数
	var mid int
	if len(nums1) > 0 {
		mid = nums1[len(nums1)/2]
	} else {
		mid = nums2[len(nums2)/2]
	}

}
