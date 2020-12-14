package _368_largest_divisible_subset

import "sort"

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	var lookup = make(map[int][]int)
	var ans []int
	for _, v := range nums {
		tmp := []int{v}
		for k := range lookup {
			if v%k == 0 && len(lookup[k])+1 > len(tmp) {
				tmp = make([]int, len(lookup[k]), len(lookup[k])+1)
				copy(tmp, lookup[k])
				tmp = append(tmp, v)
			}
		}
		lookup[v] = tmp
		if len(lookup[v]) > len(ans) {
			ans = lookup[v]
		}
	}
	return ans
}

func largestDivisibleSubset1(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	var dp = make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	var preIndex = make([]int, n)
	for i := range preIndex {
		preIndex[i] = - 1
	}
	var length, lastIndex = 0, -1
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
				preIndex[i] = j
			}
		}
		if dp[i] > length {
			length = dp[i]
			lastIndex = i
		}
	}
	var res = make([]int, length)
	for i := lastIndex; i != -1; i = preIndex[i] {
		res[length-1] = nums[i]
		length--
	}
	return res
}
