package _055_jump_game

/*
45. 计算最少jump次数
55. 判断能否jump到最后

从最后一个位置反向推理，前一个位置的下标i，加上对应的值，是能到达的最远位置
如果比lastPos大，表示lastPos能到达，继续反向推理
*/
func canJump(nums []int) bool {
	lastPos := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if lastPos <= nums[i]+i {
			lastPos = i
		}
	}
	return lastPos == 0
}

/*
正向判断
维护一个变量：最远位置maxPos，每个可判断的位置，都在maxPos之内
如果i+num[i]>maxPos，最远位置更新，maxPos>=n-1,表示能到达，反之不能
*/
func canJump1(nums []int) bool {
	n := len(nums)
	maxPos := 0
	for i := range nums {
		if i <= maxPos {
			if i+nums[i] > maxPos {
				maxPos = i + nums[i]
			}
			if maxPos >= n-1 {
				return true
			}
		}
	}
	return false
}
