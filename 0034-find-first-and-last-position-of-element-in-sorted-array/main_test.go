package _034_find_first_and_last_position_of_element_in_sorted_array

import (
	`fmt`
	"testing"
)

func Test_searchRange(t *testing.T) {
	var nums = []int{5, 7, 7, 8, 8, 10}
	var target = 8
	fmt.Println(searchRange(nums, target))
	nums = []int{5, 7, 7, 8, 8, 10}
	target = 6
	fmt.Println(searchRange(nums, target))
	nums = []int{6, 6, 6, 6, 6}
	target = 6
	fmt.Println(searchRange(nums, target))
}
