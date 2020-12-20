package _316_remove_duplicate_letters

import (
	"bytes"
)

func removeDuplicateLetters(s string) string {
	b := []byte(s)
	var stack []byte
	for i, c := range b {
		// 去重
		if bytes.IndexByte(stack, c) != -1 {
			continue
		}
		// 单调栈栈顶字符大于当前字符，且此栈顶字符后续会再次出现，则弹栈
		for len(stack) > 0 && stack[len(stack)-1] > c && bytes.IndexByte(b[i+1:], stack[len(stack)-1]) != -1 {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, c)
	}
	return string(stack)
}
