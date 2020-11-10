package _031_next_permutation

import (
	`fmt`
	`testing`
)

func Test_nextPermutation(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	for i := 0; i < 6; i++ {
		nextPermutation1(nums)
		fmt.Println(nums)
	}
}
