package _072_edit_distance

/*
1、状态定义：
dp[i][j]表示word1的前i个字母转换成word2的前j个字母所使用的最小编辑距离

2、状态转移：
i指向word1，j指向word2
若当前字母相同，则dp[i][j] = dp[i - 1][j - 1];
否则取增删替三个操作的最小值 + 1， 即:
dp[i][j] = min(dp[i][j - 1], dp[i - 1][j], dp[i - 1][j - 1]) + 1
*/
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	var dp = make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 插入、删除、替换的操作数均为1
				dp[i][j] = 1 + min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1]))
			}
		}
	}
	return dp[m][n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
