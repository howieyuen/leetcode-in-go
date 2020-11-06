package _356_sort_integers_by_the_number_of_1_bits

import (
	"sort"
)

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		a, b := zeroBit(arr[i]), zeroBit(arr[j])
		if a > b {
			return false
		} else if a < b {
			return true
		} else {
			return arr[i] < arr[j]
		}
	})
	return arr
}

func zeroBit(a int) int {
	var ans int
	for a != 0 {
		a &= a - 1
		ans++
	}
	return ans
}
