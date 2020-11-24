package _401_binary_watch

import (
	"fmt"
	"strings"
)

func readBinaryWatch(num int) []string {
	var bases = []int{8, 4, 2, 1, 32, 16, 8, 4, 2, 1}
	var backtrace func(left int, h, m int, index int)
	var ans []string
	backtrace = func(left int, h, m int, index int) {
		if left == 0 {
			ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
			return
		}
		for i := index; i < len(bases); i++ {
			if i < 4 {
				h += bases[i]
			} else {
				m += bases[i]
			}
			if h <= 11 && m <= 59 {
				backtrace(left-1, h, m, i+1)
			}
			if i < 4 {
				h -= bases[i]
			} else {
				m -= bases[i]
			}
		}
	}
	backtrace(num, 0, 0, 0)
	return ans
}

func readBinaryWatch1(num int) []string {
	var result []string
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			b1 := fmt.Sprintf("%b", i)
			b2 := fmt.Sprintf("%b", j)
			sumOne := strings.Count(b1, "1") + strings.Count(b2, "1")
			if sumOne == num {
				result = append(result, fmt.Sprintf("%d:%02d", i, j))
			}
		}
	}
	return result
}
