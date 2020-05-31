package _084_largest_rectangle_in_histogram

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	// monoStack存放的是下标，下标对用的值是单调递增的
	var monoStack []int
	for i := 0; i < n; i++ {
		// 当前元素值比栈顶元素小
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			// 更新当前栈顶元素的右边界
			right[monoStack[len(monoStack)-1]] = i
			// 出栈，保持单调性
			monoStack = monoStack[:len(monoStack)-1]
		}
		// 单调栈为空，添加哨兵，避免栈空判断，同时也为了计算宽度时，避免单个元素宽度j-i-1为负数
		if len(monoStack) == 0 {
			left[i] = -1
		} else {
			// 更新栈顶元素对应的左边界
			left[i] = monoStack[len(monoStack)-1]
		}
		// 当前元素下标入栈
		monoStack = append(monoStack, i)
	}
	ans := 0
	// 根据左右边界确定宽度，找到最大值
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
