package leetcode

import (
	"slices"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	slices.Sort(nums1)
	var median float64
	if len(nums1)%2 == 0 {
		median = (float64(nums1[len(nums1)/2]) + float64(nums1[(len(nums1)/2)-1])) / 2
	} else {
		median = float64(nums1[len(nums1)/2])
	}
	return median
}
