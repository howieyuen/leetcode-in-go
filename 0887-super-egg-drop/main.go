package _887_super_egg_drop

import (
	"math"
)

/*
k是鸡蛋数，n表示楼层高度，
求找到安全楼层，需要扔鸡蛋的最少次数

反向思考，有i个鸡蛋，扔鸡蛋j次，能检测到的楼层数m，若m大于等于n，j即为所求
dp[i][j]表示用i鸡蛋，扔j次，检测到的最高楼层数
鸡蛋只有碎和不碎两种状态，j是递增，因此，扔j次的结果，依赖扔j-1次的结果
dp[i][j] = dp[i-1][j-1] + dp[i][j-1] + 1
*/
func superEggDrop(k int, n int) int {
	var dp = make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for j := 1; j <= n; j++ {
		for i := 1; i <= k; i++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j-1] + 1
			if dp[i][j] >= n {
				return j
			}
		}
	}
	return n
}

/*
可以采用动态规划的思想，转换成一个函数f(n,k)，代表的就是最优解的最大尝试次数，假设我们第一个鸡蛋扔出的位置在第x层(1≤x≤n)，会出现两种情况：

第一个鸡蛋没碎，剩下n-x层楼，剩余k个鸡蛋，函数转变为

f(n-x,k) + 1,1≤x≤n f(n−x,k)+1,1≤x≤n

第一个鸡蛋碎了，剩余x-1层楼，剩余k-1个鸡蛋，函数转变为

f(x-1,k-1) + 1,1≤x≤n f(x−1,k−1)+1,1≤x≤n

整体而言，我们要求出的是n层楼，k个鸡蛋条件下，能找到出安全楼层的最小扔鸡蛋次数，所以这个题目的状态转移方程式如下：

f(n，k) = min(max(f(n-x，k) + 1,f(x-1,k-1) + 1)),1<=x<=M
f(n，k) = min(max(f(n−x，k) + 1,f(x−1,k−1) + 1)),1<=x<=M
*/
func superEggDrop1(k int, n int) int {
	if k == 0 || n == 0 {
		return 0
	}
	if k == 1 {
		return n
	}
	if n == 1 {
		return 1
	}
	
	dp := make([][]int, k+1)
	for i := 0; i < k+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= n; i++ {
		dp[1][i] = i
	}
	for i := 1; i <= k; i++ {
		dp[i][1] = 1
	}
	for i := 2; i <= k; i++ {
		for j := 2; j <= n; j++ {
			for x := 1; x <= j; x++ {
				if dp[i][j] == 0 {
					dp[i][j] = math.MaxInt32
				}
				dp[i][j] = min(dp[i][j], max(dp[i][j-x], dp[i-1][x-1])+1)
			}
		}
	}
	return dp[k][n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
