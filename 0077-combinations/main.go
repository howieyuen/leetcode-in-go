package _077_combinations

func combine(n int, k int) [][]int {
	var base = make([]int, n)
	for i := range base {
		base[i] = i + 1
	}
	var ans = make([][]int, 0, factor(n)/factor(k)/factor(n-k))
	var backtrace func(index int, cur []int)
	backtrace = func(index int, cur []int) {
		if len(cur) == k {
			temp := make([]int, k)
			copy(temp, cur)
			ans = append(ans, temp)
			return
		}
		for i := index; i < n; i++ {
			cur = append(cur, base[i])
			backtrace(i+1, cur)
			cur = cur[:len(cur)-1]
		}
	}
	cur := make([]int, 0, k)
	backtrace(0, cur)
	return ans
}

func factor(n int) int {
	f := 1
	for i := 1; i <= n; i++ {
		f *= i
	}
	return f
}

func combine1(n int, k int) [][]int {
	l := make([]int, k)
	i := 0
	var ret [][]int
	for i >= 0 {
		l[i]++
		if l[i] > n-k+i+1 {
			i--
		} else if i == k-1 {
			t := make([]int, k)
			copy(t, l)
			ret = append(ret, t)
		} else {
			i++
			l[i] = l[i-1]
		}
	}
	return ret
}
