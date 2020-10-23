package _249_minimum_remove_to_make_valid_parentheses

func minRemoveToMakeValid(s string) string {
	chars := []byte(s)
	var stack []int
	for i := range chars {
		switch chars[i] {
		case '(':
			stack = append(stack, i)
		case ')':
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				chars[i] = ' '
			}
		}
	}
	for i := range stack {
		chars[stack[i]] = ' '
	}
	start := 0
	for i := range chars {
		if chars[i] != ' ' {
			chars[start] = chars[i]
			start++
		}
	}
	return string(chars[:start])
}
