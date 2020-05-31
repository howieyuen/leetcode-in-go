package _085_maximal_rectangle

/*
按层计算最大面积
*/
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	var heights = make([]int, len(matrix[0]))
	var ans int
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		ans = max(ans, increaseStack(heights))
	}
	return ans
}

func increaseStack(heights []int) int {
	var stack = make([]int, 0, len(heights))
	stack = append(stack, -1)
	last := 0
	var area int
	for i := range heights {
		for stack[last] != -1 && heights[stack[last]] >= heights[i] {
			h := heights[stack[last]]
			stack = stack[:last]
			last--
			w := i - stack[last] - 1
			area = max(area, h*w)
		}
		stack = append(stack, i)
		last++
	}
	for stack[last] != -1 {
		h := heights[stack[last]]
		stack = stack[:last]
		last--
		w := len(heights) - stack[last] - 1
		area = max(area, h*w)
	}
	return area
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
