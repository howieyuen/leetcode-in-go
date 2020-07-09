package _7_13

func respace(dictionary []string, sentence string) int {
	var dp = make([]int, len(sentence)+1)
	for i := 1; i <= len(sentence); i++ {
		dp[i] = dp[i-1] + 1
		for _, word := range dictionary {
			if len(word) <= i {
				if sentence[i-len(word):i] == word {
					dp[i] = min(dp[i], dp[i-len(word)])
				}
			}
		}
	}
	return dp[len(sentence)]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
