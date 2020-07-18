package _042_trapping_rain_water

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

找左右两边最大值中小者，减去本身高度
*/

// 方法一：暴力法
// 对于每个柱子，找到左边的最高点和右边的最高点，取二者较小者，与当前柱子相减，即为能接到的雨水
func trap1(height []int) int {
	var ans int
	for i := 1; i < len(height)-1; i++ {
		var leftMax int
		for j := 0; j <= i; j++ {
			leftMax = max(leftMax, height[j])
		}
		var rightMax int
		for j := i; j < len(height); j++ {
			rightMax = max(rightMax, height[j])
		}
		ans += min(leftMax, rightMax) - height[i]
	}
	return ans
}

// 方法二：动态规划
// 暴力法中，来回遍历多次height数组，其中有多个重复计算的部分，可以通过记忆法保存结果
// 即：对于每个柱子height[i]，保存leftMax[i]和rightMax[i]
func trap2(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	lefts := make([]int, n)
	lefts[0] = height[0]
	for i := 1; i < n; i++ {
		lefts[i] = max(height[i], lefts[i-1])
	}
	rights := make([]int, n)
	rights[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rights[i] = max(height[i], rights[i+1])
	}
	var ans = 0
	for i := 1; i < n-1; i++ {
		ans += min(lefts[i], rights[i]) - height[i]
	}
	return ans
}

// 方法三：双指针，优化动态规划方法的空间复杂度
// 动态规划过程中，只依赖leftMax[i]和rightMax[i]
// 因此递推过程中，只需要定义leftMax和rightMax2个变量即可
func trap3(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	ans := 0
	for left < right {
		if height[left] < height[right] {
			leftMax = max(leftMax, height[left])
			ans += leftMax - height[left]
			left++
		} else {
			rightMax = max(rightMax, height[right])
			ans += rightMax - height[right]
			right--
		}
	}
	return ans
}

// 方法四：单调栈
// 当前柱子小于栈顶元素，说明无法形成凹槽，将当前柱子入栈；反之，能形成凹槽，将元素弹出，计算结果
func trap4(height []int) int {
	var stack []int
	var ans int
	// 遍历每个柱子，入栈
	for i := 0; i < len(height); i++ {
		for len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
			bottomIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 如果栈顶元素相等，一直出栈
			// 接雨水的量相同，只需要乘上系数即可
			for len(stack) > 0 && height[stack[len(stack)-1]] == height[bottomIndex] {
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 {
				// height[i]是右边界，height[stack[len(stack)-1]是左边界
				// height[bottomIndex]是最矮的柱子，能形成凹槽的点
				// i - stack[len(stack)-1] - 1是系数
				ans += (min(height[i], height[stack[len(stack)-1]]) - height[bottomIndex]) * (i - stack[len(stack)-1] - 1)
			}
		}
		stack = append(stack, i)
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
