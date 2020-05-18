package _582_lcof_14_I

const p = 1e9 + 7

func pow3N(n int) int {
	ans := 1
	for i := 0; i < n; i++ {
		ans = (ans * 3) % p
	}
	return ans
}

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	a := n / 3
	b := n % 3
	if b == 0 {
		return pow3N(a) % p
	}
	if b == 1 {
		return pow3N(a-1) * 4 % p
	}
	return pow3N(a) * 2 % p
}
