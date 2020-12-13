package _528_random_pick_with_weight

import (
	"math/rand"
)

type Solution struct {
	preW []int
	sum  int
}

func Constructor(w []int) Solution {
	var preW = make([]int, len(w))
	sum := 0
	for i, v := range w {
		sum += v
		preW[i] = sum
	}
	return Solution{
		preW: preW,
		sum:  sum,
	}
}

func (s *Solution) PickIndex() int {
	tmp := rand.Intn(s.sum) + 1
	l, r := 0, len(s.preW)-1
	for l <= r {
		mid := (l + r) / 2
		if tmp == s.preW[mid] {
			return mid
		} else if tmp < s.preW[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
