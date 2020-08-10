package _696_count_binary_substrings

func countBinarySubstrings(s string) int {
	var index, last, ans int
	n := len(s)
	for index < n {
		char := s[index]
		count := 0
		for index < n && char == s[index] {
			index++
			count++
		}
		ans += min(count, last)
		last = count
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
