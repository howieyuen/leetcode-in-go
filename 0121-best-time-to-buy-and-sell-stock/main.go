package _121_best_time_to_buy_and_sell_stock

import (
	"math"
)

// 只允许交易一次
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	// 继续不持有or今天卖出：dp[i][0] = max(dp[i-1][0], dp[i-1][1]+price[i])
	// 继续持有，or今天买入：dp[i][1] = max(dp[i-1][1], -price[i])
	// 遍历过程与i无关
	var todayNoHold = 0           // 第i天，不持有
	var todayHold = math.MinInt64 // 第i天，持有
	for i := range prices {
		todayNoHold = max(todayNoHold, todayHold+prices[i])
		todayHold = max(todayHold, -prices[i])
	}
	return todayNoHold
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProfit1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	max := 0
	buy := prices[0]
	for i := 1; i < len(prices); i++ {
		if buy > prices[i] {
			buy = prices[i]
		} else if prices[i]-buy > max {
			max = prices[i] - buy
		}
	}
	return max
}
