package _518_coin_change_2

func change(amount int, coins []int) int {
	var dp = make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for x := coin; x <= amount; x++ {
			dp[x] += dp[x-coin]
		}
	}
	return dp[amount]
}

// 超时
func change1(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}
	var ans int
	var dfs func(left int, index int)
	dfs = func(left int, index int) {
		if index == len(coins) {
			return
		}
		if left == 0 {
			ans++
			return
		}
		for i := index; i < len(coins); i++ {
			if coins[i] > left {
				continue
			}
			dfs(left-coins[i], i)
		}
	}
	dfs(amount, 0)
	return ans
}
