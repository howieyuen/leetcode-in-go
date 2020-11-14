package _122_relative_sort_array

import (
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	var dict = make(map[int]int)
	for i, a := range arr2 {
		dict[a] = i - len(arr2)
	}
	sort.Slice(arr1, func(i, j int) bool {
		x, y := arr1[i], arr1[j]
		if r, has := dict[x]; has {
			x = r
		}
		if r, has := dict[y]; has {
			y = r
		}
		return x<y
	})
	return arr1
}
