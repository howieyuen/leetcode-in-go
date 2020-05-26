package _058_length_of_last_word

import (
	"testing"
)

func Test_lengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{s: "Hello World"}, want: 5},
		{args: args{s: "Hello World "}, want: 5},
		{args: args{s: "Hello Wor ld "}, want: 2},
		{args: args{s: "HelloWorld"}, want: 10},
		{args: args{s: "  "}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
