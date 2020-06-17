package _704_delete_and_earn

// 类比打家劫舍，打劫一家，相邻的就会报警
// 获取点数，就得删除+1和-1的点数
func deleteAndEarn(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var m = nums[0]
	for i := range nums {
		m = max(m, nums[i])
	}
	var times = make([]int, m+1)
	for i := range nums {
		times[nums[i]] += nums[i]
	}
	var pre = times[0]
	var cur = times[1]
	for i := 2; i < len(times); i++ {
		pre, cur = cur, max(cur, pre+times[i])
	}
	return cur
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
