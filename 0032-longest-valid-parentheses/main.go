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

// 使用贪心算法
// '(' -> left++
// ')' -> right++
// 当left == right，找到一个有效的括号串
// 当left < right, 右括号多，从0开始
// 但存在一种例外，就是遍历的时候左括号的数量始终大于右括号的数量，即 (()
// 因此使用双向遍历，从右到左，贪心选择右括号，可以覆盖所有场景
func longestValidParentheses1(s string) int {
	left, right := 0, 0
	ans := 0
	for i := range s {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			ans = maxInt(ans, 2*left)
		} else if left < right {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			right++
		} else {
			left++
		}
		if left == right {
			ans = maxInt(ans, 2*right)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return ans
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
