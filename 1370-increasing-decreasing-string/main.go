package _370_increasing_decreasing_string

func sortString(s string) string {
	var dict = make([]int, 26)
	for i := range s {
		dict[s[i]-'a']++
	}
	var ans = make([]byte, len(s))
	index := 0
	for index < len(ans) {
		for i := 0; i < len(dict); i++ {
			if dict[i] > 0 {
				ans[index] = byte(i + 'a')
				index++
				dict[i]--
			}
		}
		for i := len(dict) - 1; i >= 0; i-- {
			if dict[i] > 0 {
				ans[index] = byte(i + 'a')
				index++
				dict[i]--
			}
		}
	}
	return string(ans)
}
