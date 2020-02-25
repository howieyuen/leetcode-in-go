package _189_rotate_array

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	if k == 0 {
		return
	}
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func rotate1(nums []int, k int) {
	n := len(nums)
	k %= n
	if k == 0 {
		return
	}
	for i := 0; i < k; i++ {
		pre := nums[n-1]
		for j := 0; j < n; j++ {
			nums[j], pre = pre, nums[j]
		}
	}
}

func rotate2(nums []int, k int) {
	n := len(nums)
	k %= n
	if k == 0 {
		return
	}
	start := 0
	for i := 0; i < n; {
		cur := start
		pre := nums[start]
		for start != cur {
			next := (cur + k) % n
			nums[next], pre = pre, nums[next]
			i++
		}
		start++
	}
}

// 1, 2, 3, 4, 5, 6, 7
// 5, 6, 7, 1, 2, 3, 4
