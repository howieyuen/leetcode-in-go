package _088_merge_sorted_array

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	right := m + n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[right] = nums1[i]
			i--
		} else {
			nums1[right] = nums2[j]
			j--
		}
		right--
	}
	for i >= 0 {
		nums1[right] = nums1[i]
		i--
		right--
	}
	for j >= 0 {
		nums1[right] = nums2[j]
		j--
		right--
	}
}
