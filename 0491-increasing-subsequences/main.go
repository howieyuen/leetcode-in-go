package _491_increasing_subsequences

import (
	"math"
)

func findSubsequences(nums []int) [][]int {
	var ans [][]int
	var dfs func(index int, last int, cur []int)
	dfs = func(index int, last int, cur []int) {
		if index == len(nums) {
			if len(cur) >= 2 {
				var tmp = make([]int, len(cur))
				copy(tmp, cur)
				ans = append(ans, tmp)
			}
			return
		}
		if nums[index] >= last {
			cur = append(cur, nums[index])
			dfs(index+1, nums[index], cur)
			cur = cur[:len(cur)-1]
		}
		if nums[index] != last {
			dfs(index+1,last,  cur)
		}
	}
	dfs(0, math.MinInt32, nil)
	return ans
}
