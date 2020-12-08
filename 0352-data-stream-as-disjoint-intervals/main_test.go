package _352_data_stream_as_disjoint_intervals

import (
	"reflect"
	"testing"
)

func TestSummaryRanges_GetIntervals(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "test1",
			nums: []int{1},
			want: [][]int{{1, 1}},
		},
		{
			name: "test2",
			nums: []int{1, 3},
			want: [][]int{{1, 1}, {3, 3}},
		},
		{
			name: "test3",
			nums: []int{1, 3, 7},
			want: [][]int{{1, 1}, {3, 3}, {7, 7}},
		},
		{
			name: "test4",
			nums: []int{1, 3, 7, 2},
			want: [][]int{{1, 3}, {7, 7}},
		},
		{
			name: "test4",
			nums: []int{1, 3, 7, 2, 6},
			want: [][]int{{1, 3}, {6, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Constructor()
			for _, x := range tt.nums {
				s.AddNum(x)
			}
			if got := s.GetIntervals(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
