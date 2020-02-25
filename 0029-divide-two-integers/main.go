package _029_divide_two_integers

import (
	`math`
)

func divide(dividend int, divisor int) int {
	var flag = 1
	if dividend > 0 && divisor < 0 || dividend < 0 && divisor > 0 {
		flag = -1
	}
	var a = abs(dividend)
	var b = abs(divisor)
	var count = 0
	for a >= b {
		a -= b
		count++
	}
	count *= flag
	if count >= math.MaxInt32 {
		count = math.MaxInt32
	}
	if count < math.MinInt32 {
		count = math.MinInt32
	}
	return count
}

func abs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}