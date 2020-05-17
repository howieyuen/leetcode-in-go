package _588_lcof_19

/*
模式中的字符'.'表示任意一个字符，而'*'表示它前面的字符可以出现任意次（含0次）
*/
func isMatch(s string, p string) bool {
	if s == p {
		return true
	}
	return match(s, 0, p, 0)
}

func match(s string, i int, p string, j int) bool {
	// 匹配完成
	if i >= len(s) && j >= len(p) {
		return true
	}
	// 匹配失败，模式串过短
	if i < len(s) && j >= len(p) {
		return false
	}
	// 模式串下个是'*'
	if j+1 < len(p) && p[j+1] == '*' {
		// 当前字符匹配
		if i < len(s) && (s[i] == p[j] || p[j] == '.') {
			// * 可用多次或者0次
			return match(s, i+1, p, j) || match(s, i, p, j+2)
		} else {
			// 当前字符串不匹配，借*消除前面的字符，即匹配0次
			return match(s, i, p, j+2)
		}
	}
	// 模式串下个不是'*'，正常匹配
	if i < len(s) && (s[i] == p[j] || p[j] == '.') {
		return match(s, i+1, p, j+1)
	}
	return false
}
