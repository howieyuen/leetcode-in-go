package _328_break_a_palindrome

func breakPalindrome(palindrome string) string {
	var ans string
	chars := []byte(palindrome)
	for i := 0; i < len(chars); i++ {
		c := chars[i]
		for delta := byte('a'); delta <= 'z'; delta++ {
			if delta == c {
				continue
			}
			chars[i] = delta
			if !isPalindrome(chars) {
				ans = min(ans, string(chars))
				chars[i] = c
				break
			}
			chars[i] = c
		}
	}
	return ans
}

func min(a, b string) string {
	if a != "" && b != "" {
		for i := 0; i < len(a) && i < len(b); i++ {
			if a[i] < b[i] {
				return a
			} else if a[i] > b[i] {
				return b
			}
		}
		if len(a) <= len(b) {
			return a
		}
		return b
	} else if a != "" {
		return a
	} else {
		return b
	}
}

func isPalindrome(chars []byte) bool {
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		if chars[i] != chars[j] {
			return false
		}
	}
	return true
}

func breakPalindrome1(palindrome string) string {
	if len(palindrome) <= 1 {
		return ""
	}
	
	for i := 0; i < len(palindrome)/2; i++ {
		if palindrome[i] != 'a' {
			return palindrome[0:i] + "a" + palindrome[i+1:]
		}
	}
	
	return palindrome[0:len(palindrome)-1] + "b"
}
