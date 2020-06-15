package _029_divide_two_integers

func divide(dividend int, divisor int) int {
	if dividend == -1<<31 && divisor == -1 {
		return 1<<31 - 1
	}
	flag := true
	if (dividend >= 0 && divisor < 0) || (dividend <= 0 && divisor > 0) {
		flag = false
	}
	if divisor > 0 {
		divisor = -divisor
	}
	if dividend > 0 {
		dividend = -dividend
	}
	bit := 0
	for i := 0; (divisor << i) >= dividend; i++ {
		bit = i
	}
	res := 0
	for i := bit; i >= 0; i-- {
		if dividend <= (divisor << i) {
			res += 1 << i
			dividend -= divisor << i
		}
	}
	if !flag {
		res = -res
	}
	return res
}

func divide1(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if dividend == -1<<31 && divisor == -1 {
		return 1<<31 - 1
	}
	negative := dividend^divisor < 0
	if divisor < 0 {
		divisor = -divisor
	}
	if dividend < 0 {
		dividend = -dividend
	}
	var res = 0
	for i := 31; i >= 0; i-- {
		if (dividend >> i) >= divisor {
			res += 1 << i
			dividend -= divisor << i
		}
	}
	if negative {
		res = -res
	}
	return res
}
