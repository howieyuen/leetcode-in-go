package _779_k_th_symbol_in_grammar

import (
	"testing"
)

func Test_kthGrammar(t *testing.T) {
	type args struct {
		N int
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				N: 1,
				K: 1,
			},
			want: 0,
		},
		{
			args: args{
				N: 2,
				K: 1,
			},
			want: 0,
		},
		{
			args: args{
				N: 2,
				K: 2,
			},
			want: 1,
		},
		{
			args: args{
				N: 4,
				K: 5,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthGrammar(tt.args.N, tt.args.K); got != tt.want {
				t.Errorf("kthGrammar() = %v, want %v", got, tt.want)
			}
		})
	}
}
