package _6_18

func patternMatching(pattern string, value string) bool {
	var countA, countB = 0, 0
	for i := range pattern {
		if pattern[i] == 'a' {
			countA++
		} else {
			countB++
		}
	}
	if countA < countB {
		countA, countB = countB, countA
		tmp := ""
		for i := range pattern {
			if pattern[i] == 'a' {
				tmp += "b"
			} else {
				tmp += "a"
			}
		}
		pattern = tmp
	}
	if len(value) == 0 {
		return countB == 0
	}
	if len(pattern) == 0 {
		return false
	}
	for lenA := 0; lenA*countA <= len(value); lenA++ {
		rest := len(value) - lenA*countA
		if countB == 0 && rest == 0 || countB != 0 && rest%countB == 0 {
			var lenB int
			if countB == 0 {
				lenB = 0
			} else {
				lenB = rest / countB
			}
			pos, correct := 0, true
			valueA, valueB := "", ""
			for i := 0; i < len(pattern); i++ {
				if pattern[i] == 'a' {
					sub := value[pos : pos+lenA]
					if len(valueA) == 0 {
						valueA = sub
					} else if sub != valueA {
						correct = false
						break
					}
					pos += lenA
				} else if pattern[i] == 'b' {
					sub := value[pos : pos+lenB]
					if len(valueB) == 0 {
						valueB = sub
					} else if sub != valueB {
						correct = false
						break
					}
					pos += lenB
				}
			}
			if correct && valueA != valueB {
				return true
			}
		}
	}
	return false
}
