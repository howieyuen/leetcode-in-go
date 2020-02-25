package _367_valid_perfect_square

/*
给定一个正整数 num，编写一个函数，如果 num 是一个完全平方数，则返回 True，否则返回 False。
*/
func isPerfectSquare(num int) bool {
	if num < 0 {
		return false
	}
	i, j := 0, num
	for i <= j {
		k := i + (j-i)/2
		if k*k == num {
			return true
		} else if k*k < num {
			i = k + 1
		} else {
			j = k - 1
		}
	}
	return false
}
