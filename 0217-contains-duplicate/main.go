package _217_contains_duplicate

func containsDuplicate(nums []int) bool {
	maps := make(map[int]int, len(nums))
	for i := range nums {
		if _, ok := maps[nums[i]]; !ok {
			maps[nums[i]] = 1
		} else {
			return true
		}
	}
	return false
}
