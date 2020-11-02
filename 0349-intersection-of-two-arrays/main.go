package _349_intersection_of_two_arrays

import "sort"

func intersection(nums1 []int, nums2 []int) []int {
	var ans []int
	var maps1 = make(map[int]bool)
	for _, num := range nums1 {
		if _, ok := maps1[num]; !ok {
			maps1[num] = true
		}
	}
	var maps2 = make(map[int]bool)
	for _, num := range nums2 {
		if _, ok := maps2[num]; !ok {
			maps2[num] = true
		}
	}
	for key := range maps1 {
		if maps2[key] {
			ans = append(ans, key)
		}
	}
	return ans
}

func intersection1(nums1 []int, nums2 []int) (res []int) {
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		x, y := nums1[i], nums2[j]
		if x == y {
			if res == nil || x > res[len(res)-1] {
				res = append(res, x)
			}
			i++
			j++
		} else if x < y {
			i++
		} else {
			j++
		}
	}
	return
}
