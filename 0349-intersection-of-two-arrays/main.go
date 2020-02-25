package _349_intersection_of_two_arrays

func intersection(nums1 []int, nums2 []int) []int {
	var ans []int
	var maps1 = make(map[int]bool)
	for _, num := range nums1 {
		if _, ok := maps1[num]; !ok {
			maps1[num] = true
		}
	}
	var maps2 = make(map[int]bool)
	for _, num := range nums2 {
		if _, ok := maps2[num]; !ok {
			maps2[num] = true
		}
	}
	for key := range maps1 {
		if maps2[key] {
			ans = append(ans, key)
		}
	}
	return ans
}
