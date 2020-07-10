package _714_best_time_to_buy_and_sell_stock_with_transaction_fee

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	// dp[i][0] = max(dp[i-1][0], dp[i-1][1]+price[i]-fee)
	// dp[i][1] = max(dp[i-1][1], dp[i-1][0]-price[i])
	var todayNoHold = 0
	var todayHold = -prices[0]
	for i := range prices {
		tmp := todayNoHold
		todayNoHold = max(todayNoHold, todayHold+prices[i]-fee)
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
