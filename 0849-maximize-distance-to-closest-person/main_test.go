package _849_maximize_distance_to_closest_person

import "testing"

func Test_maxDistToClosest(t *testing.T) {
	type args struct {
		seats []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{seats: []int{1, 0, 0, 0, 1, 0, 1}}, want: 2},
		{args: args{seats: []int{1, 0, 0, 0}}, want: 3},
		{args: args{seats: []int{1, 0}}, want: 1},
		{args: args{seats: []int{0, 1}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistToClosest(tt.args.seats); got != tt.want {
				t.Errorf("maxDistToClosest() = %v, want %v", got, tt.want)
			}
			if got := maxDistToClosest1(tt.args.seats); got != tt.want {
				t.Errorf("maxDistToClosest1() = %v, want %v", got, tt.want)
			}
			if got := maxDistToClosest2(tt.args.seats); got != tt.want {
				t.Errorf("maxDistToClosest2() = %v, want %v", got, tt.want)
			}
		})
	}
}
