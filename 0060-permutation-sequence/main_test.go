package _060_permutation_sequence

import (
	"reflect"
	"testing"
)

func Test_getPermutation(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{n: 3, k: 1}, want: "123"},
		{args: args{n: 3, k: 2}, want: "132"},
		{args: args{n: 3, k: 3}, want: "213"},
		{args: args{n: 3, k: 4}, want: "231"},
		{args: args{n: 3, k: 5}, want: "312"},
		{args: args{n: 3, k: 6}, want: "321"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPermutation(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("getPermutation() = %v, want %v", got, tt.want)
			}
			if got := getPermutation1(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("getPermutation1() = %v, want %v", got, tt.want)
			}
			if got := getPermutation2(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("getPermutation2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAllPermutations(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{args: args{n: 3}, want: []string{"123", "132", "213", "231", "321", "312"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAllPermutations(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllPermutations() = %v, want %v", got, tt.want)
			}
		})
	}
}
