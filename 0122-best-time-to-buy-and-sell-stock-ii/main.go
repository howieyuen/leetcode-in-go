package _122_best_time_to_buy_and_sell_stock_ii

func maxProfit(prices []int) int {
	var max = 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			max += prices[i] - prices[i-1]
		}
	}
	return max
}

func maxProfit1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var max = 0
	low, high := prices[0], prices[0]
	for i := 1; i < len(prices)-1; {
		for i < len(prices)-1 && prices[i] > prices[i+1] {
			i++
		}
		low = prices[i]
		for i < len(prices)-1 && prices[i] <= prices[i+1] {
			i++
		}
		high = prices[i]
		max += high - low
	}
	return max
}
