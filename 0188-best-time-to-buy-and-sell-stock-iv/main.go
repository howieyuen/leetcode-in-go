package _188_best_time_to_buy_and_sell_stock_iv

import "math"

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

func maxProfit1(k int, prices []int) int {
	if k <= 0 {
		return 0
	}
	if k >= len(prices)/2 {
		return greedy(prices)
	}
	// dp[i][0] 表示i次交易，当前不持有
	// dp[i][1] 表示i次交易，当前持有
	// 交易要先买入，再卖出，一对买入卖出是一次完整交易
	// 所以，dp[i][1]与dp[i-1][0]有关，dp[i][0]与dp[i][1]有关
	var dp = make([][]int, k)
	for i := range dp {
		dp[i] = make([]int, 2)
		dp[i][1] = math.MinInt32
	}
	for _, p := range prices {
		dp[0][1] = max(dp[0][1], -p)
		dp[0][0] = max(dp[0][0], dp[0][1]+p)
		for i := 1; i < k; i++ {
			dp[i][1] = max(dp[i][1], dp[i-1][0]-p)
			dp[i][0] = max(dp[i][0], dp[i][1]+p)
		}
	}
	return dp[k-1][0]
}

func greedy(prices []int) int {
	var res int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}
