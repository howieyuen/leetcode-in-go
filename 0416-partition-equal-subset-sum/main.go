package _416_partition_equal_subset_sum

// 0-1背包问题
// dp[i][target]表示nums[0, i]区间内是否能找到和为target的组合
// 对于每个nums[i]，如果nums[i] <= target，可以选择or不选，但只要有一个为true，dp[i][target]=true
//      dp[i][target] = dp[i-1][target] || dp[i][target-nums[i]]
// 如果nums[i] > target，只能不选，故：
//      dp[i][target] = dp[i-1][target]
func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	
	sum := 0
	maxNum := 0
	for _, num := range nums {
		sum += num
		if num > maxNum {
			maxNum = num
		}
	}
	// 和为奇数
	if sum%2 != 0 {
		return false
	}
	
	// 最大值超过和的一半
	target := sum / 2
	if maxNum > target {
		return false
	}
	
	var dp = make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, target+1)
		// nums数组均为正数，如果不选取任何nums[i]，则被选取的正整数的和等于0
		dp[i][0] = true
	}
	// 只有一个元素nums[0]可取时，构成初始值dp[0][nums[0]]为true
	dp[0][nums[0]] = true
	for i := 1; i < n; i++ {
		for j := 1; j <= target; j++ {
			if nums[i] <= j {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
			// 剪枝
			if dp[i][target] {
				return true
			}
		}
	}
	return dp[n-1][target]
}

func canPartition1(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	
	sum := 0
	maxNum := 0
	for _, num := range nums {
		sum += num
		if num > maxNum {
			maxNum = num
		}
	}
	// 和为奇数
	if sum%2 != 0 {
		return false
	}
	
	// 最大值超过和的一半
	target := sum / 2
	if maxNum > target {
		return false
	}
	
	var dp = make([]bool, target+1)
	dp[0] = true
	dp[nums[0]] = true
	for i := 1; i < len(nums); i++ {
		// 采用逆序，如果采用正序dp[j-nums[i]]会被之前的操作更新为新值
		// dp[j](new) = dp[j](old) || dp[j - nums[i]](old)
		for j := target; j >= nums[i]; j-- {
			if dp[target] {
				return true
			}
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[target]
}