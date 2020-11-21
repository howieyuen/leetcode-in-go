package _886_possible_bipartition

import (
	"testing"
)

func Test_possibleBipartition(t *testing.T) {
	type args struct {
		N        int
		dislikes [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				N:        4,
				dislikes: [][]int{{1, 2}, {1, 3}, {2, 4}},
			},
			want: true,
		},
		{
			args: args{
				N:        3,
				dislikes: [][]int{{1, 2}, {1, 3}, {2, 3}},
			},
			want: false,
		},
		{
			args: args{
				N:        5,
				dislikes: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := possibleBipartition(tt.args.N, tt.args.dislikes); got != tt.want {
				t.Errorf("possibleBipartition() = %v, want %v", got, tt.want)
			}
			if got := possibleBipartition1(tt.args.N, tt.args.dislikes); got != tt.want {
				t.Errorf("possibleBipartition1() = %v, want %v", got, tt.want)
			}
			if got := possibleBipartition2(tt.args.N, tt.args.dislikes); got != tt.want {
				t.Errorf("possibleBipartition2() = %v, want %v", got, tt.want)
			}
			if got := possibleBipartition3(tt.args.N, tt.args.dislikes); got != tt.want {
				t.Errorf("possibleBipartition3() = %v, want %v", got, tt.want)
			}
		})
	}
}
