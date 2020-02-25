package _219_contains_duplicate_ii

func containsNearbyDuplicate(nums []int, k int) bool {
	var maps = make(map[int]bool)
	for i := range nums {
		if maps[nums[i]] {
			return true
		}
		maps[nums[i]] = true
		if len(maps) > k {
			delete(maps, nums[i-k])
		}
	}
	return false
}
