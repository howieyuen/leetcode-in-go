package _546_remove_boxes

// 用dp[l][r][k]表示移除区间[l, r]的元素，加上该区间右边等于boxes[r]的k个元素组成的序列的最大积分
func removeBoxes(boxes []int) int {
	dp := [100][100][100]int{}
	var dfs func(l, r, k int) int
	dfs = func(l, r, k int) int {
		if l > r {
			return 0
		}
		if dp[l][r][k] != 0 {
			return dp[l][r][k]
		}
		// 尽可能缩小右边界
		for l < r && boxes[r] == boxes[r-1] {
			r--
			k++
		}
		dp[l][r][k] = dfs(l, r-1, 0) + (k+1)*(k+1)
		// 在boxes[l,r]区间中，还有boxes[i]==boxes[r]，尝试其他方式，把等于boxes[r]的元素凑在一起
		for i := l; i < r; i++ {
			if boxes[i] == boxes[r] {
				dp[l][r][k] = max(dp[l][r][k], dfs(l, i, k+1)+dfs(i+1, r-1, 0))
			}
		}
		return dp[l][r][k]
	}
	return dfs(0, len(boxes)-1, 0)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
