package _746_min_cost_climbing_stairs

func minCostClimbingStairs(cost []int) int {
	f0 := 0
	f1 := 0
	f2 := 0
	for i := len(cost) - 1; i >= 0; i-- {
		if f1 < f2 {
			f0 = cost[i] + f1
		} else {
			f0 = cost[i] + f2
		}
		f2 = f1
		f1 = f0
	}
	if f1 < f2 {
		return f1
	} else {
		return f2
	}
}

func minCostClimbingStairs1(cost []int) int {
	if len(cost) < 2 {
		return 0
	}
	for i := 2; i < len(cost); i++ {
		cost[i] = min(cost[i-1], cost[i-2]) + cost[i]

	}
	return min(cost[len(cost)-1], cost[len(cost)-2])
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
