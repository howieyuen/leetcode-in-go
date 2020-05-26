package _058_length_of_last_word

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	var end = len(s) - 1
	for end >= 0 && s[end] == ' ' {
		end--
	}
	if end < 0 {
		return 0
	}
	var start = end
	for start >= 0 && s[start] != ' ' {
		start--
	}
	return end - start
}

func lengthOfLastWord1(s string) int {
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return 0
	}
	sArray := strings.Split(s, " ")
	return len(sArray[len(sArray)-1])
}
