package _090_subsets_ii

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	var backtrace func(index int, size int, cur []int)
	backtrace = func(index int, size int, cur []int) {
		if len(cur) == size {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			ans = append(ans, tmp)
			return
		}
		for i := index; i < len(nums); i++ {
			if i > index && nums[i] == nums[i-1] {
				continue
			}
			cur = append(cur, nums[i])
			backtrace(i+1, size, cur)
			cur = cur[:len(cur)-1]
		}
	}
	for i := 0; i <= len(nums); i++ {
		backtrace(0, i, []int{})
	}
	return ans
}
