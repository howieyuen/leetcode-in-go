package _496_next_greater_element_i

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var stack = make([]int, 0, len(nums2))
	var greater = make(map[int]int)
	for i := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < nums2[i] {
			key := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			greater[key] = nums2[i]
		}
		stack = append(stack, nums2[i])
	}
	var ans = make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		ans[i] = greater[nums1[i]]
		if ans[i] == 0 {
			ans[i] = -1
		}
	}
	return ans
}
