package _213_house_robber_ii

// 198题目的扩展，相当于从nums[0:n-1]和nums[1:n]两个队列中找最大值
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	pre1 := 0
	cur1 := 0
	for i := 0; i < len(nums)-1; i++ {
		pre1, cur1 = cur1, max(cur1, pre1+nums[i])
	}
	pre2 := 0
	cur2 := 0
	for i := 1; i < len(nums); i++ {
		pre2, cur2 = cur2, max(cur2, pre2+nums[i])
	}
	return max(cur1, cur2)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
