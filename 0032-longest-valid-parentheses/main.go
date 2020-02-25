package _032_longest_valid_parentheses

/*
	有效括号的最长字串
	用数组模拟进栈出栈操作，保存括号字符串的下标，下标之差表示字串长度
*/
func longestValidParentheses(s string) int {
	var stack []int
	max := 0
	stack = append(stack, -1)
	for i := range s {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				max = maxInt(max, i-stack[len(stack)-1])
			}
		}
	}
	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
