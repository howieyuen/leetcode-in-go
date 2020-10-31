package _7_15

import (
	"testing"
)

func Test_longestWord(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				words: []string{"cat", "banana", "dog", "nana", "walk", "walker", "dogwalker"},
			},
			want: "dogwalker",
		},
		{
			args: args{
				words: []string{"qlmql", "qlmqlmqqlqmqqlq", "mqqlqmqqlqmqqlq", "mqqlq", "mqqlqlmlsmqq", "qmlmmmmsm",
					"lmlsmqq", "slmsqq", "mslqssl", "mqqlqmqqlq"},
			},
			want: "mqqlqmqqlqmqqlq",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestWord(tt.args.words); got != tt.want {
				t.Errorf("longestWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
