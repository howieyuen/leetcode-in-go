package _448_find_all_numbers_disappeared_in_an_array

func findDisappearedNumbers(nums []int) []int {
	var ans []int
	for i := 0; i < len(nums); i++ {
		for nums[i] != nums[nums[i]-1] {
			swap(nums, i, nums[i]-1)
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			ans = append(ans, i+1)
		}
	}
	return ans
}

func swap(nums []int, i, j int) {
	if i != j {
		nums[i] = nums[i] ^ nums[j]
		nums[j] = nums[i] ^ nums[j]
		nums[i] = nums[i] ^ nums[j]
	}
}
