package _354_russian_doll_envelopes

import (
	"sort"
)
// see more 300.longest-increasing-subsequence
func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}
	// 长度按照升序排列，宽度按照降序排列
	// 排序后等于把在二维(长、宽)上的最长递增子序列问题转换成一维(宽)上的最长递增子序列的查找
	// 宽度按降序排列是为了避免长相等，宽度小的也被列入递增序列中
	// 例如：[3,4]和[3,5]是不能嵌套的
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	var dp = make([]int, len(envelopes))
	for i := range dp {
		dp[i] = 1
	}
	var ans = 0
	for _, env := range envelopes {
		left, right := 0, ans
		for left < right {
			mid := (left + right) >> 1
			if env[1] > dp[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		dp[left] = env[1]
		if left == ans {
			ans++
		}
	}
	return ans
}

func maxEnvelopes1(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] < envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	var dp = make([]int, len(envelopes))
	for i := range dp {
		dp[i] = 1
	}
	var ans = 1
	for i := 1; i < len(envelopes); i++ {
		for j := 0; j < i; j++ {
			if envelopes[i][0] > envelopes[j][0] && envelopes[i][1] > envelopes[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
