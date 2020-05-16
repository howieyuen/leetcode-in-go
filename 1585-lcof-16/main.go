package _585_lcof_16

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	var ans float64 = 1
	var cur = x
	for i := n; i > 0; i /= 2 {
		if i%2 == 1 {
			ans *= cur
		}
		cur *= cur
	}
	return ans
}
