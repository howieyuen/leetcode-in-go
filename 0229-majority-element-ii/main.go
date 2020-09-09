package _229_majority_element_ii

func majorityElement(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	
	var ans []int
	first, count1 := nums[0], 0
	second, count2 := nums[0], 0
	for _, num := range nums {
		switch {
		case num == first:
			count1++
		case num == second:
			count2++
		case count1 == 0:
			first = num
			count1++
		case count2 == 0:
			second = num
			count2++
		default:
			count1--
			count2--
		}
	}
	count1, count2 = 0, 0
	for _, num := range nums {
		switch num {
		case first:
			count1++
		case second:
			count2++
		}
	}
	if count1 > len(nums)/3 {
		ans = append(ans, first)
	}
	if count2 > len(nums)/3 {
		ans = append(ans, second)
	}
	return ans
}
