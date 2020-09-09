package _040_combination_sum_ii

import (
	`sort`
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var ans [][]int
	var dfs func(index int, cur int, combination []int)
	dfs = func(index int, cur int, combination []int) {
		if cur == target {
			tmp := make([]int, len(combination))
			copy(tmp, combination)
			ans = append(ans, tmp)
			return
		}
		for i := index; i < len(candidates); i++ {
			if cur+candidates[i] > target {
				break
			}
			cur += candidates[i]
			combination = append(combination, candidates[i])
			dfs(i+1, cur, combination)
			combination = combination[:len(combination)-1]
			cur -= candidates[i]
			for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
				i++
			}
		}
	}
	dfs(0, 0, nil)
	return ans
}
