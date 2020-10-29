package _432_max_difference_you_can_get_from_changing_an_integer

import (
	"strconv"
	"strings"
)

func maxDiff(num int) int {
	numStr := strconv.Itoa(num)
	var max, min = num, num
	for i := 0; i < len(numStr); i++ {
		if numStr[i] != '9' {
			max, _ = strconv.Atoi(strings.Replace(numStr, string(numStr[i]), "9", -1))
			break
		}
	}

	if numStr[0] != '1' {
		min, _ = strconv.Atoi(strings.Replace(numStr, string(numStr[0]), "1", -1))
	} else {
		for i := 1; i < len(numStr); i++ {
			if numStr[i] != '0' && numStr[i] != '1' {
				min, _ = strconv.Atoi(strings.Replace(numStr, string(numStr[i]), "0", -1))
				break
			}
		}
	}

	return max - min
}
