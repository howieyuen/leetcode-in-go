package main

func maxArea(height []int) int {
	maxArea := -1
	for i, j := 0, len(height)-1; i < j; {
		maxArea = max(maxArea, min(height[i], height[j])*(j-i))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxArea
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
