package _983_minimum_cost_for_tickets

/*
	days数组记录要出行的日期，取值1~365且为升序
	cost数组表示买1天，7天，30天票的假期，固定长度为3
	求旅行计划最少开销

	动态规划问题，核心在于转移方程的分析，dp[i]表示，到第i天一共需要的最小开销
第i天不出行，当天开销就是0，即
	dp[i] = dp[i-1]
第i天要出行，当天开销可以是现买1天的票，可以是7天前买的7连票，可以是30天买的30连票，即
	dp[i] = min(dp[i-1]+cost[0], dp[i-7]+cost[1], dp[i-30]+cost[2])
*/
func mincostTickets(days []int, costs []int) int {
	n := days[len(days)-1]
	dp := make([]int, n+1)
	for i := range days { // 记录出行日期
		dp[days[i]] = 1
	}
	for i := 1; i <= n; i++ {
		if dp[i] == 0 { // 不出行
			dp[i] = dp[i-1]
		} else {
			a := dp[i-1] + costs[0]
			b := costs[1]
			if i >= 7 {
				b = dp[i-7] + costs[1]
			}
			c := costs[2]
			if i >= 30 {
				c = dp[i-30] + costs[2]
			}
			dp[i] = min(a, min(b, c))
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
