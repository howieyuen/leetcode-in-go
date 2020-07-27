package _407_trapping_rain_water_ii

import (
	"testing"
)

func Test_trapRainWater(t *testing.T) {
	type args struct {
		heightMap [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{heightMap: [][]int{
				{12, 13, 1, 12},
				{13, 4, 13, 12},
				{13, 8, 10, 12},
				{12, 13, 12, 12},
				{13, 13, 13, 13},
			}},
			want: 14,
		},
		{
			args: args{heightMap: [][]int{
				{1, 4, 3, 1, 3, 2},
				{3, 2, 1, 3, 2, 4},
				{2, 3, 3, 2, 3, 1},
			}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trapRainWater(tt.args.heightMap); got != tt.want {
				t.Errorf("trapRainWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
