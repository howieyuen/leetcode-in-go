package _122_best_time_to_buy_and_sell_stock_ii

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	// dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
	// dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
	var todayNoHold = 0
	var todayHold = -prices[0]
	for i := range prices {
		tmp := todayNoHold
		todayNoHold = max(todayNoHold, todayHold+prices[i])
		todayHold = max(todayHold, tmp-prices[i])
	}
	return todayNoHold
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProfit2(prices []int) int {
	var max = 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			max += prices[i] - prices[i-1]
		}
	}
	return max
}

func maxProfit1(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}
	i := 0
	valley := prices[0]
	peak := prices[0]
	maxProfit := 0
	for i < len(prices)-1 {
		for i < len(prices)-1 && prices[i] > prices[i+1] {
			i++
		}
		valley = prices[i]
		for i < len(prices)-1 && prices[i] <= prices[i+1] {
			i++
		}
		peak = prices[i]
		maxProfit += peak - valley
	}
	return maxProfit
}
