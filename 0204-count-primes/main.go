package _204_count_primes

// 素数筛
func countPrimes(n int) int {
	var primes = make([]bool, n)
	count := 0
	for i := 2; i < n; i++ {
		if !primes[i] {
			count++
			for j := 2; j*i < n; j++ {
				primes[i*j] = true
			}
		}
	}
	return count
}
