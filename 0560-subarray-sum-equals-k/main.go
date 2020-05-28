package _560_subarray_sum_equals_k

/*
给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。

示例 1 :

输入:nums = [1,1,1], k = 2
输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
说明 :

数组的长度为 [1, 20,000]。
数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subarray-sum-equals-k
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func subarraySum(nums []int, k int) int {
	ans := 0
	store := make(map[int]int, 0)
	sum := 0
	store[0] = 1
	for i := range nums {
		sum += nums[i]
		if store[sum-k] > 0 {
			ans += store[sum-k]
		}
		store[sum]++
	}
	return ans
}

func subarraySum1(nums []int, k int) int {
	n := len(nums)
	var dp = make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] += dp[i-1] + nums[i-1]
	}
	ans := 0
	for i := 0; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[i]-dp[j] == k {
				ans++
			} else if dp[i]-dp[j] < k {
				continue
			}
		}
	}
	return ans
}

func subarraySum2(a []int, k int) int {
	var count int
	for i := 1; i < len(a); i++ {
		a[i] += a[i-1]
		
	}
	for i := 0; i < len(a); i++ {
		if a[i] == k {
			count++
		}
		for j := i + 1; j < len(a); j++ {
			if a[j]-a[i] == k {
				count++
			}
		}
	}
	return count
}
