package _816_ambiguous_coordinates

func ambiguousCoordinates(S string) []string {
	build := func(i, j int) []string {
		var ans []string
		for d := 1; d <= j-i; d++ {
			left := S[i : i+d]
			right := S[i+d : j]
			// 小数点前面，要么等于0，要么不为0
			// 小数点后面，最后的字符不能为0
			if (left[0] != '0' || left == "0") && (right == "" || right[len(right)-1] != '0') {
				if right != "" {
					ans = append(ans, left+"."+right)
				} else {
					ans = append(ans, left+right)
				}
			}
		}
		return ans
	}
	var ans []string
	for i := 2; i < len(S)-1; i++ {
		for _, left := range build(1, i) {
			for _, right := range build(i, len(S)-1) {
				ans = append(ans, "("+left+", "+right+")")
			}
		}
	}
	return ans
}
