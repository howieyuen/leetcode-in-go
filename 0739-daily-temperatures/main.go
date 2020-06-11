package _739_daily_temperatures

import (
	`math`
)

/*
维护单调递减栈
当前温度比栈顶元素对应的温度高，则栈顶元素找到了第一个温度更高的天数，相减即为等待天数
*/
func dailyTemperatures(t []int) []int {
	if len(t) == 0 {
		return []int{}
	}
	var ans = make([]int, len(t))
	var stack = make([]int, 0, len(t))
	for i := 0; i < len(t); i++ {
		for len(stack) > 0 && t[stack[len(stack)-1]] < t[i] {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[j] = i - j
		}
		stack = append(stack, i)
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
