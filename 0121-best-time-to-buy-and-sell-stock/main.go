package _121_best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
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