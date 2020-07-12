package test4

func minInteger(num string, k int) string {
	if k == 0 {
		return num
	}
	var chars = []byte(num)
	for i := 0; i < len(chars); i++ {
		min := chars[i]
		findBigger := false
		var index int
		// 贪心选择最小值
		for j := i + 1; j <= i+k && j < len(chars); j++ {
			if min > chars[j] {
				min = chars[j]
				index = j
				findBigger = true
			}
		}
		// 找到最小值，交换
		if findBigger {
			for a := index; a > i && k != 0; a-- {
				chars[a], chars[a-1] = chars[a-1], chars[a]
				k--
			}
		}
		// 交换次数剩余为0
		if k == 0 {
			break
		}
	}
	return string(chars)
}
