package _205_isomorphic_strings

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var maps1 = make([]int, 128)
	var maps2 = make([]int, 128)
	for i := 0; i < len(s); i++ {
		if maps1[s[i]] != maps2[t[i]] {
			return false
		} else {
			if maps1[s[i]] == 0 {
				maps1[s[i]] = i + 1
				maps2[t[i]] = i + 1
			}
		}
	}
	return true
}

func isIsomorphic1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var dict1 = make(map[byte]byte)
	var dict2 = make(map[byte]byte)
	for i := range s {
		if v, ok := dict1[s[i]]; !ok {
			if _, ok := dict2[t[i]]; ok {
				return false
			}
			dict1[s[i]] = t[i]
			dict2[t[i]] = s[i]
		} else {
			if v != t[i] || s[i] != dict2[t[i]] {
				return false
			}
		}
	}
	return true
}
