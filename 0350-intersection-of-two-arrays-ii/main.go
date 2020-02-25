package _350_intersection_of_two_arrays_ii

import (
	`sort`
)

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	i, j, k := 0, 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			nums1[k] = nums2[j]
			k++
			i++
			j++
		}
	}
	return nums1[:k]
}
