package _016_3Sum_Closest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closest := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			min := nums[i] + nums[left] + nums[left+1]
			if min > target {
				if math.Abs(float64(min-target)) < math.Abs(float64(closest-target)) {
					closest = min
					break
				}
			}
			
			max := nums[i] + nums[right-1] + nums[right]
			if max < target {
				if math.Abs(float64(max-target)) < math.Abs(float64(closest-target)) {
					closest = max
					break
				}
			}
			
			cur := nums[i] + nums[left] + nums[right]
			if math.Abs(float64(cur-target)) < math.Abs(float64(closest-target)) {
				closest = cur
			}
			if cur == target {
				return target
			} else if cur < target {
				left++
				for left < right && nums[left-1] == nums[left] {
					left++
				}
			} else {
				right--
				for left < right && nums[right+1] == nums[right] {
					right--
				}
			}
		}
	}
	return closest
}
