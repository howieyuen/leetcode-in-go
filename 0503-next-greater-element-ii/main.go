package _503_next_greater_element_ii

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	var stack = make([]int, 0, n)
	var ans = make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	for i := 0; i < n*2; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[j] = nums[i%n]
		}
		if i < n {
			stack = append(stack, i)
		}
	}
	return ans
}
