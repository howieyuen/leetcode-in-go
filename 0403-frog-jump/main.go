package _403_frog_jump

func canCross(stones []int) bool {
	n := len(stones)
	if stones[1] != 1 {
		return false
	}
	// dp[i][k] 对于第i石头来说,能否通过k步到达。
	// 其中三种情况：j石头走k-1步；j石头走k步；j石头走k+1步
	// 对于j来说，可取值为[0, i)
	// 对于i来说，最小为1（两个数，0和1位置）
	var dp = make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			k := stones[i] - stones[j]
			if k <= j+1 {
				dp[i][k] = dp[j][k-1] || dp[j][k] || dp[j][k+1]
				if i == n-1 && dp[i][k] {
					return true
				}
			}
		}
	}
	return false
}

func canCross1(stones []int) bool {
	var dp = make(map[int]map[int]bool)
	for _, v := range stones {
		dp[v] = make(map[int]bool)
	}
	dp[0][0] = true
	for _, v := range stones {
		for k := range dp[v] {
			for step := k - 1; step <= k+1; step++ {
				if step > 0 {
					if _, ok := dp[v+step]; ok {
						dp[v+step][step] = true
					}
				}
			}
		}
	}
	return len(dp[stones[len(stones)-1]]) > 0
}
