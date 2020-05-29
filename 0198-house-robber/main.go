package _198_house_robber

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	pre := nums[0]
	if len(nums) == 1 {
		return pre
	}
	cur := max(pre, nums[1])
	for i := 2; i < len(nums); i++ {
		pre, cur = cur, max(cur, pre+nums[i])
	}
	return cur
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}