package _990_satisfiability_of_equality_equations

import (
	"testing"
)

func Test_equationsPossible(t *testing.T) {
	type args struct {
		equations []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{equations: []string{"a==b", "b!=a"}}, want: false},
		{args: args{equations: []string{"b==a", "a==b"}}, want: true},
		{args: args{equations: []string{"a==b", "b==c", "a==c"}}, want: true},
		{args: args{equations: []string{"a==b", "b!=c", "c==a"}}, want: false},
		{args: args{equations: []string{"c==c", "b==d", "x!=z"}}, want: true},
		{args: args{equations: []string{"c==c","f!=a","f==b","b==c"}}, want: true},
		{args: args{equations: []string{"c==c"}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equationsPossible(tt.args.equations); got != tt.want {
				t.Errorf("equationsPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
