package _273_integer_to_english_words

import (
	"testing"
)

func Test_numberToWords(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{num: 0},
			want: "Zero",
		},
		{
			args: args{num: 123},
			want: "One Hundred Twenty Three",
		},
		{
			args: args{num: 12345},
			want: "Twelve Thousand Three Hundred Forty Five",
		},
		{
			args: args{num: 1234567},
			want: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven",
		},
		{
			args: args{num: 1234567891},
			want: "One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberToWords(tt.args.num); got != tt.want {
				t.Errorf("numberToWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
