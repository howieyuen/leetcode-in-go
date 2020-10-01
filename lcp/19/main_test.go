package _9

import (
	"testing"
)

func Test_minimumOperations(t *testing.T) {
	type args struct {
		leaves string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{leaves: "rrryyyrryyyrr"}, want: 2},
		{args: args{leaves: "ryr"}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumOperations(tt.args.leaves); got != tt.want {
				t.Errorf("minimumOperations() = %v, want %v", got, tt.want)
			}
			if got := minimumOperations1(tt.args.leaves); got != tt.want {
				t.Errorf("minimumOperations1() = %v, want %v", got, tt.want)
			}
		})
	}
}
