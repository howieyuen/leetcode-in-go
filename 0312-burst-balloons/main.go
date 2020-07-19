package _312_burst_balloons

// 方法一：记忆法递归
// 每次戳完气球后，本不相邻的气球变成相邻，难以处理，因此将整个过程看作是插入硬币数为1的气球
// 对于整个过程，就是对每个子区间选择最优解，合并子最优达到整体最优
// 枚举开区间(i, j)所有可能的切分位置mid，寻找左区间(i, mid)和右区间(mid, j)最优
// 递归方程为：dp[i][j] = num[i]*num[mid]*nums[j] + solve(i,mid) + solve(mid, j)
func maxCoins1(nums []int) int {
	n := len(nums)
	
	values := make([]int, n+2)
	for i := 1; i <= n; i++ {
		values[i] = nums[i-1]
	}
	values[0], values[n+1] = 1, 1
	
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	
	return solve(0, n+1, values, dp)
}

func solve(left, right int, values []int, dp [][]int) int {
	if left+1 >= right {
		return 0
	}
	if dp[left][right] != -1 {
		return dp[left][right]
	}
	for mid := left + 1; mid < right; mid++ {
		sum := values[left] * values[mid] * values[right]
		sum += solve(left, mid, values, dp) + solve(mid, right, values, dp)
		dp[left][right] = max(dp[left][right], sum)
	}
	return dp[left][right]
}

// 方法二：动态规划
// 即把递归的写法改成循环
func maxCoins2(nums []int) int {
	n := len(nums)
	values := make([]int, n+2)
	for i := 1; i <= n; i++ {
		values[i] = nums[i-1]
	}
	values[0], values[n+1] = 1, 1
	
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
		
	}
	
	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j < n+2; j++ {
			for k := i + 1; k < j; k++ {
				sum := values[i] * values[k] * values[j]
				sum += dp[i][k] + dp[k][j]
				dp[i][j] = max(dp[i][j], sum)
			}
		}
	}
	return dp[0][n+1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
