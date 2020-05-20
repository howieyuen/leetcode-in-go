package _371_find_the_longest_substring_containing_vowels_in_even_counts

/*
给你一个字符串 s ，请你返回满足以下条件的最长子字符串的长度：每个元音字母，即 'a'，'e'，'i'，'o'，'u' ，在子字符串中都恰好出现了偶数次。

第一步：暴力破解，枚举所有子串，判断元音字母的个数是否都满足偶数这一要求
第二步：前缀和优化枚举时间，暴力会超时，原因是重复计算，如果能把子串s[i:j]的元音个数记录下来，优化时间
       本题要记录5个字母个数，可以用map保存；又只需要记录奇偶性，相当于0和1的变化，那么可以优化空间，用5个bit表示
       'a': 00001，'a'出现奇数次，其他出现偶数次
       'e': 00010，'e'出现奇数次，其他出现偶数次
       'i': 00100，'i'出现奇数次，其他出现偶数次
       'o': 01000，'o'出现奇数次，其他出现偶数次
       'u': 10000，'u'出现奇数次，其他出现偶数次
第三步：保存了中间状态，剩下的就是记录最大长度，只要在遍历字符串时，维护中间状态和字符下标的对应关系
*/
func findTheLongestSubstring(s string) int {
	var pos = make([]int, 1<<5)
	var ans = 0
	for i := range pos {
		pos[i] = -1
	}
	pos[0] = 0
	var status = 0
	for i := range s {
		switch s[i] {
		case 'a':
			status ^= 1 << 0
		case 'e':
			status ^= 1 << 1
		case 'i':
			status ^= 1 << 2
		case 'o':
			status ^= 1 << 3
		case 'u':
			status ^= 1 << 4
		}
		if pos[status] >= 0 { // 中间状态再次出现，元音字母出现偶数次
			ans = max(ans, i+1-pos[status])
		} else { // 中间状态首次出现，显然包含的奇数个元音字母
			pos[status] = i + 1
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
