package _169_majority_element

func majorityElement(nums []int) int {
	count := 0
	candidate := 0
	for i := range nums {
		if count == 0 {
			candidate = nums[i]
		}
		if nums[i] == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
