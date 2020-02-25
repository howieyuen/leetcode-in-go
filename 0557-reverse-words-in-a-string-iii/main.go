package _557_reverse_words_in_a_string_iii

func reverseWords(s string) string {
	tmp := []byte(s)
	for i := 0; i < len(tmp); i++ {
		j := i
		for j < len(tmp) {
			if tmp[j] == ' ' {
				break
			}
			j++
		}
		reverse(tmp, i, j-1)
		i = j
	}
	return string(tmp)
}

func reverse(tmp []byte, i, j int) {
	for ; i < j; i, j = i+1, j-1 {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}
}
