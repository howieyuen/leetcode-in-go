package _071_simplify_path

import (
	"strings"
)

func simplifyPath(path string) string {
	var ans []string
	splits := strings.Split(path, "/")
	for i := range splits {
		if splits[i] == "" || splits[i] == "." {
			continue
		}
		if splits[i] == ".." && len(ans) > 0 {
			ans = ans[:len(ans)-1]
		} else if splits[i] != ".." {
			ans = append(ans, splits[i])
		}
	}
	return "/" + strings.Join(ans, "/")
}
