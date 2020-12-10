package _441_arranging_coins

import "math"

func arrangeCoins(n int) int {
	var level int
	for n > level {
		level++
		n -= level
	}
	return level
}

func arrangeCoins1(n int) int {
	l, r := 1, (n+1)/2
	for l <= r {
		m := l + (r-l)/2
		s := m * (m + 1) / 2
		if s == n {
			return m
		} else if s > n {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return r
}

func arrangeCoins2(n int) int {
	return int(math.Sqrt(float64(n*2)+0.25) - 0.5)
}
