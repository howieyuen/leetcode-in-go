package _498_number_of_subsequences_that_satisfy_the_given_sum_condition

import "sort"

func numSubseq(nums []int, target int) int {
	const base int = 1e9 + 7
	sort.Ints(nums)

	var pow = make([]int, len(nums))
	pow[0] = 1
	for n := 1; n < len(nums); n++ {
		pow[n] = (pow[n-1] * 2) % (1e9 + 7)
	}

	var ans int
	l, r := 0, len(nums)-1
	for l <= r {
		if nums[l]+nums[r] > target {
			r--
		} else {
			ans = (ans + pow[r-l]) % base
			l++
		}
	}
	return ans
}
