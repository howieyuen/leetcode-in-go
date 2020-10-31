package _487_making_file_names_unique

import (
	"strconv"
)

func getFolderNames(names []string) []string {
	var dict = make(map[string]int)
	var ans []string
	for i := range names {
		name := names[i]
		next, ok := dict[name]
		for ok {
			next++
			name = names[i] + "(" + strconv.Itoa(next) + ")"
			_, ok = dict[name]
		}
		dict[name] = 0
		dict[names[i]] = next
		ans = append(ans, name)
	}
	return ans
}
