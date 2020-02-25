package _202_happy_number

func isHappy(n int) bool {
	var maps = make(map[int]bool)
	maps[n] = true
	for n != 1 {
		n = change(n)
		if maps[n] {
			return false
		}
		maps[n] = true
	}
	return true
}

func change(n int) int {
	sum := 0
	for n != 0 {
		tmp := n % 10
		sum += tmp * tmp
		n /= 10
	}
	return sum
}
