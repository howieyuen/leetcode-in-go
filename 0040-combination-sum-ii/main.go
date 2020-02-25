package _040_combination_sum_ii

import (
	`sort`
)

var res [][]int

func combinationSum2(candidates []int, target int) [][]int {
	res = [][]int{}
	sort.Ints(candidates)
	dfs(candidates, []int{}, 0, 0, target)
	return res
}

func dfs(candidates, ans []int, index, cur, target int) {
	if cur == target {
		var tmp []int
		tmp = append(tmp, ans...)
		res = append(res, tmp)
		return
	}
	if index >= len(candidates) || cur > target {
		return
	}
	for i := index; i < len(candidates); i++ {
		if cur+candidates[i] > target {
			break
		}
		ans = append(ans, candidates[i])
		dfs(candidates, ans, i+1, cur+candidates[i], target)
		ans = ans[:len(ans)-1]
		for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
			i++
		}
	}
}
