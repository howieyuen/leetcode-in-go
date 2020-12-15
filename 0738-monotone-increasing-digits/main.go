package _738_monotone_increasing_digits

import "strconv"

func monotoneIncreasingDigits(N int) int {
	chars := []byte(strconv.Itoa(N))
	var index = len(chars)
	for i := len(chars) - 1; i > 0; i-- {
		if chars[i-1] > chars[i] {
			chars[i-1]--
			index = i
		}
	}
	for i := index; i < len(chars); i++ {
		chars[i] = '9'
	}
	n, _ := strconv.Atoi(string(chars))
	return n
}
