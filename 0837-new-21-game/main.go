package _837_new_21_game

/*
21点游戏，底分是0，每张牌的的大小从1到w
总分s大于=k时，停止抽牌，当s<=n时，即为胜利
求获胜的概率

获胜的概率，只和最后一次抽牌有关系，dp[x]表示，得分x开始游戏，并且获胜的概率，则dp[0]即为所求
x在[k, min(n, k+w-1)]之间时, dp[x] = 1，因为下一轮不论怎么抽，都不会超过n，给出递推方程：
	dp[x] = (dp[x+1]+dp[x+2]+....+dp[x+w])/w
	dp[x+1] = (dp[x+2]+dp[x+3]+....+dp[x+w+1])/w

计算关联关系：
	=> dp[x] - dp[x+1] = (dp[x+1]-dp[x+w+1])/w
	=> dp[x] = dp[x+1] - (dp[x+w-1]-dp[x+1])/w

但是上面的方程，不适合x=k-1的情况，x=k-1是临界值，需要手动计算：
	dp[k-1] = (dp[k]+dp[k+1]+...+dp[k+w]-1)/w

当x在[k, min(n, k+w-1)]之间时, dp[x] = 1，x > min(n, k+w-1)，dp[x] = 0，所以
	dp[k-1] = min(n-k+1, w)/w

*/
func new21Game(n int, k int, w int) float64 {
	if k == 0 || n >= k+w-1 {
		return 1.0
	}
	var dp = make([]float64, k+w)
	for i := k; i <= n && i < k+w; i++ {
		dp[i] = 1.0
	}
	dp[k-1] = float64(min(n-k+1, w)) / float64(w)
	for i := k - 2; i >= 0; i-- {
		dp[i] = dp[i+1] - (dp[i+w+1]-dp[i+1])/float64(w)
	}
	return dp[0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
