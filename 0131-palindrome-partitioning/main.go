package _131_palindrome_partitioning

func partition1(s string) [][]string {
	var ans [][]string
	var dp = make([][]bool, len(s))
	for right := range dp {
		dp[right] = make([]bool, len(s))
		for left := 0; left <= right; left++ {
			if s[left] == s[right] && ((right-left) <= 2 || dp[left+1][right-1]) {
				dp[left][right] = true
			}
		}
	}
	var dfs func(left int, partition []string)
	dfs = func(left int, partition []string) {
		if left == len(s) {
			var tmp = make([]string, len(partition))
			copy(tmp, partition)
			ans = append(ans, tmp)
			return
		}
		for right := left; right < len(s); right++ {
			if !dp[left][right] {
				continue
			}
			partition = append(partition, s[left:right+1])
			dfs(right+1, partition)
			partition = partition[:len(partition)-1]
		}
	}
	dfs(0, nil)
	return ans
}

func partition(s string) [][]string {
	var ans [][]string
	var dfs func(left int, partition []string)
	dfs = func(left int, partition []string) {
		if left == len(s) {
			var tmp = make([]string, len(partition))
			copy(tmp, partition)
			ans = append(ans, tmp)
			return
		}
		for right := left; right < len(s); right++ {
			if !isPalindrome(s[left : right+1]) {
				continue
			}
			partition = append(partition, s[left:right+1])
			dfs(right+1, partition)
			partition = partition[:len(partition)-1]
		}
	}
	dfs(0, nil)
	return ans
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
