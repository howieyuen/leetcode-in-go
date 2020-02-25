package _030_substring_with_concatenation_of_all_words

import (
	`fmt`
	"testing"
)

func Test_findSubstring(t *testing.T) {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	fmt.Println(findSubstring(s, words))
	s = "wordgoodgoodgoodbestword"
	words = []string{"word", "good", "best", "word"}
	fmt.Println(findSubstring(s, words))
	s = ""
	words = []string{}
	fmt.Println(findSubstring(s, words))
}
