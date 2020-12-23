package _523_continuous_subarray_sum

// 前缀和，O(n^2)的时间复杂度，O(n)的空间复杂度，计算每个子数组的和
func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	var preSum = make([]int, n)
	preSum[0] = nums[0]
	for i := 1; i < n; i++ {
		preSum[i] = preSum[i-1] + nums[i]
	}
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			tmp := preSum[j] - preSum[i] + nums[i]
			if tmp == k || (k != 0 && tmp%k == 0) {
				return true
			}
		}
	}
	return false
}

// 用map记录前i个元素的和，模k之后的值sum与下标i的关系
// 计算前缀和sum%=k，如果sum在map存在
// 表示之间的j个元素的和对k的余数也是sum，两个相减就是k的倍数
func checkSubarraySum1(nums []int, k int) bool {
	var sum int
	var dict = make(map[int]int)
	dict[0] = -1
	for i := range nums {
		sum += nums[i]
		if k != 0 {
			sum %= k
		}
		if j, ok := dict[sum]; ok {
			if i-j > 1 {
				return true
			}
		} else {
			dict[sum] = i
		}
	}
	return false
}
