package main

func isValid(s string) bool {
	brackets := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	for i := range s {
		if left, ok := brackets[s[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != left {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
