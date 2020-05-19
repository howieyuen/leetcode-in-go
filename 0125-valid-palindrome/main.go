package _125_valid_palindrome

import (
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		for i < j {
			if s[i] >= 'a' && s[i] <= 'z' || s[i] >= '0' && s[i] <= '9' {
				break
			} else {
				i++
			}
		}
		
		for i < j {
			if s[j] >= 'a' && s[j] <= 'z' || s[j] >= '0' && s[j] <= '9' {
				break
			} else {
				j--
			}
		}
		if i > j {
			break
		}
		if s[i] == s[j] {
			continue
		}
		return false
	}
	return true
}
