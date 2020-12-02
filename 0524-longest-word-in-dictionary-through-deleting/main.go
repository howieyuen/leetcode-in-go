package _524_longest_word_in_dictionary_through_deleting

func findLongestWord(s string, d []string) string {
	var candidate string
	var maxLen = 0
	for _, can := range d {
		i, j := 0, 0
		for i < len(s) && j < len(can) {
			if s[i] == can[j] {
				j++
			}
			i++
		}
		if j != len(can) {
			continue
		}
		if maxLen < len(can) || maxLen == len(can) && can < candidate {
			candidate = can
			maxLen = len(can)
		}
	}
	return candidate
}
