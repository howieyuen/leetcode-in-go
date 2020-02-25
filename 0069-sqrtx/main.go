package _069_sqrtx

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	left := 2
	right := x / 2
	var mid int
	for left <= right {
		mid = left + (right-left)/2
		t := mid * mid
		if t == x {
			return mid
		} else if t < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}
