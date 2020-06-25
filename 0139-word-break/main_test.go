package _139_word_break

import (
	"testing"
)

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				s:        "leetcode",
				wordDict: []string{"leet", "code"},
			},
			want: true,
		},
		{
			args: args{
				s:        "applepenapple",
				wordDict: []string{"apple", "pen"},
			},
			want: true,
		},
		{
			args: args{
				s:        "catsandog",
				wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak(tt.args.s, tt.args.wordDict); got != tt.want {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
			if got := wordBreak1(tt.args.s, tt.args.wordDict); got != tt.want {
				t.Errorf("wordBreak1() = %v, want %v", got, tt.want)
			}
		})
	}
}
