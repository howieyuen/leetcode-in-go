package _352_data_stream_as_disjoint_intervals

import (
	"sort"
)

type SummaryRanges struct {
	sort.IntSlice
}

/** Initialize your data structure here. */
func Constructor() SummaryRanges {
	return SummaryRanges{}
}

func (s *SummaryRanges) AddNum(val int) {
	if index := s.Search(val); index < s.Len() && s.IntSlice[index] == val {
		return
	} else {
		var deepCopy = make([]int, 0, s.Len()+1)
		deepCopy = append(deepCopy, s.IntSlice[:index]...)
		deepCopy = append(deepCopy, val)
		deepCopy = append(deepCopy, s.IntSlice[index:]...)
		s.IntSlice = deepCopy
	}
}

func (s *SummaryRanges) GetIntervals() [][]int {
	var ans [][]int
	if s.Len() == 0 {
		return ans
	}
	var last, begin = s.IntSlice[0], s.IntSlice[0]
	for _, x := range s.IntSlice {
		if last == x || last+1 == x {
			last = x
		} else {
			ans = append(ans, []int{begin, last})
			begin = x
			last = x
		}
	}
	ans = append(ans, []int{begin, s.IntSlice[s.Len()-1]})
	return ans
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(val);
 * param_2 := obj.GetIntervals();
 */
