package _002_find_common_characters

import (
	"math"
)

func commonChars(A []string) []string {
	var minFreq = make([]int, 26)
	for i := range minFreq {
		minFreq[i] = math.MaxInt32
	}
	for _, word := range A {
		var curFreq = make([]int, 26)
		for _, char := range word {
			curFreq[char-'a']++
		}
		for j := range curFreq {
			if minFreq[j] > curFreq[j] {
				minFreq[j] = curFreq[j]
			}
		}
	}
	var ans []string
	for i := range minFreq {
		for j := 0; j < minFreq[i]; j++ {
			ans = append(ans, string(byte('a'+i)))
		}
	}
	return ans
}
