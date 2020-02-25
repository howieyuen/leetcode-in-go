package _028_implement_strstr

func strStr(haystack string, needle string) int {
	length := len(needle)
	for i := 0; i <= len(haystack)-length; i++ {
		if string(haystack[i:i+length]) == needle {
			return i
		}
	}
	return -1
}
