package _268_missing_number

// 类似只出现一次的数，只要把下标i和num[i]都做一次异或，结果就是只出现的下标
func missingNumber(nums []int) int {
	var res = len(nums)
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
		res ^= i
	}
	return res
}
