package _384_shuffle_an_array

import (
	"math/rand"
)

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}

/** Resets the array to its original configuration and return it. */
func (s *Solution) Reset() []int {
	return s.nums
}

/** Returns a random shuffling of the array. */
func (s *Solution) Shuffle() []int {
	tmp := make([]int, len(s.nums))
	copy(tmp, s.nums)
	for i := len(tmp) - 1; i >= 0; i-- {
		next := rand.Intn(i + 1)
		tmp[i], tmp[next] = tmp[next], tmp[i]
	}
	return tmp
}
