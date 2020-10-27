package _513_number_of_substrings_with_only_1s

func numSub(s string) int {
	const base = 1000000007
	sum := func(n int) int {
		return n * (n + 1) / 2
	}
	var ans int
	for i := 0; i < len(s); i++ {
		start := i
		for start < len(s) && s[start] == '1' {
			start++
		}
		ans += sum(start - i)
		ans %= base
		i = start
	}
	return ans
}
