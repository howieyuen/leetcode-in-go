package main

import (
	`math`
)

func reverse(x int) int {
	var out int
	for x != 0 {
		out = out*10 + x%10
		if out > math.MaxInt32 || out < math.MinInt32 {
			return 0
		}
		x /= 10
	}
	return out
}