package _553_optimal_division

import (
	"math"
	"strconv"
)

func optimalDivision(nums []int) string {
	if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}
	if len(nums) == 2 {
		return strconv.Itoa(nums[0]) + "/" + strconv.Itoa(nums[1])
	}
	var ans = strconv.Itoa(nums[0]) + "/(" + strconv.Itoa(nums[1])
	for i := 2; i < len(nums); i++ {
		ans += "/" + strconv.Itoa(nums[i])
	}
	ans += ")"
	return ans
}

type T struct {
	minVal, maxVal float64
	minStr, maxStr string
}

func optimalDivision1(nums []int) string {
	memo := make([][]*T, len(nums))
	for i := range memo {
		memo[i] = make([]*T, len(nums))
	}
	t := optimal(nums, 0, len(nums)-1, memo)
	return t.maxStr
}

func optimal(nums []int, start, end int, memo [][]*T) *T {
	if memo[start][end] != nil {
		return memo[start][end]
	}
	t := new(T)
	if start == end {
		t.maxVal = float64(nums[start])
		t.minVal = float64(nums[start])
		t.maxStr = strconv.Itoa(nums[start])
		t.minStr = strconv.Itoa(nums[start])
		return t
	}
	t.maxVal = math.SmallestNonzeroFloat64
	t.minVal = math.MaxFloat64
	t.maxStr = ""
	t.minStr = ""
	for i := start; i < end; i++ {
		left := optimal(nums, start, i, memo)
		right := optimal(nums, i+1, end, memo)
		if t.minVal > left.minVal/right.maxVal {
			t.minVal = left.minVal / right.maxVal
			t.minStr = left.minStr + "/"
			if i+1 != end {
				t.minStr += "("
			}
			t.minStr += right.maxStr
			if i+1 != end {
				t.minStr += ")"
			}
		}
		if t.maxVal < left.maxVal/right.minVal {
			t.maxVal = left.maxVal / right.minVal
			t.maxStr = left.maxStr + "/"
			if i+1 != end {
				t.maxStr += "("
			}
			t.maxStr += right.minStr
			if i+1 != end {
				t.maxStr += ")"
			}
		}
	}
	memo[start][end] = t
	return t
}
