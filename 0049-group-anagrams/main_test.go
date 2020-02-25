package _049_group_anagrams

import (
	`reflect`
	`testing`
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			args: args{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			want: [][]string{
				{"ate", "eat", "tea"},
				{"tan", "nat"},
				{"bat"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
