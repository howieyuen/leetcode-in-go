package _526_beautiful_arrangement

func countArrangement(N int) int {
	if N <= 0 {
		return 0
	}
	var used = make([]bool, N+1)
	var ans int
	var backtrace func(index int)
	backtrace = func(index int) {
		if index > N {
			ans++
			return
		}
		for num := 1; num <= N; num++ {
			if used[num] {
				continue
			}
			if num%index == 0 || index%num == 0 {
				used[num] = true
				backtrace(index + 1)
				used[num] = false
			}
		}
	}
	backtrace(1)
	return ans
}
