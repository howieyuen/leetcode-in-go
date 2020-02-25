package _049_group_anagrams

import (
	`sort`
)

func groupAnagrams(strs []string) [][]string {
	var maps = make(map[string][]string)
	for _, str := range strs {
		bytes := []byte(str)
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		tmp := string(bytes)
		if _, ok := maps[tmp]; ok {
			maps[tmp] = append(maps[tmp], str)
		} else {
			maps[tmp] = []string{str}
		}
	}
	var ans [][]string
	for _, val := range maps {
		ans = append(ans, val)
	}
	return ans
}
