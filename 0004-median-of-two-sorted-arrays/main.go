package _004_Median_of_Two_Sorted_Arrays

/*
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}
	iMin, iMax := 0, m
	half := (m + n + 1) / 2
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := half - i
		if i < iMax && nums1[i] < nums2[j-1] { // i需要向右移动
			iMin = i + 1
		} else if i > iMin && nums1[i-1] > nums2[j] { // i需要向左移动
			iMax = i - 1
		} else {
			maxLeft := 0
			if i == 0 { // num1的数较大，中位数靠近num2的左侧最大数
				maxLeft = nums2[j-1]
			} else if j == 0 { // num2的数较大，中位数靠近num1左侧最大数
				maxLeft = nums1[i-1]
			} else {
				maxLeft = max(nums1[i-1], nums2[j-1])
			}
			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}
			minRight := 0
			if i == m { // num1的数较小，中位数靠近num2的右侧最小数
				minRight = nums2[j]
			} else if j == n { // num2的数较小，中位数靠近num1的右侧最小数
				minRight = nums1[i]
			} else {
				minRight = min(nums1[i], nums2[j])
			}
			return float64(maxLeft+minRight) / 2.0
		}
	}
	return 0.0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
