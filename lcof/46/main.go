package _6

// 类比跳台阶
func translateNum(num int) int {
	if num < 10 {
		return 1
	}
	b := num % 100
	if b < 10 || b >= 26 {
		return translateNum(num / 10)
	}
	return translateNum(num/10) + translateNum(num/100)
}

