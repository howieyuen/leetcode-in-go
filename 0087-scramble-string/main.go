package _087_scramble_string

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	var letters = make([]int, 26)
	for i := range s1 {
		letters[s1[i]-'a']++
		letters[s2[i]-'a']--
	}
	for i := range letters {
		if letters[i] != 0 {
			return false
		}
	}
	for i := 1; i < len(s1); i++ {
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}
		if isScramble(s1[:i], s2[len(s2)-i:]) && isScramble(s1[i:], s2[:len(s2)-i]) {
			return true
		}
	}
	return false
}

/*
dp[len][i][j]来表示s1[i:i+len]和s2[j:j+len]两个字符串是否满足条件。
换句话说，就是s1从i开始的len个字符是否能转换为S2从j开始的len个字符

假设左半部分长度是q，
左s1对左s2，右s1对右s2: dp[len][i][j]=dp[q][i][j]&&dp[len-q][i+q][j+q]
左s1对右s2，右s1对左s2: dp[len][i][j]=dp[q][i][j+len-q]&&dp[len-q][i+q][j]
*/
func isScramble1(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	var letters = make([]int, 26)
	for i := range s1 {
		letters[s1[i]-'a']++
		letters[s2[i]-'a']--
	}
	for i := range letters {
		if letters[i] != 0 {
			return false
		}
	}
	n := len(s1)
	var dp = make([][][]bool, n)
	for i := range dp {
		dp[i] = make([][]bool, n)
		for j := range dp[i] {
			dp[i][j] = make([]bool, n+1)
			dp[i][j][1] = s1[i] == s2[j]
		}
	}
	for l := 2; l <= n; l++ {
		for i := 0; i+l <= n; i++ {
			for j := 0; j+l <= n; j++ {
				for k := 1; k < l; k++ {
					if dp[i][j][k] && dp[i+k][j+k][l-k] {
						dp[i][j][l] = true
						break
					}
					if dp[i][j+l-k][k] && dp[i+k][j][l-k] {
						dp[i][j][l] = true
						break
					}
				}
			}
		}
	}
	return dp[0][0][n]
}
