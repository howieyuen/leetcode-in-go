package _066_plus_one

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		tmp := digits[i] + carry
		digits[i] = tmp % 10
		carry = tmp / 10
	}
	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}
	return digits
}
