package _7_21

import (
	"container/list"
)

// 方法1：暴力法
// 对于柱子i，找到左边最高leftMax和右边最高柱子rightMax，形成凹槽
func trap(height []int) int {
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

// 方法2：记忆法
// 预先保存leftMax和rightMax
func trap1(height []int) int {
	if len(height) == 0 {
		return 0
	}
	n := len(height)
	var leftMax = make([]int, n)
	var rightMax = make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(height[i], leftMax[i-1])
	}
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(height[i], rightMax[i+1])
	}
	var ans int
	for i := 0; i < n; i++ {
		ans += min(leftMax[i], rightMax[i]) - height[i]
	}
	return ans
}

// 方法3：双指针——记忆法优化
// 每次遍历只需要leftMax[i]和rightMax[i]，递推过程只需要更新这2个变量
func trap2(height []int) int {
	i, j := 0, len(height)-1
	var leftMax, rightMax int
	var ans int
	for i < j {
		if height[i] < height[j] {
			leftMax = max(leftMax, height[i])
			ans += leftMax - height[i]
			i++
		} else {
			rightMax = max(rightMax, height[j])
			ans += rightMax - height[j]
			j--
		}
	}
	return ans
}

// 方法4：单调栈
// 栈中保存高度递减的下标序列，柱子i的高度小于栈顶元素，表示无法形成凹槽，入栈
// 反之，栈中元素弹出，计算形成的凹槽深度
func trap3(height []int) int {
	var stack = list.New()
	var ans int
	for i := range height {
		for stack.Len() > 0 && height[i] > height[stack.Back().Value.(int)] {
			bottomIndex := stack.Remove(stack.Back()).(int)
			for stack.Len() > 0 && height[bottomIndex] == height[stack.Back().Value.(int)] {
				stack.Remove(stack.Back())
			}
			if stack.Len() > 0 {
				leftIndex := stack.Back().Value.(int)
				ans += (min(height[i], height[leftIndex]) - height[bottomIndex]) * (i - leftIndex - 1)
			}
		}
		stack.PushBack(i)
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
