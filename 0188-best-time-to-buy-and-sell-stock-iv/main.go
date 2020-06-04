package _188_best_time_to_buy_and_sell_stock_iv

func maxProfit(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	if k >= len(prices)/2 {
		return maxProfitWithInfiniteK(prices)
	}
	return maxProfitWithFiniteK(prices, k)
}

func maxProfitWithFiniteK(prices []int, k int) int {
	var dp = make([][][]int, len(prices))
	for i := range dp {
		dp[i] = make([][]int, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}
	for i := range prices {
		for j := 1; j <= k; j++ {
			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
			} else {
				dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
				dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
			}
		}
	}
	return dp[len(prices)-1][k][0]
}

func maxProfitWithInfiniteK(prices []int) int {
	var max = 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			max += prices[i] - prices[i-1]
		}
	}
	return max
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
