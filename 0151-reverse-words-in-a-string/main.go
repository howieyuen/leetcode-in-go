package _151_reverse_words_in_a_string

import (
	`strings`
)

func reverseWords(s string) string {
	subs := strings.Split(s, " ")
	i := 0
	for j := 0; j < len(subs); j++ {
		if len(subs[j]) > 0 {
			subs[i] = subs[j]
			i++
		}
	}
	return reverse(subs[0:i])
}

func reverse(tmp []string) string {
	for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}
	return strings.Join(tmp, " ")
}
