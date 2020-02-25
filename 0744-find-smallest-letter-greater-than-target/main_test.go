package _744_find_smallest_letter_greater_than_target

import (
	"testing"
)

func Test_nextGreatestLetter(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			args: args{
				letters: []byte("cfj"),
				target:  'a',
			},
			want: 'c',
		},
		{
			args: args{
				letters: []byte("cfj"),
				target:  'c',
			},
			want: 'f',
		},
		{
			args: args{
				letters: []byte("cfj"),
				target:  'd',
			},
			want: 'f',
		},
		{
			args: args{
				letters: []byte("cfj"),
				target:  'g',
			},
			want: 'j',
		},
		{
			args: args{
				letters: []byte("cfj"),
				target:  'j',
			},
			want: 'c',
		},
		{
			args: args{
				letters: []byte("cfj"),
				target:  'k',
			},
			want: 'c',
		},
		{
			args: args{
				letters: []byte("eeeeeennnn"),
				target:  'e',
			},
			want: 'n',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreatestLetter(tt.args.letters, tt.args.target); got != tt.want {
				t.Errorf("nextGreatestLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
