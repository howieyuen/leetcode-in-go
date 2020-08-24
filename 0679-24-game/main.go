package _679_24_game

import "math"

const (
	TARGET  = 24
	EPSILON = 1E-6
)

func judgePoint24(nums []int) bool {
	var list []float64
	for _, num := range nums {
		list = append(list, float64(num))
	}
	return solve(list)
}

func solve(nums []float64) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return math.Abs(nums[0]-TARGET) < EPSILON
	}
	
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j {
				var newNums []float64
				for k := 0; k < len(nums); k++ {
					if k != i && k != j {
						newNums = append(newNums, nums[k])
					}
				}
				// 处理 +、* 时，条件i < j 是为了避免重复计算，因为运算的结果是一样的
				if i < j && solve(append(newNums, nums[i]+nums[j])) {
					return true
				}
				if i < j && solve(append(newNums, nums[i]*nums[j])) {
					return true
				}
				if solve(append(newNums, nums[i]-nums[j])) {
					return true
				}
				if nums[j] > EPSILON && solve(append(newNums, nums[i]/nums[j])) {
					return true
				}
			}
		}
	}
	return false
}
