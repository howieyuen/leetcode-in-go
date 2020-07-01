package _718_maximum_length_of_repeated_subarray

//  j依赖j-1的状态，如果从前向后，j-1先更新了，结果就不对了
func findLength(A []int, B []int) int {
	dp := make([]int, len(B)+1)
	ans := 0
	for i := 0; i < len(A); i++ {
		for j := len(B) - 1; j >= 0; j-- {
			if A[i] == B[j] {
				dp[j+1] = 1 + dp[j]
				ans = max(ans, dp[j+1])
			} else {
				dp[j+1] = 0
			}
		}
	}
	return ans
}

func findLength2(A []int, B []int) int {
	dp := make([]int, len(B))
	ans := 0
	for i := 0; i < len(A); i++ {
		prev := 0
		for j := 0; j < len(B); j++ {
			ori := dp[j]
			if A[i] == B[j] {
				dp[j] = 1 + prev
				ans = max(ans, dp[j])
			}
			prev = ori
		}
	}
	return ans
}

func findLength1(a []int, b []int) int {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}
	var dp = make([][]int, len(a)+1)
	for i := range dp {
		dp[i] = make([]int, len(b)+1)
	}
	var res = 0
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				res = max(res, dp[i][j])
			}
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
