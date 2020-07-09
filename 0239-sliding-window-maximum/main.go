package _239_sliding_window_maximum

// 首先将nums数组，切分成k个元素的子数组，最后一个可能不足k个
// eg:   [1, 3, -1, -3, 5, 3, 6, 7]
//       [1, 3, -1][-3, 5, 3][6, 7]
// 在移动窗口时，i,j分别表示开始和结束元素的下标，会出现两种情况：
// 情况1：i,j在同一个子数组内
// 针对这种情况，只需要建立left数组，从左向右，用于保存到每个子数组的开始位置到j的最大值
// left[i] = max(left[i-1], nums[i])
// 情况2：i,j在横跨2个子数组
// 针对这种情况，除了需要left数组，还需要right数组，从右向左，保存每个子数组的结束位置到j的最大值
// right[i] = max(right[i+1], nums[i])
// 每个窗口的最大元素，为left的尾巴和right的头，此二者的最大值
// 即: res[i] = max(left[i+k-1], right[i])
func maxSlidingWindow1(nums []int, k int) []int {
	n := len(nums)
	if n*k == 0 {
		return nil
	}
	if n == 1 {
		return nums
	}
	var left = make([]int, n)
	left[0] = nums[0]
	var right = make([]int, n)
	right[n-1] = nums[n-1]
	for i := 1; i < n; i++ {
		if i%k == 0 { // 子数组的起始位置
			left[i] = nums[i]
		} else {
			left[i] = max(nums[i], left[i-1])
		}
		j := n - i - 1
		if (j+1)%k == 0 { // 子数组的结束位置
			right[j] = nums[j]
		} else {
			right[j] = max(nums[j], right[j+1])
		}
	}
	var res = make([]int, n-k+1)
	for i := 0; i < n-k+1; i++ {
		res[i] = max(left[i+k-1], right[i])
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 双端队列，保存下标，保持队列指向元素单调递减
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < 2 {
		return nums
	}
	var res = make([]int, 0, len(nums)-k+1)
	var dequeue = make([]int, 0, len(nums))
	for i := range nums {
		// 从队尾向对头，与当前元素比较，保持队头指向的元素值最大
		for len(dequeue) > 0 && nums[dequeue[len(dequeue)-1]] <= nums[i] {
			dequeue = dequeue[:len(dequeue)-1]
		}
		// 当前元素入队尾
		dequeue = append(dequeue, i)
		// 超过k个元素已经比较，表示队头元素可能因为窗口移动过远，不得不丢弃
		if i >= k && dequeue[0] <= i-k {
			dequeue = dequeue[1:]
		}
		// 窗口内元素已满k个，保存最大值
		if i >= k-1 {
			res = append(res, nums[dequeue[0]])
		}
	}
	return res
}
