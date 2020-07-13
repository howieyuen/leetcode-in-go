package _8

func permutation(s string) []string {
	var ans []string
	if s == "" {
		return nil
	}
	var dfs func(chars []byte, index int)
	dfs = func(chars []byte, index int) {
		if index == len(chars) {
			ans = append(ans, string(chars))
		}
		m := make(map[byte]bool)
		for i := index; i < len(chars); i++ {
			if m[chars[i]] {
				continue
			}
			m[chars[i]] = true
			chars[i], chars[index] = chars[index], chars[i]
			dfs(chars, index+1)
			chars[i], chars[index] = chars[index], chars[i]
		}
	}
	dfs([]byte(s), 0)
	return ans
}
