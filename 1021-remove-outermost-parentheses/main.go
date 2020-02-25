package _021_remove_outermost_parentheses

import (
	`strings`
)

func removeOuterParentheses(S string) string {
	count := 0
	var sb strings.Builder
	for _, v := range S {
		if v == '(' {
			if count > 0 {
				sb.WriteRune(v)
			}
			count++
		}
		if v == ')' {
			count--
			if count > 0 {
				sb.WriteRune(v)
			}
		}
	}
	return sb.String()
}

func removeOuterParentheses1(s string) string {
	count := 0
	ret := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			count -= 1
		}
		if count != 0 {
			ret = append(ret, s[i])
		}
		if s[i] == '(' {
			count += 1
		}
	}
	return string(ret)
}
