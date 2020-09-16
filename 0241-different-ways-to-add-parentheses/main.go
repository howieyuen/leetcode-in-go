package _241_different_ways_to_add_parentheses

import (
	"strconv"
)

func diffWaysToCompute(input string) []int {
	var ans []int
	for i, c := range input {
		if c == '+' || c == '-' || c == '*' {
			lefts := diffWaysToCompute(input[:i])
			rights := diffWaysToCompute(input[i+1:])
			for _, left := range lefts {
				for _, right := range rights {
					if c == '-' {
						ans = append(ans, left-right)
					} else if c == '+' {
						ans = append(ans, left+right)
					} else if c == '*' {
						ans = append(ans, left*right)
					}
				}
			}
		}
	}
	if len(ans) == 0 {
		num, _ := strconv.Atoi(input)
		ans = append(ans, num)
	}
	return ans
}
