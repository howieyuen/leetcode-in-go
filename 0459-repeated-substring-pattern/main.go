package _459_repeated_substring_pattern

import (
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	count := 1
	for i := n / 2; i > 0; i-- {
		if n%i != 0 {
			continue
		}
		count = n / i
		if s == strings.Repeat(s[:i], count) {
			return true
		}
	}
	return false
}
