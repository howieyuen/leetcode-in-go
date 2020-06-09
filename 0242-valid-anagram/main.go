package _242_valid_anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var count = make([]byte, 26)
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}
	for i := range count {
		if count[i] != 0 {
			return false
		}
	}
	return true
}
