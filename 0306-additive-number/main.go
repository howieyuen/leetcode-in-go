package _306_additive_number

import "strconv"

func isAdditiveNumber(num string) bool {
	var dfs func(index int, sum int, pre int, k int) bool
	dfs = func(index int, sum int, pre int, k int) bool {
		if index == len(num) {
			return k > 2
		}
		for i := index; i < len(num); i++ {
			if index < i && num[index] == '0' {
				continue
			}
			cur, _ := strconv.Atoi(num[index : i+1])
			if k >= 2 && cur != sum {
				continue
			}
			if dfs(i+1, pre+cur, cur, k+1) {
				return true
			}
		}
		return false
	}
	return dfs(0, 0, 0, 0)
}
