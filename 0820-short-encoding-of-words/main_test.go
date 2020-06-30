package _820_short_encoding_of_words

import (
	"testing"
)

func Test_minimumLengthEncoding(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{words: []string{"time", "me", "bell"}}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumLengthEncoding(tt.args.words); got != tt.want {
				t.Errorf("minimumLengthEncoding() = %v, want %v", got, tt.want)
			}
			if got := minimumLengthEncoding1(tt.args.words); got != tt.want {
				t.Errorf("minimumLengthEncoding1() = %v, want %v", got, tt.want)
			}
		})
	}
}
