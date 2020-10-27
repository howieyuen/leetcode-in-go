package _392_longest_happy_prefix

// Rabin-Karp 字符串编码
// 前缀增长，末尾多一位，x*10 + y，即原编码值*base + 新字符的编码值
// 后缀增长，前面多一位，x + y^(i-1)，即原编码值 + 新字符的编码值 * n^(i-1)
func longestPrefix(s string) string {
	// 取模
	const mod = int(1e9 + 7)
	// 前缀编码和后缀编码
	preCode, postCode := 0, 0
	// 编码的最大长度
	n := len(s) - 1
	// 编码基数
	base := 31
	// 幂
	pow := 1
	// 最长前缀的最后下标
	happy := -1
	for i := 0; i < n; i++ {
		pre := int(s[i] - 'a')
		post := int(s[n-i] - 'a')
		preCode = (preCode*base + pre) % mod
		postCode = (post*pow + postCode) % mod
		pow = pow * base % mod
		if preCode == postCode {
			happy = i
		}
	}
	if happy == -1 {
		return ""
	}
	return s[:happy+1]
}

func longestPrefix1(s string) string {
	n := len(s)
	for d := n - 1; d > 0; d-- {
		if s[0:d] == s[n-d:] {
			return s[0:d]
		}
	}
	return ""
}
