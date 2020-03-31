package _089_gray_code

func grayCode(n int) []int {
	var ans []int
	ans = append(ans, 0)
	head := 1
	for i := 0; i < n; i++ {
		for j := len(ans) - 1; j >= 0; j-- {
			ans = append(ans, head+ans[j])
		}
		head <<= 1
	}
	return ans
}

