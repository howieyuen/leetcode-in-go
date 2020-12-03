package _948_bag_of_tokens

import (
	"sort"
)

func bagOfTokensScore(tokens []int, P int) int {
	sort.Ints(tokens)
	i, n := 0, len(tokens)-1
	var score int
	for i <= n {
		if P < tokens[i] { // 能量不够，尝试分换能量
			if score > 0 {
				// 置换能量后，消除当前token，score不增不减
				P += tokens[n] - tokens[i]
				n--
				i++
			} else {
				// score 不够置换，直接返回
				return score
			}
		} else {
			// 能量够，换分
			P -= tokens[i]
			score++
			i++
		}
	}
	return score
}
