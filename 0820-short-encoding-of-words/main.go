package _820_short_encoding_of_words

import (
	"sort"
)

func minimumLengthEncoding(words []string) int {
	m := make(map[string]byte, len(words))
	result := 0
	for _, word := range words {
		if _, ok := m[word]; !ok {
			m[word] = 1
		}
	}
	for _, word := range words {
		for i := 1; i < len(word); i++ {
			delete(m, word[i:])
		}
	}
	for item := range m {
		result += len(item) + 1
	}
	return result
}

func minimumLengthEncoding1(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	root := new(Trie)
	ans := 0
	for _, word := range words {
		ans += root.insert(word)
	}
	return ans
}

type Trie struct {
	children [26]*Trie
}

func (t *Trie) insert(word string) int {
	var length = 0
	for i := len(word) - 1; i >= 0; i-- {
		index := word[i] - 'a'
		if t.children[index] == nil {
			length = len(word) + 1
			t.children[index] = new(Trie)
		}
		t = t.children[index]
	}
	return length
}
