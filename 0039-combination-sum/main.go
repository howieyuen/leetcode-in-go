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

func combinationSum1(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var ans [][]int
	var dfs func(index int, combination []int, cur int)
	dfs = func(index int, combination []int, cur int) {
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
			dfs(i, combination, cur)
			combination = combination[:len(combination)-1]
			cur -= candidates[i]
		}
	}
	dfs(0, nil, 0)
	return ans
}
