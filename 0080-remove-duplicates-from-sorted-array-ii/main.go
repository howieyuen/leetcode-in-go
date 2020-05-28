package _080_remove_duplicates_from_sorted_array_ii

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	index := 0
	for i := range nums {
		if i < 2 || nums[i] != nums[index-2] {
			nums[index] = nums[i]
			index++
		}
	}
	fmt.Println(nums)
	return index
}
