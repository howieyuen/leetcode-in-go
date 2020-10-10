package _438_longest_continuous_subarray_with_absolute_diff_less_than_or_equal_to_limit

// 暴力破解
func longestSubarray(nums []int, limit int) int {
	longest := 0
	for i := 0; i < len(nums); i++ {
		min, max := nums[i], nums[i]
		for j := i; j < len(nums); j++ {
			if nums[j] > max {
				max = nums[j]
			} else if nums[j] < min {
				min = nums[j]
			}
			if max-min <= limit {
				cur := j - i + 1
				if cur > longest {
					longest = cur
				}
			} else {
				break
			}
		}
		// 剪枝
		if longest >= len(nums)-i {
			break
		}
	}
	return longest
}

// 单调栈
func longestSubarray1(nums []int, limit int) int {
	// 单调递增栈保存最大值、单调递减栈保存最小值
	var maxStack, minStack []int
	size := 0
	for leftSide, rightSide := 0, 0; rightSide < len(nums); rightSide++ {
		// 维护单调递增栈
		for len(maxStack) > 0 && nums[rightSide] > nums[maxStack[len(maxStack)-1]] {
			maxStack = maxStack[:len(maxStack)-1]
		}
		maxStack = append(maxStack, rightSide)
		
		// 维护单调递减栈
		for len(minStack) > 0 && nums[rightSide] < nums[minStack[len(minStack)-1]] {
			minStack = minStack[:len(minStack)-1]
		}
		minStack = append(minStack, rightSide)
		
		// 取出最大值最小值
		max, min := nums[maxStack[0]], nums[minStack[0]]
		// 不满足条件，窗口缩小
		for max-min > limit {
			// 如果leftSide对应的是max
			if maxStack[0] == leftSide {
				maxStack = maxStack[1:]
			}
			// 如果leftSide对应的是min
			if minStack[0] == leftSide {
				minStack = minStack[1:]
			}
			// 修改窗口左侧，窗口缩小
			leftSide++
			max, min = nums[maxStack[0]], nums[minStack[0]]
		}
		
		// 更新结果
		curSize := rightSide - leftSide + 1
		if size < curSize {
			size = curSize
		}
	}
	return size
}
