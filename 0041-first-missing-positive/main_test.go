package _041_first_missing_positive

import (
	`fmt`
	`testing`
)

func Test_firstMissingPositive(t *testing.T) {
	var nums = []int{1, 2, 3}
	fmt.Println(firstMissingPositive(nums))

	nums = []int{3, 4, -1, 1}
	fmt.Println(firstMissingPositive(nums))

	nums = []int{7, 8, 9, 11, 12}
	fmt.Println(firstMissingPositive(nums))

}
