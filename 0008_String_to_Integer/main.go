package _008_String_to_Integer

const (
	Int32MaxValue = int(^uint32(0) >> 1)
	Int32MinValue = ^Int32MaxValue
)

func myAtoi(str string) int {
	first := true
	flag := 1
	ans := 0
	for i := range str {
		if first {
			if str[i] == ' ' {
				continue
			} else if str[i] == '+' {
				first = false
			} else if str[i] == '-' {
				flag = -1
				first = false
			} else if str[i] >= '0' && str[i] <= '9' {
				ans = ans*10 + int(str[i]-'0')
				first = false
			} else {
				return 0
			}
		} else {
			if str[i] >= '0' && str[i] <= '9' {
				if flag == 1 && (ans > Int32MaxValue/10 || (ans == Int32MaxValue/10 && str[i] > '7')) {
					return Int32MaxValue
				}
				if flag == -1 && (ans > Int32MaxValue/10 || (ans == Int32MaxValue/10 && str[i] > '8')) {
					return Int32MinValue
				}
				ans = ans*10 + int(str[i]-'0')
			} else {
				return flag * ans
			}
		}
	}
	return flag * ans
}
