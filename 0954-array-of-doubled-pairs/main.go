package _954_array_of_doubled_pairs

import (
	"sort"
)

func canReorderDoubled(A []int) bool {
	n := len(A)
	if n == 0 {
		return true
	}
	sort.Ints(A)
	var positives, negatives []int
	for _, x := range A {
		if x < 0 {
			if len(negatives) > 0 && x*2 == negatives[0] {
				negatives = negatives[1:]
			} else {
				negatives = append(negatives, x)
			}
		} else {
			if len(positives) > 0 && x == 2*positives[0] {
				positives = positives[1:]
			} else {
				positives = append(positives, x)
			}
		}
	}
	return len(positives) == 0 && len(negatives) == 0
}

func canReorderDoubled1(A []int) bool {
	dict := map[int]int{}
	for _, v := range A {
		dict[v]++
	}
	
	s := make([]int, 0, len(dict))
	for k := range dict {
		s = append(s, k)
	}
	sort.Ints(s)
	for _, v := range s {
		n := dict[v]
		if n == 0 {
			continue
		}
		dict[v] -= n
		if v < 0 {
			if v%2 == -1 {
				return false
			}
			v /= 2
		} else if v > 0 {
			v *= 2
		} else {
			if n%2 != 0 {
				return false
			}
		}
		if v != 0 {
			if m, ok := dict[v]; !ok || m < n {
				return false
			}
			dict[v] -= n
		}
	}
	return true
}
