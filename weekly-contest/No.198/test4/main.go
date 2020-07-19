package test4

func closestToTarget(arr []int, target int) int {
	var ans = 1000000000
	var dp = make(map[int]bool)
	for _, v := range arr {
		var tmp = make(map[int]bool)
		tmp[v] = true
		ans = min(ans, abs(v-target))
		for k := range dp {
			ans = min(ans, abs(v&k-target))
			tmp[v&k] = true
		}
		dp = tmp
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
