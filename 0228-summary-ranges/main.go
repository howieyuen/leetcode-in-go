package _228_summary_ranges

import (
	"strconv"
)

func summaryRanges(nums []int) []string {
	var ans []string
	for i, j := 0, 0; j < len(nums); j++ {
		if j+1 < len(nums) && nums[j]+1 == nums[j+1] {
			continue
		}
		if i == j {
			ans = append(ans, strconv.Itoa(nums[i]))
		} else {
			ans = append(ans, strconv.Itoa(nums[i])+"->"+strconv.Itoa(nums[j]))
		}
		i = j + 1
	}
	return ans
}
