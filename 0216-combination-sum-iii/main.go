package _216_combination_sum_iii

func combinationSum3(k int, n int) [][]int {
	var ans [][]int
	var dfs func(cur, sum int, combine []int)
	dfs = func(cur, sum int, combine []int) {
		if sum == n && len(combine) == k {
			var tmp = make([]int, len(combine))
			copy(tmp, combine)
			ans = append(ans, tmp)
		} else if sum >= n {
			return
		} else if len(combine) >= k {
			return
		}
		for next := cur; next <= 9; next++ {
			combine = append(combine, next)
			sum += next
			dfs(next+1, sum, combine)
			sum -= next
			combine = combine[:len(combine)-1]
		}
	}
	dfs(1, 0, nil)
	return ans
}
