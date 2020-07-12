package test2

// 两个蚂蚁相撞之后会互相调头，那么是不是蚂蚁碰撞的调头 就等于 穿透？
// 那么题目就变成了求单只最晚落地的蚂蚁，与碰撞无关
func getLastMoment(n int, left []int, right []int) int {
	var ans = -1
	for i := range left {
		ans = max(ans, left[i])
	}
	for i := range right {
		ans = max(ans, n-right[i])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
