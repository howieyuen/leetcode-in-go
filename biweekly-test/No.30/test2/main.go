package test2

import (
	"sort"
)

const base = 1e9 + 7

func rangeSum(nums []int, n int, left int, right int) int {
	var sum = make([]int, n*(n+1)/2)
	var index = 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if index == 0 {
				sum[index] = nums[0]
			} else {
				sum[index] += nums[j]
				if i != j {
					sum[index] += sum[index-1]
				}
			}
			index++
		}
	}
	sort.Ints(sum)
	var ans int
	for i := left - 1; i < right; i++ {
		ans = (ans + sum[i]) % base
	}
	return ans
}
