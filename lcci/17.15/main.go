package _7_15

import (
	"sort"
	"strings"
)

func longestWord(words []string) string {
	var dict = make(map[string]bool)
	var maxLen = 0
	for _, word := range words {
		if len(word) > maxLen {
			maxLen = len(word)
		}
		dict[word] = true
	}
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) == len(words[j]) {
			return strings.Compare(words[i], words[j]) < 0
		}
		return len(words[i]) > len(words[j])
	})
	
	var dfs func(word string) bool
	dfs = func(word string) bool {
		if word == "" {
			return true
		}
		for split := 1; split <= len(word); split++ {
			if _, ok := dict[word[:split]]; ok {
				if dfs(word[split:]) {
					return true
				}
			}
		}
		return false
	}
	
	for _, word := range words {
		delete(dict, word)
		if dfs(word) {
			return word
		}
	}
	return ""
}
