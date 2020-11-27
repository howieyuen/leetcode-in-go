package _842_split_array_into_fibonacci_sequence

import (
	"math"
	"strconv"
)

func splitIntoFibonacci(S string) []int {
	var ans []int
	var dfs func(index int, pre1, pre2, deep int) bool
	dfs = func(index int, pre1, pre2, deep int) bool {
		if index == len(S) {
			return deep >= 3
		}
		for i := 1; i <= 11; i++ {
			if index+i > len(S) || S[index] == '0' && i > 1 {
				break
			}
			num, _ := strconv.Atoi(S[index : index+i])
			if num > math.MaxInt32 || (deep != 0 && deep != 1 && num > pre1+pre2) {
				break
			}
			if deep == 0 || deep == 1 || pre1+pre2 == num {
				ans = append(ans, num)
				if dfs(index+i, pre2, num, deep+1) {
					return true
				}
				ans = ans[:len(ans)-1]
			}
		}
		return false
	}
	dfs(0, 0, 0, 0)
	return ans
}
