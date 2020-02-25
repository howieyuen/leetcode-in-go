package _050_powx_n

/*
实现 pow(x, n) ，即计算 x 的 n 次幂函数。
*/

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

func myPow1(x float64, n int) float64 {
	if n < 0 {
		return 1.0 / myPow(x, -n)
	}
	if n == 0 {
		return 1.0
	}
	if n == 1 {
		return x
	}
	if n == 2 {
		return x * x
	}
	return myPow(myPow(x, n/2), 2) * myPow(x, n%2)
}
