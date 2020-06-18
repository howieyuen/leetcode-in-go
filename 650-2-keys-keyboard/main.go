package _50_2_keys_keyboard

// 如果这个数是质数 则这个数只能一个一个的复制得到 操作步数就是这个数本身
// 如果不是质数 则可以由复制得到
// eg: 20可以由10复制得到 10可以由5复制得到 而5是质数 只能一个一个复制
// 所以minStep(20) = 9
func minSteps(n int) int {
	var ans int
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			ans += i
			n /= i
		}
	}
	return ans
}
