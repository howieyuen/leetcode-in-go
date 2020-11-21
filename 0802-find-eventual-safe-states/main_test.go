package _802_find_eventual_safe_states

import (
	"reflect"
	"testing"
)

func Test_eventualafeNodes(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{graph: [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}},
			want: []int{2, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eventualSafeNodes(tt.args.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("eventualafeNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
