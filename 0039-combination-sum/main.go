package _039_combination_sum

import (
	`sort`
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var ans [][]int
	backtrace(candidates, []int{}, 0, 0, target, &ans)
	return ans
}

func backtrace(candidates, combinations []int, index, current, target int, ans *[][]int) {
	if current == target {
		var tmp []int
		tmp = append(tmp, combinations...)
		*ans = append(*ans, tmp)
		return
	} else {
		for i := index; i < len(candidates); i++ {
			if current+candidates[i] > target {
				break
			}
			combinations = append(combinations, candidates[i])
			current += candidates[i]
			backtrace(candidates, combinations, i, current, target, ans)
			current -= candidates[i]
			combinations = combinations[:len(combinations)-1]
		}
	}
}
