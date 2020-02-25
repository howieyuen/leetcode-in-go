package _030_substring_with_concatenation_of_all_words

func findSubstring(s string, words []string) []int {
	var ans []int
	if len(words) == 0 || len(s) == 0 {
		return ans
	}
	var w = len(words[0])
	var l = len(words) * w
	var checked = map[string]int{}
	for _, word := range words {
		checked[word] ++
	}
	for i := 0; i <= len(s)-l; i++ {
		if containsAllWords(s[i:i+l], checked, w) {
			ans = append(ans, i)
		}
	}
	return ans
}

func containsAllWords(s string, dic map[string]int, w int) bool {
	sLen := len(s)
	tmpDict := map[string]int{}
	for i := 0; i <= sLen-w; i = i + w {
		tmpDict[s[i:i+w]]++
	}
	for k, v := range dic {
		if tmpDict[k] != v {
			return false
		}
	}
	return true
}
