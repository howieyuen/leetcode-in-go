package _123_best_time_to_buy_and_sell_stock_iii

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	firstBuy := -prices[0]
	firstSell := 0
	secondBuy := -prices[0]
	secondSell := 0
	for i := 1; i < len(prices); i++ {
		secondSell = max(secondSell, secondBuy+prices[i])
		secondBuy = max(secondBuy, firstSell-prices[i])
		firstSell = max(firstSell, firstBuy+prices[i])
		firstBuy = max(firstBuy, -prices[i])
	}
	return secondSell
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
	dp[i][k][0] 第i天，已交易k次，当前不持有
	dp[i][k][1] 第i天，已交易k次，当前持有
*/
func maxProfit1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var dp = make([][][]int, len(prices))
	for i := range dp {
		dp[i] = make([][]int, 3)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}
	for i := 0; i < len(prices); i++ {
		for k := 1; k <= 2; k++ {
			if i == 0 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
			} else {
				dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
				dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
			}
		}
	}
	return dp[len(prices)-1][2][0]
}
