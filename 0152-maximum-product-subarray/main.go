package _152_maximum_product_subarray

/*
nums数组存在正负，则当前最优解不是单纯地依赖上一个最优解得到
当nums[i]<0,依赖的是上个子问题的最小值，而不是最大值
则维护2个变量max，min
*/
func maxProduct(nums []int) int {
	maxP, minP, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxP, minP
		maxP = max(nums[i], max(mx*nums[i], mn*nums[i]))
		minP = min(nums[i], min(mn*nums[i], mx*nums[i]))
		ans = max(ans, max(maxP, minP))
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
