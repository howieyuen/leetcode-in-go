package _387_first_unique_character_in_a_string

func firstUniqChar(s string) int {
	var maps = make(map[byte]int)
	for i := range s {
		if _, ok := maps[s[i]]; !ok {
			maps[s[i]] = 1
		} else {
			maps[s[i]]++
		}
	}
	for i := range s {
		if maps[s[i]] == 1 {
			return i
		}
	}
	return -1
}
