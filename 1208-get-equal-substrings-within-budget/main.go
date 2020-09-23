package _208_get_equal_substrings_within_budget

func equalSubstring(s string, t string, maxCost int) int {
	var ans int
	var cost int
	left, right := 0, 0
	for right < len(s) {
		cost += abs(int(s[right]) - int(t[right]))
		for cost > maxCost {
			cost -= abs(int(s[left]) - int(t[left]))
			left++
		}
		ans = max(ans, right-left+1)
		right++
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}
