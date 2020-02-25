package _739_daily_temperatures

import (
	`math`
)

func dailyTemperatures(T []int) []int {
	ans := make([]int, len(T))
	var stack []int
	index := -1
	for i := len(T) - 1; i >= 0; i-- {
		for len(stack) > 0 && T[i] >= T[stack[index]] {
			stack = stack[:index]
			index--
		}
		if index > -1 {
			ans[i] = stack[index] - i
		}
		stack = append(stack, i)
		index++
	}
	return ans
}

func dailyTemperatures1(T []int) []int {
	ans := make([]int, len(T))
	next := make([]int, 101)
	for i := 0; i <= 100; i++ {
		next[i] = math.MaxInt64
	}
	for i := len(T) - 1; i >= 0; i-- {
		warmerIndex := math.MaxInt64
		for t := T[i] + 1; t <= 100; t++ {
			if next[t] < warmerIndex {
				warmerIndex = next[t]
			}
		}
		if warmerIndex < math.MaxInt64 {
			ans[i] = warmerIndex - i
		}
		next[T[i]] = i
	}
	return ans
}
