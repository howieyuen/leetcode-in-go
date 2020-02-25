package _027_remove_element

import (
	`fmt`
	`testing`
)

func Test_removeElement(t *testing.T) {
	var nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	i := removeElement(nums, 2)
	fmt.Println(nums[:i])
	nums = []int{3, 2, 2, 3}
	i = removeElement(nums, 3)
	fmt.Println(nums[:i])
	nums = []int{1, 2, 4, 3}
	i = removeElement(nums, 5)
	fmt.Println(nums[:i])
	nums = []int{1}
	i = removeElement(nums, 1)
	fmt.Println(nums[:i])
}
