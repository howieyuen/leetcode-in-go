package _0

import (
	"strings"
)

/*
see https://leetcode-cn.com/problems/valid-number/
*/
func isNumber(s string) bool {
	s = strings.TrimLeft(s, " ")
	s = strings.TrimRight(s, " ")
	if s == "" {
		return false
	}
	split1 := strings.Index(s, "e")
	split2 := strings.Index(s, "E")
	if split1 != -1 && split2 != -1 {
		return false
	} else if split1 != -1 || split2 != -1 {
		split := split1 + split2 + 1
		left, right := s[:split], s[split+1:]
		if len(left) == 0 || len(right) == 0 {
			return false
		}
		if withFlag(left) && len(left[1:]) == 0 {
			return false
		}
		if withFlag(right) && len(right[1:]) == 0 {
			return false
		}
		return allDigitsWithDot(left) && allDigitsWithFlag(right)
	}
	return allDigitsWithDot(s)
}

func allDigitsWithDot(s string) bool {
	if split := strings.Index(s, "."); split != -1 {
		left, right := s[:split], s[split+1:]
		if len(left) == 0 && len(right) == 0 {
			return false
		}
		if len(left) == 1 && withFlag(left) && len(right) == 0 {
			return false
		}
		return allDigitsWithFlag(s[:split]) && allDigits(s[split+1:])
	}
	return allDigitsWithFlag(s)
}

func allDigitsWithFlag(s string) bool {
	if len(s) == 0 {
		return true
	}
	if withFlag(s) {
		return allDigits(s[1:])
	}
	return allDigits(s)
}

func withFlag(s string) bool {
	return s[0] == '+' || s[0] == '-'
}

func allDigits(s string) bool {
	var t uint8
	for i := range s {
		t = s[i] - '0'
		if t < 0 || t >= 10 {
			return false
		}
	}
	return true
}
