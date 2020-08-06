package _211_add_and_search_word_data_structure_design

import (
	"testing"
)

func TestWordDictionary_Search1(t *testing.T) {
	type args struct {
		add    []string
		search []string
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			args: args{
				add:    []string{"bad", "dad", "mad"},
				search: []string{"pad", "bad", ".ad", "b.."},
			},
			want: []bool{false, true, true, true},
		},
		{
			args: args{
				add:    []string{"a", "a"},
				search: []string{".", "a", "aa", ".a", "a."},
			},
			want: []bool{true, true, false, false, false},
		},
	}
	
	for _, tt := range tests {
		wordDictionary := Constructor()
		for _, word := range tt.args.add {
			wordDictionary.AddWord(word)
		}
		t.Run(tt.name, func(t *testing.T) {
			for i, word := range tt.args.search {
				if got := wordDictionary.Search(word); got != tt.want[i] {
					t.Errorf("Search() = %v, want %v", got, tt.want[i])
				}
			}
		})
	}
}
