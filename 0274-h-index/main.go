package _274_h_index

import (
	"sort"
)

func hIndex(citations []int) int {
	if len(citations) == 0 {
		return 0
	}
	sort.Ints(citations)
	for i := 0; i < len(citations); i++ {
		h := len(citations) - i
		if citations[i] >= h {
			return h
		}
	}
	return 0
}
