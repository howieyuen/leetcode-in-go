package _045_jump_game_ii

/*
45. 计算到达最远位置的最少的jump次数
55. 判断能否到达最后
*/
func jump(nums []int) int {
	maxPos := 0
	end := 0
	steps := 0
	for i := 0; i < len(nums)-1; i++ {
		maxPos = max(maxPos, nums[i]+i)
		if i == end {
			end = maxPos
			steps++
		}
	}
	return steps
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
