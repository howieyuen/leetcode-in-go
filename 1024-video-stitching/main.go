package _024_video_stitching

// 动态规划
func videoStitching(clips [][]int, T int) int {
	var dp = make([]int, T+1)
	for i := range dp {
		dp[i] = T + 1
	}
	dp[0] = 0
	for i := 0; i <= T; i++ {
		for _, clip := range clips {
			if clip[0] <= i && clip[1] >= i {
				dp[i] = min(dp[i], dp[clip[0]]+1)
			}
		}
	}
	if dp[T] == T+1 {
		return -1
	}
	return dp[T]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 贪心算法
func videoStitching1(clips [][]int, T int) int {
	var maxRight = make([]int, T)
	for _, clip := range clips {
		left, right := clip[0], clip[1]
		if left < T && right > maxRight[left] {
			maxRight[left] = right
		}
	}
	var rightest, lastRightest int
	var ans int
	for i, right := range maxRight {
		if right > rightest {
			rightest = right
		}
		if i == rightest {
			return -1
		}
		if i == lastRightest {
			ans++
			lastRightest = rightest
		}
	}
	return ans
}
