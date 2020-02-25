package _026_remove_duplicates_from_sorted_array

import (
	`fmt`
	`testing`
)

func Test_removeDuplicates(t *testing.T) {
	var nums = []int{0, 0}
	i := removeDuplicates(nums)
	fmt.Println(nums[:i])
}
