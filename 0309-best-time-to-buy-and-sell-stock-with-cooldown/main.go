package _309_best_time_to_buy_and_sell_stock_with_cooldown

import (
	"math"
)

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	// 第i天，不持有: dp[i][0] = max(dp[i-1][0], dp[i-1][1] + price[i])
	// 第i天，持有:   dp[i][1] = max(dp[i-1][1], dp[i-2][0] - price[i])
	var todayNoHold = 0           // dp[i][0]
	var todayHold = math.MinInt64 // dp[i][1]
	var yesterdayNoHold = 0       // dp[i-2][0]
	for i := 0; i < n; i++ {
		tmp := todayNoHold
		todayNoHold = max(todayNoHold, todayHold+prices[i])
		todayHold = max(todayHold, yesterdayNoHold-prices[i])
		yesterdayNoHold = tmp
	}
	return todayNoHold
}

func maxProfit1(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	// dp[i]: 表示第i天的最终状态
	// dp[i][0]: 手上持有股票的最大收益
	// dp[i][1]: 手上不持有股票，并且处于冷冻期中的累计最大收益
	// dp[i][2]: 手上不持有股票，并且不在冷冻期中的累计最大收益
	var dp = make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 3)
	}
	dp[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		// 第i天持有，第i-1天也持有 VS 第i-1天不持有且不在冷冻期，且在第i天买入
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		
		// 第i天不持有且在冷冻期，第i-1持有，且第i天卖出
		dp[i][1] = dp[i-1][0] + prices[i]
		
		// 第i天不持有且不在冷冻期，第i-1天不持有，且在冷冻期 VS 第i-1天不持有，也不在冷冻期
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	return max(dp[n-1][1], dp[n-1][2])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
