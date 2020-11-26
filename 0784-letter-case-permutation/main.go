package _784_letter_case_permutation

func letterCasePermutation(S string) []string {
	chars := []byte(S)
	var ans []string
	var backtrace func(index int, cur []byte)
	backtrace = func(index int, cur []byte) {
		if index == len(cur) {
			ans = append(ans, string(cur))
			return
		}
		backtrace(index+1, cur)
		if cur[index] >= 'a' && cur[index] <= 'z' {
			cur[index] -= 32
			backtrace(index+1, cur)
		} else if cur[index] >= 'A' && cur[index] <= 'Z' {
			cur[index] += 32
			backtrace(index+1, cur)
		}
	}
	backtrace(0, chars)
	return ans
}
