package _018_4Sum

import "sort"

func fourSum(nums []int, target int) [][]int {
	var ans [][]int
	if nums == nil || len(nums) < 4 {
		return ans
	}
	sort.Ints(nums)
	for a := 0; a < len(nums)-3; a++ {
		if a > 0 && nums[a-1] == nums[a] {
			continue
		}
		for b := a + 1; b < len(nums)-2; b++ {
			if b > a+1 && nums[b-1] == nums[b] {
				continue
			}
			c := b + 1
			d := len(nums) - 1
			for c < d {
				if c > b+1 && nums[c-1] == nums[c] {
					c++
					continue
				}
				if d < len(nums)-1 && nums[d+1] == nums[d] {
					d--
					continue
				}
				sum := nums[a] + nums[b] + nums[c] + nums[d]
				if sum == target {
					zero := []int{nums[a], nums[b], nums[c], nums[d]}
					ans = append(ans, zero)
					c++
					d--
				} else if sum < target {
					c++
				} else {
					d--
				}
			}
		}
	}
	return ans
}
