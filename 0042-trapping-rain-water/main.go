package _042_trapping_rain_water

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

找左右两边最大值中小者，减去本身高度
*/
func trap(height []int) int {
	n := len(height)
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

func trap1(height []int) int {
	left := 0
	right := len(height) - 1
	ans := 0
	leftMax := 0
	rightMax := 0
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				ans += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				ans += rightMax - height[right]
			}
			right--
		}
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
