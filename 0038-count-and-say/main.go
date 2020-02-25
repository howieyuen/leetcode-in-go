package _038_count_and_say

func countAndSay(n int) string {
	return generateNext(1, n, "1")
}

func generateNext(i, n int, cur string) string {
	if i == n {
		return cur
	} else {
		var count = 1
		var b = cur[0]
		var ans []byte
		for i := 1; i < len(cur); i++ {
			if cur[i] == b {
				count++
			} else {
				ans = append(ans, byte(count+'0'), b)
				count = 1
				b = cur[i]
			}
		}
		ans = append(ans, byte(count+'0'), b)
		return generateNext(i+1, n, string(ans))
	}
}
