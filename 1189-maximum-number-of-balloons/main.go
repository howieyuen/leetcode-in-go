package _189_maximum_number_of_balloons

func maxNumberOfBalloons(text string) int {
	var b, a, l, o, n int
	for _, ch := range text {
		switch ch {
		case 'b':
			b++
		case 'a':
			a++
		case 'l':
			l++
		case 'o':
			o++
		case 'n':
			n++
		}
	}
	var ans int
	for b > 0 && a > 0 && l > 1 && o > 1 && n > 0 {
		b -= 1
		a -= 1
		l -= 2
		o -= 2
		n -= 1
		ans++
	}
	return ans
}
