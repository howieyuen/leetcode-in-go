package _321_create_maximum_number

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	initialSize := 0
	if len(nums2) < k {
		initialSize = k - len(nums2)
	}
	var ans []int
	for size := initialSize; size <= k && size <= len(nums1); size++ {
		sub1 := maxSubSequence(nums1, size)
		sub2 := maxSubSequence(nums2, k-size)
		merged := merge(sub1, sub2)
		if lexicographicalLess(ans, merged) {
			ans = merged
		}
	}
	return ans
}

// 按照字典序降序合并
func merge(sub1, sub2 []int) []int {
	var merged = make([]int, len(sub1)+len(sub2))
	for i := range merged {
		if lexicographicalLess(sub1, sub2) {
			merged[i], sub2 = sub2[0], sub2[1:]
		} else {
			merged[i], sub1 = sub1[0], sub1[1:]
		}
	}
	return merged
}

// 判断字典序大小
func lexicographicalLess(sub1, sub2 []int) bool {
	for i := 0; i < len(sub1) && i < len(sub2); i++ {
		if sub1[i] != sub2[i] {
			return sub1[i] < sub2[i]
		}
	}
	return len(sub1) < len(sub2)
}

// 单调栈求指定长度的最大子序列
func maxSubSequence(nums []int, k int) []int {
	var sub []int
	for i := range nums {
		for len(sub) > 0 && nums[i] > sub[len(sub)-1] && len(sub)+len(nums)-i-1 >= k {
			sub = sub[:len(sub)-1]
		}
		if len(sub) < k {
			sub = append(sub, nums[i])
		}
	}
	return sub
}
